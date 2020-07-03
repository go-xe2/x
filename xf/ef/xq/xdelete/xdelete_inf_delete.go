package xdelete

import "github.com/go-xe2/x/xf/ef/xqi"

type SqlDeleteExecute interface {
	Execute() (int, error)
}

type SqlDelete interface {
	Where(where ...xqi.SqlCondition) SqlDeleteExecute
}
