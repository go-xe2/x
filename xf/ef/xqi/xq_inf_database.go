package xqi

import (
	"github.com/go-xe2/x/core/exception"
	"time"
)

type BasicDB interface {
	Close()
	BeginTran() (err exception.IException)
	Rollback() (err exception.IException)
	Commit() (err exception.IException)
	Transaction(closers ...func(session BasicDB) error) (err exception.IException)
	// 查询数据并使用指定的binder数据绑定器序列化数据后返回，binder绑定器名称格式为 binderName:options,
	// 不传options时写为binderName,options格式支持yaml,json,xml,toml格式
	Query(binder TBinderName, szSql string, args ...interface{}) (result interface{}, err exception.IException)
	QueryBind(binder DbQueryBinder, szSql string, args ...interface{}) (result interface{}, err exception.IException)
	//Query(makeRowFn func(row int, colInfos *[]QueryColValue) interface{}, szSql string, args ...interface{}) (interface{}, error)
	Execute(szSql string, args ...interface{}) (int64, exception.IException)
	GetConn() DbConn
	LastInsertId() int64
	LastSql() string
	This() interface{}
	// 最后执行消耗时间
	LastSqlDuration() time.Duration
}

type DatabaseInitializer interface {
	Connection(conf ...interface{}) exception.IException
	Use(closers ...func(e DbConn))
}

type Database interface {
	BasicDB
	DatabaseInitializer
	Logger() DbLogger
	Conn() DbConn
	// 表前掇
	Prefix() string
	// 数据库驱动名称
	Driver() string
	// 创建数据库表
	Create() exception.IException
	// 升级数据库表
	Upgrade() exception.IException
}
