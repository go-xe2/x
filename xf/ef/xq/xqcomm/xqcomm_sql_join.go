package xqcomm

import . "github.com/go-xe2/x/xf/ef/xqi"

type TSqlJoin struct {
	joinType  SqlJoinType
	joinTable SqlTable
	condition SqlCondition
}

var _ SqlJoin = &TSqlJoin{}

func NewSqlJoin(joinType SqlJoinType, table SqlTable, on SqlCondition) *TSqlJoin {
	return &TSqlJoin{
		joinType:  joinType,
		joinTable: table,
		condition: on,
	}
}

// 联查方式
func (sj *TSqlJoin) JoinType() SqlJoinType {
	return sj.joinType
}

// 连接的表
func (sj *TSqlJoin) JoinTable() SqlTable {
	return sj.joinTable
}

func (sj *TSqlJoin) OnCondition() SqlCondition {
	return sj.condition
}
