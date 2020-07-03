package xentity

import (
	"fmt"
	"github.com/go-xe2/x/core/exception"
	"github.com/go-xe2/x/type/t"
	"github.com/go-xe2/x/xf/ef/xdriveri"
	"github.com/go-xe2/x/xf/ef/xq/xqcomm"
	"github.com/go-xe2/x/xf/ef/xqi"
)

type TEFExprField struct {
	*baseField
	expr func(ent xqi.Entity) xqi.SqlField
}

var _ xqi.EFExpr = (*TEFExprField)(nil)

func newEFExprField(entity xqi.Entity, defineName string, attr []xqi.XqAttribute, annotations map[string]interface{}, params ...interface{}) interface{} {
	inst := &TEFExprField{}
	base := newBaseField(entity, defineName, attr, annotations, params, inst)
	inst.baseField = base
	return inst.Constructor(inst)
}

func (ef *TEFExprField) Supper() xqi.EntField {
	return ef.baseField
}

func (ef *TEFExprField) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", t.String(ef.Value()))), nil
}

func (ef *TEFExprField) UnmarshalJSON(data []byte) error {
	ef.Set(t.Bytes(string(data)))
	return nil
}

func (ef *TEFExprField) FieldType() xqi.FieldDataType {
	return xqi.FDTUnknown
}

func (ef *TEFExprField) Expression(expr func(ent xqi.Entity) xqi.SqlField) {
	ef.expr = expr
}

func (ef *TEFExprField) IsPrimary() bool {
	return false
}

func (ef *TEFExprField) IsForeign() bool {
	return false
}

func (ef *TEFExprField) NewInstance(alias string, inherited ...interface{}) xqi.EntField {
	inst := &TEFExprField{expr: ef.expr}
	var this interface{} = inst
	if len(inherited) > 0 {
		if _, ok := inherited[0].(xqi.EntField); ok {
			this = inherited[0]
		}
	}
	inst.baseField = ef.baseField.NewInstance(alias, this).(*baseField)
	return inst
}

func (ef *TEFExprField) TokenType() xqi.SqlTokenType {
	return xqi.SqlExpressTokenType
}

func (ef *TEFExprField) Compile(builder xdriveri.DbDriverSqlBuilder, cxt xqi.SqlCompileContext, unPrepare ...bool) xqi.SqlToken {
	if ef.expr == nil {
		panic(exception.Newf("字段%s未定义表达式", ef.DefineName()))
	}
	fd := ef.expr(ef.entity)
	if fd == nil {
		panic(exception.Newf("字段%s定义的表达式未正常返回字段", ef.DefineName()))
	}
	state := cxt.State()
	cxt.PushState(xqi.SCPQrSelectExprFieldState)
	defer cxt.PopState()
	if tk := fd.Alias(ef.AliasName()).Compile(builder, cxt, unPrepare...); tk == nil || tk.TType() == xqi.SqlEmptyTokenType {
		panic(exception.Newf("字段%s定义的表达式错误", ef.DefineName()))
	} else {
		if state == xqi.SCPQrSelectWhereCondState || state == xqi.SCPQrSelectJoinCondState || state == xqi.SCPBuildUpdateWhereState ||
			state == xqi.SCPBuildUpdateWhereFromState {
			// 条件表达式不能有别名
			return xqcomm.NewSqlToken(tk.Val(), ef.TokenType())
		}
		alias := ef.defineName
		if ef.AliasName() != "" {
			alias = ef.AliasName()
		}
		s := fmt.Sprintf("(%s) %s", tk.Val(), alias)
		return xqcomm.NewSqlToken(s, ef.TokenType())
	}
}
