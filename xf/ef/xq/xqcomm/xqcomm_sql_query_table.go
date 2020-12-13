package xqcomm

import (
	"fmt"
	"github.com/go-xe2/x/core/exception"
	"github.com/go-xe2/x/xf/ef/xdriveri"
	. "github.com/go-xe2/x/xf/ef/xqi"
)

// sql查询脚本虚拟表
type tSqlQueryTable struct {
	this        interface{}
	fields      []SqlField
	fieldMaps   map[string]SqlField
	query       SqlQuery
	alias       string
	compileData SqlToken
}

var _ SqlTable = &tSqlQueryTable{}

func NewSqlQueryTable(query SqlQuery, alias ...string) SqlTable {
	var s = ""
	if len(alias) > 0 {
		s = alias[0]
	}
	inst := &tSqlQueryTable{
		query: query,
		alias: s,
	}
	inst.this = inst
	return inst
}

func (sqt *tSqlQueryTable) TokenType() SqlTokenType {
	return SqlTableTokenType
}

func (sqt *tSqlQueryTable) TableName() string {
	return sqt.alias
}

func (sqt *tSqlQueryTable) TableAlias() string {
	//return sqt.alias
	return sqt.alias
}

func (sqt *tSqlQueryTable) Alias(name string) SqlTable {
	return NewSqlQueryTable(sqt.query, name)
}

func (sqt *tSqlQueryTable) initFields() {
	if sqt.fields != nil {
		return
	}
	sfItems := sqt.srcFields()
	fdCount := len(sfItems)
	sqt.fields = make([]SqlField, fdCount)
	sqt.fieldMaps = make(map[string]SqlField)
	for i := 0; i < fdCount; i++ {
		srcFd := sfItems[i]
		sqt.fields[i] = NewSqlTableField(nil, sqt, srcFd.AliasName())
		sqt.fieldMaps[srcFd.AliasName()] = sqt.fields[i]
	}
}

func (sqt *tSqlQueryTable) AllField() []SqlField {
	sqt.initFields()
	return sqt.fields
}

func (sqt *tSqlQueryTable) FieldCount() int {
	sqt.initFields()
	return len(sqt.fields)
}

func (sqt *tSqlQueryTable) Field(index int) SqlField {
	sqt.initFields()
	if index >= 0 && index < sqt.FieldCount() {
		return sqt.fields[index]
	}
	panic(exception.Newf("查询表%s字段索引超出边界%d", sqt.TableName(), index))
}

func (sqt *tSqlQueryTable) FieldByName(name string) SqlField {
	sqt.initFields()
	if f, ok := sqt.fieldMaps[name]; ok {
		return f
	}
	panic(exception.Newf("查询表%s不存在字段%s", sqt.TableName(), name))
}

func (sqt *tSqlQueryTable) SelField(field SqlField) SqlField {
	if field == nil {
		return nil
	}
	fdName := ""
	if tf, ok := field.(SqlTableField); ok {
		fdName = tf.FieldName()
	} else {
		fdName = field.AliasName()
	}
	if field.AliasName() != "" {
		fdName = field.AliasName()
	}
	return sqt.FieldByName(fdName)
}

func (sqt *tSqlQueryTable) srcFields() []SqlField {
	return sqt.query.GetQueryFields()
}

func (sqt *tSqlQueryTable) This() interface{} {
	return sqt.this
}

func (sqt *tSqlQueryTable) String() string {
	if sqt.compileData != nil {
		return sqt.compileData.Val()
	}
	return fmt.Sprintf("queryTable(%s)", sqt.TableName())
}

func (sqt *tSqlQueryTable) Compile(builder xdriveri.DbDriverSqlBuilder, cxt SqlCompileContext, unPrepare ...bool) SqlToken {
	if sqt.query == nil {
		return EmptySqlToken
	}
	result := NewSqlToken("", SqlQueryTableTokenType)
	subExpr := ""
	if subTk := sqt.query.Compile(builder, cxt, unPrepare...); subTk == nil || subTk.TType() == SqlEmptyTokenType {
		return EmptySqlToken
	} else {
		subExpr = subTk.Val()
		pms := subTk.Params()
		for _, item := range pms {
			result.AddParam(item.Name(), item.Val())
		}
	}
	tableName := sqt.TableName()
	if tableName == "" {
		tableName = cxt.MakeIndentId()
	}
	state := cxt.State()
	expr := ""
	switch state {
	case SCPQrSelectWhereCondState, SCPQrSelectHavingCondState:
		expr = fmt.Sprintf("(%s)", subExpr)
		break
	default:
		expr = fmt.Sprintf("(%s) %s", subExpr, tableName)
	}
	sqt.compileData = result
	return result.SetVal(expr)
}
