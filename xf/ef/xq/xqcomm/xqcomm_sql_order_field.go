package xqcomm

import (
	. "github.com/go-xe2/x/xf/ef/xdriveri"
	. "github.com/go-xe2/x/xf/ef/xqi"
)

type TSqlOrderField struct {
	field  SqlField
	direct SqlOrderDirect
}

var _ SqlOrderField = &TSqlOrderField{}

func NewSqlOrderField(field SqlField, direct SqlOrderDirect) SqlOrderField {
	return &TSqlOrderField{
		field:  field,
		direct: direct,
	}
}

func (sof *TSqlOrderField) Exp() interface{} {
	return sof
}

func (sof *TSqlOrderField) TokenType() SqlTokenType {
	return SqlOrderFieldTokenType
}

func (sof *TSqlOrderField) Field() SqlField {
	return sof.field
}

func (sof *TSqlOrderField) OrderDir() SqlOrderDirect {
	return sof.direct
}

func (sof *TSqlOrderField) This() interface{} {
	return sof
}

func (sof *TSqlOrderField) Compile(builder DbDriverSqlBuilder, cxt SqlCompileContext, unPrepare ...bool) SqlToken {
	fdToken := sof.field.Compile(builder, cxt, unPrepare...)
	s := fdToken.Val()
	s += sof.direct.Exp()
	return NewSqlToken(s, SqlOrderFieldTokenType)
}
