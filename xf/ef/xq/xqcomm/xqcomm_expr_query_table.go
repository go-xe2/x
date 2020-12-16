package xqcomm

import (
	"fmt"
	"github.com/go-xe2/x/xf/ef/xdriveri"
	"github.com/go-xe2/x/xf/ef/xqi"
)

type TExpressQueryTable struct {
	alias   string
	express string
}

var _ xqi.SqlTable = (*TExpressQueryTable)(nil)

// NewExpressQuery 创建静态sql查询表达式表
func NewExpressQuery(querySql string, alias string) xqi.SqlTable {
	return &TExpressQueryTable{
		alias:   alias,
		express: querySql,
	}
}

func (sq *TExpressQueryTable) TokenType() xqi.SqlTokenType {
	return xqi.SqlQueryTableTokenType
}

func (sq *TExpressQueryTable) This() interface{} {
	return sq
}

func (sq *TExpressQueryTable) TableName() string {
	return ""
}

func (sq *TExpressQueryTable) TableAlias() string {
	return sq.alias
}

func (sq *TExpressQueryTable) Alias(name string) xqi.SqlTable {
	sq.alias = name
	return sq
}

func (sq *TExpressQueryTable) AllField() []xqi.SqlField {
	return []xqi.SqlField{}
}

func (sq *TExpressQueryTable) Field(index int) xqi.SqlField {
	return nil
}

func (sq *TExpressQueryTable) FieldByName(name string) xqi.SqlField {
	return NewSqlField(sq.This(), name)
}

func (sq *TExpressQueryTable) SelField(field xqi.SqlField) xqi.SqlField {
	return field
}

func (sq *TExpressQueryTable) String() string {
	return fmt.Sprintf("(%s) %s", sq.express, sq.alias)
}

func (sq *TExpressQueryTable) FieldCount() int {
	return 0
}

func (sq *TExpressQueryTable) Compile(builder xdriveri.DbDriverSqlBuilder, cxt xqi.SqlCompileContext, unPrepare ...bool) xqi.SqlToken {
	token := NewSqlToken(sq.express, xqi.SqlQueryTableTokenType)
	return token
}
