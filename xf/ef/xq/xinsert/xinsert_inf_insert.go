package xinsert

import "github.com/go-xe2/x/xf/ef/xqi"

type SqlInsertExecute interface {
	Execute() (int, error)
}

type SqlInsertFrom interface {
	From(table xqi.SqlTable) SqlInsertExecute
	Execute() (int, error)
}

type SqlInsert interface {
	DB() xqi.Database
	Values(values ...xqi.FieldValue) SqlInsertFrom
}
