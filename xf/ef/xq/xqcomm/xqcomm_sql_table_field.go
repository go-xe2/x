package xqcomm

import (
	. "github.com/go-xe2/x/xf/ef/xdriveri"
	. "github.com/go-xe2/x/xf/ef/xqi"
)

type TSqlTableField struct {
	table      SqlTable
	fieldName  string
	fieldAlias string
	this       interface{}
}

var _ SqlTableField = &TSqlTableField{}
var _ SqlField = &TSqlTableField{}

func NewSqlTableField(inherited interface{}, table SqlTable, name string, alias ...string) *TSqlTableField {
	s := ""
	if len(alias) > 0 {
		s = alias[0]
	}
	inst := &TSqlTableField{
		table:      table,
		fieldName:  name,
		fieldAlias: s,
		this:       nil,
	}
	inst.this = inst
	if _, ok := inherited.(SqlTableField); ok {
		inst.this = inherited
	}
	return inst
}

func (stf *TSqlTableField) Constructor(instance interface{}, props ...interface{}) interface{} {
	table := stf.table
	if table != nil {
		if tb, ok := table.This().(SqlTableFields); ok {
			if f, ok := stf.This().(SqlTableField); ok {
				tb.AddField(f)
			}
		}
	}
	return instance
}

func (stf *TSqlTableField) Exp() interface{} {
	return stf
}

func (stf *TSqlTableField) TokenType() SqlTokenType {
	return SqlFieldTokenType
}

func (stf *TSqlTableField) Table() SqlTable {
	return stf.table
}

func (stf *TSqlTableField) FieldName() string {
	return stf.fieldName
}

func (stf *TSqlTableField) AliasName() string {
	if stf.fieldAlias != "" {
		return stf.fieldAlias
	}
	return stf.fieldName
}

func (stf *TSqlTableField) Alias(name string) SqlField {
	return NewSqlTableField(stf.This(), stf.table, stf.fieldName, name)
}

func (stf *TSqlTableField) This() interface{} {
	return stf.this
}

func (stf *TSqlTableField) makeQrName(builder DbDriverSqlBuilder, withTable, withAlias bool) string {
	s := ""
	if withTable && stf.table.TableAlias() != "" {
		s += stf.table.TableAlias() + "."
	}
	s += builder.QuotesName(stf.fieldName)
	if withAlias && stf.fieldAlias != "" {
		s += " " + stf.fieldAlias
	}
	return s
}

func (stf *TSqlTableField) String() string {
	return stf.fieldName
}

func (stf *TSqlTableField) Compile(builder DbDriverSqlBuilder, cxt SqlCompileContext, unPrepare ...bool) SqlToken {
	s := ""
	state := cxt.State()
	cxt.PushState(SCPQrSelectFieldState)
	defer cxt.PopState()
	switch state {
	case SCPQrSelectFieldsState:
		s = stf.makeQrName(builder, true, true)
		break
	case SCPQrSelectExprFieldState:
		s = stf.makeQrName(builder, true, false)
		break
	case SCPQrSelectWhereCondState:
		s = stf.makeQrName(builder, true, false)
		break
	case SCPBuildDeleteWhereState:
		s = stf.makeQrName(builder, true, false)
		break
	case SCPQrSelectJoinCondState:
		s = stf.makeQrName(builder, true, false)
		break
	case SCPBuildUpdateJoinOnState:
		s = stf.makeQrName(builder, true, false)
		break
	case SCPBuildUpdateWhereState:
		s = stf.makeQrName(builder, false, false)
		break
	case SCPBuildUpdateWhereFromState:
		s = stf.makeQrName(builder, true, false)
		break
	case SCPQrSelectJoinItemState:
		s = stf.makeQrName(builder, true, false)
		break
	case SCPQrSelectHavingCondState:
		s = stf.makeQrName(builder, true, false)
		break
	case SCPQrSelectGroupFieldState:
		s = stf.makeQrName(builder, true, false)
		break
	case SCPQrSelectOrderFieldState:
		s = stf.makeQrName(builder, true, false)
		break
	case SCPQrSelectParamNameState:
		s = stf.fieldName
		break
	case SCPQrSelectFunParamState:
		s = stf.makeQrName(builder, true, false)
	case SCPBuildInsertFieldValueState:
		s = stf.makeQrName(builder, true, true)
		break
	case SCPBuildUpdateFieldState:
		s = stf.makeQrName(builder, false, false)
		break
	case SCPBuildUpdateFieldWithFromState:
		s = stf.makeQrName(builder, true, false)
		break
	case SCPBuildUpdateFieldValueFromState:
		s = stf.makeQrName(builder, true, false)
		break
	case SCPBuildUpdateFieldValueState:
		s = stf.makeQrName(builder, false, false)
		break
	default:
		s = stf.makeQrName(builder, false, false)
	}
	cxt.UseTable(stf.Table())
	return NewSqlToken(s, stf.TokenType())
}
