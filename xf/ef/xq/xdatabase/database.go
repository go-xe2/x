package xdatabase

import (
	"github.com/go-xe2/x/core/exception"
	"github.com/go-xe2/x/sync/xsafeMap"
	"github.com/go-xe2/x/xf/ef/xqi"
)

type TDatabase struct {
	*TBaseDb
}

var _ xqi.Database = (*TDatabase)(nil)

const defaultDbName = "xq.database"

var databaseInstances = xsafeMap.NewStrAnyMap()

// 获取数据库
func DB(dbName ...string) xqi.Database {
	name := defaultDbName
	if len(dbName) > 0 {
		name = dbName[0]
	}
	if v := databaseInstances.Get(name); v != nil {
		return v.(xqi.Database)
	}
	inst := newDatabase()
	databaseInstances.Set(name, inst)
	return inst
}

func newDatabase() *TDatabase {
	inst := &TDatabase{}
	base := newBaseDb(nil, inst)
	inst.TBaseDb = base
	return inst
}

func (db *TDatabase) Logger() xqi.DbLogger {
	if db.conn == nil {
		return nil
	}
	return db.conn.GetLogger()
}

func (db *TDatabase) Conn() xqi.DbConn {
	return db.conn
}

// 表前掇
func (db *TDatabase) Prefix() string {
	if db.conn == nil {
		return ""
	}
	return db.conn.GetPrefix()
}

// 数据库驱动名称
func (db *TDatabase) Driver() string {
	if db.conn == nil {
		return "未配置数据库连接"
	}
	return db.conn.GetDriver()
}

// 创建数据库表方法
func (db *TDatabase) Create() exception.IException {
	return nil
}

// 升级数据库表方法
func (db *TDatabase) Upgrade() exception.IException {
	return nil
}
