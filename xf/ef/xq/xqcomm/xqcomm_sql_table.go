package xqcomm

import (
	"fmt"
	. "github.com/go-xe2/x/xf/ef/xdriveri"
	. "github.com/go-xe2/x/xf/ef/xqi"
)

type TSqlTable struct {
	tableName  string
	tableAlias string
	this       interface{}
}

var _ SqlTable = (*TSqlTable)(nil)
var _ SqlTableFields = (*TSqlTable)(nil)

func NewSqlTable(inherited interface{}, tableName string, alias ...string) *TSqlTable {
	s := ""
	if len(alias) > 0 {
		s = alias[0]
	}
	inst := &TSqlTable{
		tableName:  tableName,
		tableAlias: s,
	}
	if _, ok := inherited.(SqlTable); ok {
		inst.this = inherited
	} else {
		inst.this = inst
	}
	return inst
}

func (st *TSqlTable) TokenType() SqlTokenType {
	return SqlTableTokenType
}

func (st *TSqlTable) TableName() string {
	return st.tableName
}

func (st *TSqlTable) TableAlias() string {
	return st.tableAlias
}

func (st *TSqlTable) Alias(name string) SqlTable {
	result := NewSqlTable(st.tableName, name)
	return result
}

func (st *TSqlTable) This() interface{} {
	return st.this
}

func (st *TSqlTable) AllField() []SqlField {
	return []SqlField{}
}

func (st *TSqlTable) Field(index int) SqlField {
	return nil
}

func (st *TSqlTable) FieldByName(name string) SqlField {
	return nil
}

func (st *TSqlTable) SelField(field SqlField) SqlField {
	return field
}

func (st *TSqlTable) FieldCount() int {
	return 0
}

func (st *TSqlTable) makeQrName(cxt SqlCompileContext, builder DbDriverSqlBuilder, withAlias bool) string {
	s := builder.QuotesName(cxt.TablePrefix() + st.tableName)
	if withAlias && st.tableAlias != "" {
		s += " " + st.tableAlias
	}
	return s
}

func (st *TSqlTable) Compile(builder DbDriverSqlBuilder, cxt SqlCompileContext, unPrepare ...bool) SqlToken {
	state := cxt.State()
	var s = ""
	switch state {
	case SCPQrSelectFromState:
		// from 后的表名
		s = st.makeQrName(cxt, builder, true)
		break
	case SCPBuildInsertTableState:
		s = st.makeQrName(cxt, builder, false)
		break
	case SCPQrSelectJoinTableState:
		// join 后的表名
		s = st.makeQrName(cxt, builder, true)
		break
	case SCPBuildUpdateTableState:
		s = st.makeQrName(cxt, builder, false)
		break
	case SCPBuildUpdateTableWithFromState:
		s = st.makeQrName(cxt, builder, true)
		break
	case SCPBuildUpdateJoinTableState:
		s = st.makeQrName(cxt, builder, true)
		break
	case SCPBuildDeleteTableState:
		s = st.makeQrName(cxt, builder, true)
		break
	default:
		s = st.makeQrName(cxt, builder, false)
	}
	return NewSqlToken(s, st.TokenType())
}

func (st *TSqlTable) String() string {
	if st.tableAlias != "" {
		return fmt.Sprintf("table(%s as %s)", st.tableName, st.tableAlias)
	}
	return fmt.Sprintf("table(%s)", st.tableName)
}
