package xqcomm

import (
	"fmt"
	. "github.com/go-xe2/x/xf/ef/xdriveri"
	. "github.com/go-xe2/x/xf/ef/xqi"
)

type TSqlFunAggregate struct {
	*TSqlFun
	exp SqlCompiler
}

var _ SqlFun = (*TSqlFunAggregate)(nil)

var aggregateFunSql = map[SqlFunId]string{
	SFCount: "COUNT(%s)",
	SFAvg:   "AVG(%s)",
	SFMax:   "MAX(%s)",
	SFMin:   "MIN(%s)",
	SFSum:   "SUM(%s)",
}

func newAggregate(id SqlFunId, exp SqlCompiler) *TSqlFunAggregate {
	inst := &TSqlFunAggregate{exp: exp}
	base := newSqlFun(id, inst)
	inst.TSqlFun = base
	return inst
}

func (agg *TSqlFunAggregate) TokenType() SqlTokenType {
	return SqlAggregateFunTokenType
}

func (agg *TSqlFunAggregate) Compile(builder DbDriverSqlBuilder, cxt SqlCompileContext, unPrepare ...bool) SqlToken {
	result := NewSqlToken("", SqlAggregateFunTokenType)
	if agg.exp == nil {
		return EmptySqlToken
	}
	if !agg.funId.IsAggregation() {
		// 不是聚合函数
		return EmptySqlToken
	}
	tk := agg.exp.Compile(builder, cxt, unPrepare...)
	result.SetVal(fmt.Sprintf(aggregateFunSql[agg.funId], tk.Val()))
	return result
}
