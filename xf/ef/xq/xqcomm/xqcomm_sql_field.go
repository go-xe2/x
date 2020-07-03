package xqcomm

import (
	"fmt"
	"github.com/go-xe2/x/xf/ef/xdbUtil"
	. "github.com/go-xe2/x/xf/ef/xdriveri"
	. "github.com/go-xe2/x/xf/ef/xqi"
)

type TSqlField struct {
	val   interface{}
	alias string
	this  interface{}
}

var _ SqlField = &TSqlField{}

func NewSqlField(val interface{}, alias string, inherited ...interface{}) *TSqlField {
	inst := &TSqlField{
		val:   val,
		alias: alias,
		this:  nil,
	}
	inst.this = inst
	if len(inherited) > 0 {
		if _, ok := inherited[0].(SqlField); ok {
			inst.this = inherited[0]
		}
	}
	return inst
}

func (sf *TSqlField) Exp() interface{} {
	return sf.val
}

func (sf TSqlField) TokenType() SqlTokenType {
	v := sf.val
	if v == nil {
		return SqlEmptyTokenType
	} else if c, ok := v.(SqlCompiler); ok {
		return c.TokenType()
	} else {
		return SqlVarTokenType
	}
}

func (sf *TSqlField) AliasName() string {
	return sf.alias
}

func (sf *TSqlField) Alias(alias string) SqlField {
	return NewSqlField(sf.val, alias)
}

func (sf *TSqlField) This() interface{} {
	return sf.this
}

func (sf *TSqlField) String() string {
	return sf.alias
}

func (sf *TSqlField) Compile(builder DbDriverSqlBuilder, cxt SqlCompileContext, unPrepare ...bool) SqlToken {
	var exp = ""
	result := NewSqlToken("", SqlExpressTokenType)
	state := cxt.State()
	bExpCp := false
	if SCPQrSelectGroupFieldState == state && sf.alias != "" {
		return result.SetVal(sf.alias)
	}
	if c, ok := sf.val.(SqlCompiler); ok {
		if tk := c.Compile(builder, cxt, unPrepare...); tk != nil && tk.TType() != SqlEmptyTokenType {
			exp = tk.Val()
			bExpCp = tk.TType() == SqlExpressTokenType || tk.TType() == SqlFunExpressTokenType
		}
	} else {
		prepare := true
		if len(unPrepare) > 0 {
			prepare = !unPrepare[0]
		}
		switch state {
		case SCPQrSelectJoinCondState, SCPQrSelectJoinItemState, SCPQrSelectWhereCondState, SCPQrSelectHavingCondState:
			if prepare {
				vn := cxt.MakeParamId()
				cxt.AddParam(vn, sf.val)
				result.AddParam(vn, sf.val)
				exp = builder.PlaceHolder(vn)
			} else {
				exp = builder.MakeRealValue(sf.val)
			}
			break
		default:
			exp = builder.MakeRealValue(sf.val)
		}
	}
	if state == SCPQrSelectFieldsState || state == SCPQrSelectFieldState {
		if sf.alias != "" {
			exp = xdbUtil.IfThen(bExpCp, fmt.Sprintf("%s AS %s", exp, sf.alias), fmt.Sprintf("%s %s", exp, sf.alias)).(string)
		}
	}
	return result.SetVal(exp)
}
