package xdatabase

import (
	"github.com/go-xe2/x/core/exception"
	"github.com/go-xe2/x/xf/ef/xqi"
)

func (db *TBaseDb) Connection(conf ...interface{}) exception.IException {
	if conn, err := NewDbConn(conf...); err != nil {
		return err
	} else {
		db.conn = conn
		db.master = conn.GetExecuteDB()
		db.slave = conn.GetQueryDB()
	}
	return nil
}

func (db *TBaseDb) Use(closers ...func(e xqi.DbConn)) {
	db.conn.Use(closers...)
}
