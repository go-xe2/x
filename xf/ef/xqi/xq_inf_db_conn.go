package xqi

import (
	"database/sql"
	"github.com/go-xe2/x/core/exception"
)

type DbConn interface {
	// 使用部件
	Use(closers ...func(e DbConn))
	Ping() exception.IException
	SetPrefix(pre string)
	GetPrefix() string
	GetDriver() string
	// GetQueryDB : 获取一个从库用来做查询操作
	GetQueryDB() *sql.DB
	// GetExecuteDB : 获取一个主库用来做查询之外的操作
	GetExecuteDB() *sql.DB
	// 获取日志接口
	GetLogger() DbLogger
	SetLogger(log DbLogger)
	This() interface{}
}
