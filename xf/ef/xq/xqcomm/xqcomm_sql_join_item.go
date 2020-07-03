package xqcomm

import (
	. "github.com/go-xe2/x/xf/ef/xqi"
)

type TSqlJoinItem struct {
	joinType            SqlJoinType
	joinTable           SqlTable
	lazyGetConditionFun LazyGetJoinConditionFun
}

var _ SqlJoinItem = &TSqlJoinItem{}

func NewSqlJoinItem(joinType SqlJoinType, table SqlTable, on LazyGetJoinConditionFun) *TSqlJoinItem {
	return &TSqlJoinItem{
		joinType:            joinType,
		joinTable:           table,
		lazyGetConditionFun: on,
	}
}

// 联查方式
func (sji *TSqlJoinItem) JoinType() SqlJoinType {
	return sji.joinType
}

// 连接的表
func (sji *TSqlJoinItem) JoinTable() SqlTable {
	return sji.joinTable
}

// 连接条件表达式
func (sji *TSqlJoinItem) LazyConditionFn() LazyGetJoinConditionFun {
	return sji.lazyGetConditionFun
}
