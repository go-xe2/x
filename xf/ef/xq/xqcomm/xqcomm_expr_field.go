package xqcomm

import (
	"fmt"
	"github.com/go-xe2/x/xf/ef/xdriveri"
	"github.com/go-xe2/x/xf/ef/xqi"
)

type TExpressQueryField struct {
	express string
	alias   string
}

var _ xqi.SqlField = (*TExpressQueryField)(nil)

func NewExpressField(sqlExpress string, alias string) xqi.SqlField {
	return &TExpressQueryField{
		express: sqlExpress,
		alias:   alias,
	}
}

func (p *TExpressQueryField) Compile(builder xdriveri.DbDriverSqlBuilder, cxt xqi.SqlCompileContext, unPrepare ...bool) xqi.SqlToken {
	return NewSqlToken(fmt.Sprintf("%s", p.express), xqi.SqlFieldTokenType)
}

func (p *TExpressQueryField) TokenType() xqi.SqlTokenType {
	return xqi.SqlFieldTokenType
}

func (p *TExpressQueryField) This() interface{} {
	return p
}

func (p *TExpressQueryField) Exp() interface{} {
	return p.express
}

func (p *TExpressQueryField) AliasName() string {
	return p.alias
}

func (p *TExpressQueryField) Alias(alias string) xqi.SqlField {
	p.alias = alias
	return p
}

func (p *TExpressQueryField) Eq(val interface{}) xqi.SqlConditionItem {
	return NewSqlConditionItem(p.This(), val, xqi.SqlCompareEQType)
}

func (p *TExpressQueryField) Neq(val interface{}) xqi.SqlConditionItem {
	return NewSqlConditionItem(p.This(), val, xqi.SqlCompareNEQType)
}

func (p *TExpressQueryField) Gt(val interface{}) xqi.SqlConditionItem {
	return NewSqlConditionItem(p.This(), val, xqi.SqlCompareGTType)
}

func (p *TExpressQueryField) Gte(val interface{}) xqi.SqlConditionItem {
	return NewSqlConditionItem(p.This(), val, xqi.SqlCompareGTEType)
}

func (p *TExpressQueryField) Lt(val interface{}) xqi.SqlConditionItem {
	return NewSqlConditionItem(p.This(), val, xqi.SqlCompareLTType)
}

func (p *TExpressQueryField) Lte(val interface{}) xqi.SqlConditionItem {
	return NewSqlConditionItem(p.This(), val, xqi.SqlCompareLTEType)
}

func (p *TExpressQueryField) In(arr interface{}) xqi.SqlConditionItem {
	return NewSqlConditionItem(p.This(), arr, xqi.SqlCompareINType)
}

func (p *TExpressQueryField) NotIn(arr interface{}) xqi.SqlConditionItem {
	return NewSqlConditionItem(p.This(), arr, xqi.SqlCompareNINType)
}

func (p *TExpressQueryField) Like(val interface{}) xqi.SqlConditionItem {
	return NewSqlConditionItem(p.This(), val, xqi.SqlCompareLKType)
}

func (p *TExpressQueryField) NotLike(val interface{}) xqi.SqlConditionItem {
	return NewSqlConditionItem(p.This(), val, xqi.SqlCompareNLKType)
}

func (p *TExpressQueryField) Plus(val interface{}) xqi.SqlField {
	return NewSqlField(NewSqlAriExp(p.This(), val, xqi.SqlAriPlusType), "")
}

func (p *TExpressQueryField) Sub(val interface{}) xqi.SqlField {
	return NewSqlField(NewSqlAriExp(p.This(), val, xqi.SqlAriSubType), "")
}

func (p *TExpressQueryField) Mul(val interface{}) xqi.SqlField {
	return NewSqlField(NewSqlAriExp(p.This(), val, xqi.SqlAriMulType), "")
}

func (p *TExpressQueryField) Div(val interface{}) xqi.SqlField {
	return NewSqlField(NewSqlAriExp(p.This(), val, xqi.SqlAriDivType), "")
}

func (p *TExpressQueryField) Case() xqi.SqlCase {
	return NewSqlFunCase(p.This().(xqi.SqlField))
}

func (p *TExpressQueryField) CaseEq(v interface{}) xqi.SqlCaseThenElse {
	return NewSqlFunCase(p.Eq(v))
}

func (p *TExpressQueryField) CaseNeq(v interface{}) xqi.SqlCaseThenElse {
	return NewSqlFunCase(p.Neq(v))
}

func (p *TExpressQueryField) CaseGt(v interface{}) xqi.SqlCaseThenElse {
	return NewSqlFunCase(p.Gt(v))
}

func (p *TExpressQueryField) CaseGte(v interface{}) xqi.SqlCaseThenElse {
	return NewSqlFunCase(p.Gte(v))
}

func (p *TExpressQueryField) CaseLt(v interface{}) xqi.SqlCaseThenElse {
	return NewSqlFunCase(p.Lt(v))
}

func (p *TExpressQueryField) CaseLte(v interface{}) xqi.SqlCaseThenElse {
	return NewSqlFunCase(p.Lte(v))
}

func (p *TExpressQueryField) Cast(asType xqi.DbType) xqi.SqlField {
	return SqlFunCast(p.This(), asType)
}

func (p *TExpressQueryField) Count() xqi.SqlField {
	return SqlFunCount(p.This().(xqi.SqlField))
}

func (p *TExpressQueryField) Max() xqi.SqlField {
	return SqlFunMax(p.This().(xqi.SqlField))
}

func (p *TExpressQueryField) Min() xqi.SqlField {
	return SqlFunMin(p.This().(xqi.SqlField))
}

func (p *TExpressQueryField) Avg() xqi.SqlField {
	return SqlFunAvg(p.This().(xqi.SqlField))
}

func (p *TExpressQueryField) Sum() xqi.SqlField {
	return SqlFunSum(p.This().(xqi.SqlField))
}

func (p *TExpressQueryField) String() string {
	return fmt.Sprintf("(%s) %s", p.express, p.alias)
}

func (p *TExpressQueryField) Substr(from, l int) xqi.SqlField {
	return SqlFunSubstring(p.This(), from, l)
}

func (p *TExpressQueryField) FromBase64() xqi.SqlField {
	return SqlFunFromBase64(p.This())
}

func (p *TExpressQueryField) ToBase64() xqi.SqlField {
	return SqlFunToBase64(p.This())
}

func (p *TExpressQueryField) Concat(val interface{}, more ...interface{}) xqi.SqlField {
	return SqlFunStrConcat(p.This(), val, more...)
}

func (p *TExpressQueryField) SubEq(from, l int, eqVal interface{}) xqi.SqlConditionItem {
	return NewSqlConditionItem(p.Substr(from, l), eqVal, xqi.SqlCompareEQType)
}

func (p *TExpressQueryField) SubNeq(from, l int, neqVal interface{}) xqi.SqlConditionItem {
	return NewSqlConditionItem(p.Substr(from, l), neqVal, xqi.SqlCompareNEQType)
}

func (p *TExpressQueryField) SubIn(from, l int, arr interface{}) xqi.SqlConditionItem {
	return NewSqlConditionItem(p.Substr(from, l), arr, xqi.SqlCompareINType)
}

func (p *TExpressQueryField) SubNotIn(from, l int, arr interface{}) xqi.SqlConditionItem {
	return NewSqlConditionItem(p.Substr(from, l), arr, xqi.SqlCompareNINType)
}

func (p *TExpressQueryField) SubLike(from, l int, val interface{}) xqi.SqlConditionItem {
	return NewSqlConditionItem(p.Substr(from, l), val, xqi.SqlCompareLKType)
}

func (p *TExpressQueryField) SubNotLike(from, l int, val interface{}) xqi.SqlConditionItem {
	return NewSqlConditionItem(p.Substr(from, l), val, xqi.SqlCompareNLKType)
}

func (p *TExpressQueryField) Not() xqi.SqlConditionItem {
	return NewSqlConditionItem(p.This(), nil, xqi.SqlCompareNTType)
}

func (p *TExpressQueryField) ConcatEq(catVal interface{}, eqVal interface{}) xqi.SqlConditionItem {
	return NewSqlConditionItem(p.Concat(catVal), eqVal, xqi.SqlCompareEQType)
}

func (p *TExpressQueryField) ConcatNeq(catVal interface{}, neqVal interface{}) xqi.SqlConditionItem {
	return NewSqlConditionItem(p.Concat(catVal), neqVal, xqi.SqlCompareNEQType)
}

func (p *TExpressQueryField) ConcatIn(catVal interface{}, arr interface{}) xqi.SqlConditionItem {
	return NewSqlConditionItem(p.Concat(catVal), arr, xqi.SqlCompareINType)
}

func (p *TExpressQueryField) ConcatNotIn(catVal interface{}, arr interface{}) xqi.SqlConditionItem {
	return NewSqlConditionItem(p.Concat(catVal), arr, xqi.SqlCompareNINType)
}

func (p *TExpressQueryField) ConcatLike(catVal interface{}, val interface{}) xqi.SqlConditionItem {
	return NewSqlConditionItem(p.Concat(catVal), val, xqi.SqlCompareLKType)
}

func (p *TExpressQueryField) ConcatNotLike(catVal interface{}, val interface{}) xqi.SqlConditionItem {
	return NewSqlConditionItem(p.Concat(catVal), val, xqi.SqlCompareNLKType)
}

func (p *TExpressQueryField) Desc() xqi.SqlOrderField {
	return NewSqlOrderField(p.This().(xqi.SqlField), xqi.SqlOrderDescDirect)
}

func (p *TExpressQueryField) Asc() xqi.SqlOrderField {
	return NewSqlOrderField(p.This().(xqi.SqlField), xqi.SqlOrderAscDirect)
}
