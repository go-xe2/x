package xqcomm

import (
	"fmt"
	. "github.com/go-xe2/x/xf/ef/xdriveri"
	. "github.com/go-xe2/x/xf/ef/xqi"
)

type TSqlCondition struct {
	lExpr SqlLogic
	logic SqlConditionLogic
	items []SqlLogic
}

var _ SqlCondition = &TSqlCondition{}

func NewSqlCondition(logic ...SqlConditionLogic) *TSqlCondition {
	l := SqlConditionAndLogic
	if len(logic) > 0 {
		l = logic[0]
	}
	return &TSqlCondition{
		lExpr: nil,
		logic: l,
		items: make([]SqlLogic, 0),
	}
}

func (sc *TSqlCondition) LExp() interface{} {
	return sc.lExpr
}

func (sc *TSqlCondition) RExp() interface{} {
	return sc.items
}

func (sc *TSqlCondition) Exp() interface{} {
	return sc
}

func (sc *TSqlCondition) TokenType() SqlTokenType {
	return SqlConditionTokenType
}

func (sc *TSqlCondition) Logic() SqlConditionLogic {
	return sc.logic
}

func (sc *TSqlCondition) Items() []SqlLogic {
	return sc.items
}

func (sc *TSqlCondition) And(exp ...SqlLogic) SqlCondition {
	result := NewSqlCondition(SqlConditionAndLogic)
	result.lExpr = sc.lExpr
	result.items = exp
	sc.lExpr = result
	return sc
}

func (sc *TSqlCondition) Or(exp ...SqlLogic) SqlCondition {
	result := NewSqlCondition(SqlConditionOrLogic)
	result.lExpr = sc.lExpr
	result.items = exp
	sc.lExpr = result
	return sc
}

func (sc *TSqlCondition) Xor(exp ...SqlLogic) SqlCondition {
	result := NewSqlCondition(SqlConditionXorLogic)
	result.lExpr = sc.lExpr
	result.items = exp
	sc.lExpr = result
	return sc
}

func (sc *TSqlCondition) This() interface{} {
	return sc
}

func (sc *TSqlCondition) Compile(builder DbDriverSqlBuilder, cxt SqlCompileContext, unPrepare ...bool) SqlToken {
	var lExp = ""
	if sc.lExpr != nil {
		lToken := sc.lExpr.Compile(builder, cxt, unPrepare...)
		lExp = lToken.Val()
	}
	rExp := ""
	for _, item := range sc.items {
		if item != nil {
			itemToken := item.Compile(builder, cxt, unPrepare...)
			if rExp != "" {
				rExp += fmt.Sprintf(" AND %s", itemToken.Val())
			} else {
				rExp = itemToken.Val()
			}
		}
	}
	var exp = ""
	var transRExpr = func(s string) string {
		if sc.logic == SqlConditionXorLogic {
			return fmt.Sprintf("NOT (%s)", s)
		} else {
			return s
		}
	}

	if lExp == "" && rExp != "" {
		exp = transRExpr(rExp)
	} else if lExp != "" && rExp == "" {
		exp = transRExpr(lExp)
	} else if lExp != "" && rExp != "" {
		exp = fmt.Sprintf("(%s %s (%s))", lExp, sc.logic.String(), rExp)
	}
	//
	//if rExp != "" {
	//	fmt.Println("======? rExp is not empty , rExp:", rExp)
	//	if sc.logic == SqlConditionXorLogic {
	//		exp = fmt.Sprintf("(%s) AND (NOT (%s))", rExp, lExp)
	//	} else {
	//		exp = fmt.Sprintf("(%s) %s (%s)", rExp, sc.logic.String(), lExp)
	//	}
	//} else {
	//	exp = lExp
	//}
	return NewSqlToken(exp, SqlConditionTokenType)
}
