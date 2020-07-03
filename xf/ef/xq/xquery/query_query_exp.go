package xquery

import (
	"github.com/go-xe2/x/core/exception"
	"github.com/go-xe2/x/xf/ef/xdriveri"
	"github.com/go-xe2/x/xf/ef/xq/xqcomm"
	. "github.com/go-xe2/x/xf/ef/xqi"
)

type tQueryExp struct {
	db             Database
	com            *xqcomm.TSqlQuery
	fields         func(tables SqlTables) []SqlField
	bQueryAllField bool
	exprAlias      string
}

var _ Query = &tQueryExp{}

func NewQueryExp(db Database) QuerySelect {
	return &tQueryExp{
		db:             db,
		com:            nil,
		bQueryAllField: false,
		exprAlias:      "",
	}
}

func (qe *tQueryExp) DB() Database {
	return qe.db
}

func (qe *tQueryExp) raiseUnSetMainTable() {
	panic(exception.NewText("未选择查询表主"))
}

func (qe *tQueryExp) checkMainTableSet() {
	if qe.com == nil {
		qe.raiseUnSetMainTable()
	}
}

func (qe *tQueryExp) TokenType() SqlTokenType {
	return SqlQueryTokenType
}

func (qe *tQueryExp) Info() QueryInfo {
	qe.checkMainTableSet()
	if qe.com == nil {
		return nil
	}
	return qe
}

// 查询指定字段
func (qe *tQueryExp) Fields(fields func(tables SqlTables) []SqlField) QueryFrom {
	qe.fields = fields
	qe.bQueryAllField = false
	return qe
}

func (qe *tQueryExp) This() interface{} {
	return qe
}

// 查询所有字段
func (qe *tQueryExp) All() QueryFrom {
	qe.bQueryAllField = true
	qe.fields = func(tables SqlTables) []SqlField {
		result := make([]SqlField, 0)
		allTables := tables.All()
		for _, table := range allTables {
			fields := table.AllField()
			result = append(result, fields...)
		}
		return result
	}
	return qe
}

func (qe *tQueryExp) Alias(alias string) SqlTable {
	qe.checkMainTableSet()
	// 生成查询sql脚表本
	return qe.com.Alias(alias)
}

func (qe *tQueryExp) Compile(builder xdriveri.DbDriverSqlBuilder, cxt SqlCompileContext, unPrepare ...bool) SqlToken {
	qe.checkMainTableSet()
	return qe.com.Compile(builder, cxt, unPrepare...)
}
