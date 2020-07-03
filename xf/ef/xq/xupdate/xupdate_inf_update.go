package xupdate

import "github.com/go-xe2/x/xf/ef/xqi"

type SqlUpdateInfo interface {
	DB() xqi.Database
}

type SqlUpdateExecute interface {
	SqlUpdateInfo
	Execute() (int, error)
}

type SqlUpdateWhere interface {
	SqlUpdateInfo
	Where(where xqi.SqlCondition) SqlUpdateExecute
}

type SqlUpdateJoin interface {
	SqlUpdateWhere
	Join(table xqi.SqlTable, on func(joinEnt xqi.SqlTable, leftTables xqi.SqlTables) xqi.SqlCondition) SqlUpdateJoin
	LeftJoin(table xqi.SqlTable, on func(joinEnt xqi.SqlTable, leftTables xqi.SqlTables) xqi.SqlCondition) SqlUpdateJoin
	RightJoin(table xqi.SqlTable, on func(joinEnt xqi.SqlTable, leftTables xqi.SqlTables) xqi.SqlCondition) SqlUpdateJoin
	CrossJoin(table xqi.SqlTable, on func(joinEnt xqi.SqlTable, leftTables xqi.SqlTables) xqi.SqlCondition) SqlUpdateJoin
}

type SqlUpdateSet interface {
	SqlUpdateWhere
	Join(table xqi.SqlTable, on func(joinEnt xqi.SqlTable, leftTables xqi.SqlTables) xqi.SqlCondition) SqlUpdateJoin
}

type SqlUpdate interface {
	SqlUpdateInfo
	Set(values ...xqi.FieldValue) SqlUpdateSet
}
