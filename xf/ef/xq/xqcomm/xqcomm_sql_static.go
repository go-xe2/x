package xqcomm

import (
	"github.com/go-xe2/x/xf/ef/xdriveri"
	"github.com/go-xe2/x/xf/ef/xqi"
)

type tSqlStaticExpr struct {
	val interface{}
}

var _ xqi.SqlStaticExpr = (*tSqlStaticExpr)(nil)

func NewSqlStatic(v interface{}) xqi.SqlStaticExpr {
	return &tSqlStaticExpr{
		val: v,
	}
}

func (v *tSqlStaticExpr) This() interface{} {
	return v
}

func (v *tSqlStaticExpr) TokenType() xqi.SqlTokenType {
	return xqi.SqlStaticTokenType
}

func (v *tSqlStaticExpr) Compile(builder xdriveri.DbDriverSqlBuilder, cxt xqi.SqlCompileContext, unPrepare ...bool) xqi.SqlToken {
	if c, ok := v.val.(xqi.SqlCompiler); ok {
		// 常量编译，不进行参数化处理:unPrepare = true
		if tk := c.Compile(builder, cxt, true); tk != nil && tk.TType() != xqi.SqlEmptyTokenType {
			return NewSqlToken(tk.Val(), v.TokenType())
		}
		return EmptySqlToken
	}
	s := builder.MakeRealValue(v.val)
	return NewSqlToken(s, v.TokenType())
}

func (v *tSqlStaticExpr) Val() interface{} {
	return v.val
}
