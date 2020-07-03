package xdatabase

import (
	"database/sql"
	"encoding/hex"
	"fmt"
	"github.com/go-xe2/x/core/exception"
	"github.com/go-xe2/x/xf/ef/xdbUtil"
	"github.com/go-xe2/x/xf/ef/xq/xbinder"
	. "github.com/go-xe2/x/xf/ef/xqi"
	"strings"
	"time"
)

type TBaseDb struct {
	this         interface{}
	conn         DbConn
	tx           *sql.Tx
	lastInsertId int64
	sqlLogs      []string
	lastSql      string
	union        interface{}
	transaction  bool
	master       *sql.DB
	slave        *sql.DB
	lastErr      exception.IException
	// 最后执行时间
	lastDuration time.Duration
}

var _ BasicDB = (*TBaseDb)(nil)

func newBaseDb(inherited ...interface{}) *TBaseDb {
	inst := &TBaseDb{
		master: nil,
		slave:  nil,
	}
	inst.this = inst
	if len(inherited) > 0 {
		if _, ok := inherited[0].(BasicDB); ok {
			inst.this = inherited[0]
		}
	}
	return inst
}

func (db *TBaseDb) This() interface{} {
	return db.this
}

func (db *TBaseDb) Close() {
}

func (db *TBaseDb) BeginTran() exception.IException {
	var e error
	if db.tx, e = db.master.Begin(); e != nil {
		return exception.Wrap(e, "开启事务失败")
	}
	db.transaction = true
	return nil
}

func (db *TBaseDb) Rollback() exception.IException {
	if err := db.tx.Rollback(); err != nil {
		return exception.Wrap(err, "回滚事务失败")
	}
	db.tx = nil
	db.transaction = false
	return nil
}

func (db *TBaseDb) Commit() exception.IException {
	if err := db.tx.Commit(); err != nil {
		return exception.Wrap(err, "提交事务失败")
	}
	db.tx = nil
	db.transaction = false
	return nil
}

func (db *TBaseDb) Transaction(closers ...func(session BasicDB) error) exception.IException {
	logger := db.GetConn().GetLogger()
	if err := db.BeginTran(); err != nil {
		logger.Error(err.Error())
		return err
	}
	for _, closer := range closers {
		if err := closer(db); err != nil {
			logger.Error(err.Error())
			_ = db.Rollback()
			return exception.Wrap(err, "执行事务失败")
		}
	}
	return db.Commit()
}

func (db *TBaseDb) scan(rows *sql.Rows, binder DbQueryBinder) (result interface{}, err error) {
	// 如果不需要绑定, 则需要初始化一下binder
	columnTypes, _ := rows.ColumnTypes()
	count := len(columnTypes)
	values := make([]interface{}, count)
	valueRef := make([]interface{}, count)
	colInfos := make([]*TQueryColInfo, count)
	colValues := make([]QueryColValue, count)
	for i := 0; i < count; i++ {
		col := columnTypes[i]
		valueRef[i] = &values[i]
		colInfos[i] = NewQueryColInfo(i, col.Name(), col.ScanType(), col.DatabaseTypeName())
		colValues[i] = &TQueryColValue{
			TQueryColInfo: colInfos[i],
			ColValue:      nil,
		}
	}
	binder.StartBuild(colInfos...)
	row := 1
	for rows.Next() {
		if err = rows.Scan(valueRef...); err != nil {
			return
		}
		if !binder.StartBuildRow(row, count) {
			continue
		}
		for i := 0; i < count; i++ {
			var v = values[i]
			col := colInfos[i]
			if vs, ok := v.([]byte); ok {
				if col.DbType == "BINARY" {
					v = "0x" + hex.EncodeToString(vs)
				} else {
					v = string(vs)
				}
			}
			colValues[i].ColValue = v
		}
		item, exit := binder.BuildRow(row, &colValues)
		binder.EndBuildRow(item)
		row++
		if exit {
			break
		}
	}
	return binder.EndBuild(), nil
}

func (db *TBaseDb) Query(binder TBinderName, szSql string, args ...interface{}) (result interface{}, err exception.IException) {
	oBinder := xbinder.GetQueryBinder(binder)
	if oBinder == nil {
		return nil, exception.Newf("数据绑定器%s不存在", binder)
	}
	return db.query(oBinder, szSql, args...)
}

func (db *TBaseDb) QueryBind(binder DbQueryBinder, szSql string, args ...interface{}) (result interface{}, err exception.IException) {
	return db.query(binder, szSql, args...)
}

func (db *TBaseDb) query(binder DbQueryBinder, szSql string, args ...interface{}) (interface{}, exception.IException) {
	if binder == nil {
		binder = xbinder.NewQryMapBinder(nil)
	}
	// 清除错误
	db.lastErr = nil

	var result interface{} = nil

	var err error
	logger := db.GetConn().GetLogger()

	result, err = xdbUtil.WithRunTimeContext(func() (interface{}, error) {

		// 记录sqlLog
		db.lastSql = fmt.Sprint(szSql, ", ", args)

		var stmt *sql.Stmt
		var rows *sql.Rows
		// 如果是事务, 则从主库中读写
		if db.tx == nil {
			stmt, err = db.slave.Prepare(szSql)
		} else {
			stmt, err = db.master.Prepare(szSql)
		}
		if err != nil {
			return nil, err
		}

		defer func() {
			if err = stmt.Close(); err != nil {
				panic(err)
			}
		}()

		rows, err = stmt.Query(args...)
		if err != nil {
			return nil, err
		}

		// make sure we always close rows
		defer func() {
			if err = rows.Close(); err != nil {
				panic(err)
			}
		}()

		result, err = db.scan(rows, binder)
		if err != nil {
			return nil, err
		}

		return result, nil
	}, func(duration time.Duration, err error) {
		if err == nil {
			// 记录最后执行时间
			db.lastDuration = duration
			if duration.Seconds() > logger.EnableSlowLog() {
				logger.Slow(db.LastSql(), duration)
			} else {
				logger.Sql(db.LastSql(), duration)
			}
		}
	})
	if err != nil {
		db.lastErr = exception.Wrap(err, "查询出错")
		logger.Error("查询出错:" + db.lastErr.Error() + ", sql:" + db.lastSql)
		return nil, db.lastErr
	}
	// 返回数据结果
	return result, nil
}

func (db *TBaseDb) Execute(szSql string, args ...interface{}) (int64, exception.IException) {
	// 清除错误
	db.lastErr = nil
	db.lastDuration = 0
	var rowsAffected int64

	// 记录开始时间
	logger := db.GetConn().GetLogger()

	_, err := xdbUtil.WithRunTimeContext(func() (interface{}, error) {

		db.lastSql = fmt.Sprint(szSql, ", ", args)

		var opType = strings.ToLower(szSql[0:6])

		if opType == "select" {
			return 0, exception.NewText("不能使用Execute执行查询语句")
		}

		var stmt *sql.Stmt
		var err error
		if db.tx == nil {
			stmt, err = db.master.Prepare(szSql)
		} else {
			stmt, err = db.tx.Prepare(szSql)
		}

		if err != nil {
			return 0, err
		}
		//var err error
		defer func() {
			if err = stmt.Close(); err != nil {
				panic(err)
			}
		}()

		result, err := stmt.Exec(args...)
		if err != nil {
			return 0, err
		}

		if opType == "insert" {
			// 获取最后插入的主键id,如果主键不是int类型将返回0
			var lastInsertId int64 = 0
			if lastInsertId, err = result.LastInsertId(); err == nil {
				db.lastInsertId = lastInsertId
			} else {
				return 0, err
			}
		}
		rowsAffected, err = result.RowsAffected()
		if err != nil {
			return 0, err
		}
		if opType == "insert" && db.lastInsertId == 0 {
			// 获取最后插入的主键id,如果主键不是int类型则返回影响行数
			db.lastInsertId = rowsAffected
		}
		return rowsAffected, nil
	}, func(duration time.Duration, err error) {
		if err == nil {
			// 记录最后执行时间
			db.lastDuration = duration
			if duration.Seconds() > logger.EnableSlowLog() {
				logger.Slow(db.LastSql(), duration)
			} else {
				logger.Sql(db.LastSql(), duration)
			}
		}
	})

	if err != nil {
		logger.Error("执行语句出错:" + err.Error() + ", sql:" + db.lastSql)
		db.lastErr = exception.Wrap(err, "执行语句出错")
		return 0, db.lastErr
	}
	return rowsAffected, nil
}

func (db *TBaseDb) GetConn() DbConn {
	return db.conn
}

func (db *TBaseDb) LastInsertId() int64 {
	return db.lastInsertId
}

func (db *TBaseDb) LastSql() string {
	return db.lastSql
}

func (db *TBaseDb) GetErr() exception.IException {
	return db.lastErr
}

func (db *TBaseDb) LastSqlDuration() time.Duration {
	return db.LastSqlDuration()
}
