package xqcomm

import (
	"fmt"
	"github.com/go-xe2/x/type/t"
	"github.com/go-xe2/x/xf/ef/xdbUtil"
	. "github.com/go-xe2/x/xf/ef/xdriveri"
	. "github.com/go-xe2/x/xf/ef/xqi"
	"reflect"
	"strings"
)

type TSqlConditionItem struct {
	lExp     interface{}
	rExp     interface{}
	comparer SqlCompareType
}

var _ SqlConditionItem = &TSqlConditionItem{}

func NewSqlConditionItem(lExp, rExp interface{}, compare SqlCompareType) *TSqlConditionItem {
	return &TSqlConditionItem{
		lExp:     lExp,
		rExp:     rExp,
		comparer: compare,
	}
}

func (sci *TSqlConditionItem) Exp() interface{} {
	return sci
}

func (sci *TSqlConditionItem) TokenType() SqlTokenType {
	return SqlConditionItemTokenType
}

func (sci *TSqlConditionItem) LExp() interface{} {
	return sci.lExp
}

func (sci *TSqlConditionItem) RExp() interface{} {
	return sci.rExp
}

func (sci *TSqlConditionItem) This() interface{} {
	return sci
}

func (sci *TSqlConditionItem) Comparer() SqlCompareType {
	return sci.comparer
}

var compareType2LogicTypes = map[SqlCompareType]SqlConditionLogic{
	// ==
	SqlCompareEQType: SqlConditionEqLogic,
	// <>
	SqlCompareNEQType: SqlConditionNeqLogic,
	// >
	SqlCompareGTType: SqlConditionGtLogic,
	// >=
	SqlCompareGTEType: SqlConditionGteLogic,
	// <
	SqlCompareLTType: SqlConditionLtLogic,
	// <=
	SqlCompareLTEType: SqlConditionLteLogic,
	// like
	SqlCompareLKType: SqlConditionLkLogic,
	// not like
	SqlCompareNLKType: SqlConditionNlkLogic,
	// in
	SqlCompareINType: SqlConditionInLogic,
	// not in
	SqlCompareNINType: SqlConditionNinLogic,
	// not
	SqlCompareNTType: SqlConditionInLogic,
}

func (sci *TSqlConditionItem) Logic() SqlConditionLogic {
	if v, ok := compareType2LogicTypes[sci.comparer]; ok {
		return v
	}
	return SqlConditionEqLogic
}

func (sci *TSqlConditionItem) Compile(builder DbDriverSqlBuilder, cxt SqlCompileContext, unPrepare ...bool) SqlToken {

	prepare := true
	if len(unPrepare) > 0 {
		prepare = !unPrepare[0]
	}

	if sci.lExp == nil && sci.rExp == nil {
		return EmptySqlToken
	}
	var szLeft, szRight = "", ""
	var bLeftCp, bRightCp = false, false

	result := NewSqlToken("", SqlConditionItemTokenType)
	if l, ok := sci.lExp.(SqlCompiler); ok {
		if tk := l.Compile(builder, cxt, unPrepare...); tk != nil && tk.TType() != SqlEmptyTokenType {
			bLeftCp = tk.TType() == SqlExpressTokenType || tk.TType() == SqlFunExpressTokenType
			szLeft = tk.Val()
		}
	} else {
		if prepare {
			vn := cxt.MakeParamId()
			cxt.AddParam(vn, sci.lExp)
			result.AddParam(vn, sci.lExp)
			szLeft = builder.PlaceHolder(vn)
		} else {
			szLeft = builder.MakeRealValue(sci.lExp)
		}
	}

	if r, ok := sci.rExp.(SqlCompiler); ok {
		if tk := r.Compile(builder, cxt, unPrepare...); tk != nil {
			szRight = tk.Val()
			bRightCp = tk.TType() == SqlExpressTokenType || tk.TType() == SqlFunExpressTokenType
		}
	} else {
		if prepare {
			if sci.rExp == nil {
				szRight = " NULL "
			} else {
				if sci.comparer == SqlCompareINType || sci.comparer == SqlCompareNINType {
					// in 操作需要处理数组
					var vt = reflect.TypeOf(sci.rExp)
					if vt.Kind() == reflect.Ptr {
						vt = vt.Elem()
					}
					if vt.Kind() == reflect.Slice {
						slItems := t.New(sci.rExp).SliceAny()
						count := len(slItems)
						params := make([]string, count)
						for i := 0; i < count; i++ {
							item := slItems[i]
							vn := cxt.MakeParamId()
							cxt.AddParam(vn, item)
							result.AddParam(vn, item)
							params[i] = builder.PlaceHolder(vn)
						}
						szRight = "(" + strings.Join(params, ",") + ")"
					} else {
						vn := cxt.MakeParamId()
						cxt.AddParam(vn, sci.rExp)
						result.AddParam(vn, sci.rExp)
						szRight = builder.PlaceHolder(vn)
					}
				} else {
					vn := cxt.MakeParamId()
					cxt.AddParam(vn, sci.rExp)
					result.AddParam(vn, sci.rExp)
					szRight = builder.PlaceHolder(vn)
				}
			}
		} else {
			szRight = builder.MakeRealValue(sci.rExp)
		}
	}
	var exp = ""
	if sci.comparer == SqlCompareNTType {
		if szLeft == "" && szRight == "" {
			return EmptySqlToken
		}
		exp = "NOT " + xdbUtil.IfThen(szLeft != "", szLeft, szRight).(string)
	} else {
		if szLeft != "" && szRight == "" {
			exp = xdbUtil.IfThen(sci.comparer == SqlCompareNTType, fmt.Sprintf("(%s %s)", sci.comparer.Val(), szLeft), szLeft).(string)
		} else if szLeft != "" && szRight != "" {
			szLeft = xdbUtil.IfThen(bLeftCp, fmt.Sprintf("(%s)", szLeft), szLeft).(string)
			szRight = xdbUtil.IfThen(bRightCp, fmt.Sprintf("(%s)", szRight), szRight).(string)
			exp = fmt.Sprintf("%s %s %s", szLeft, sci.comparer.Val(), szRight)
		} else if szLeft == "" && szRight != "" {
			exp = xdbUtil.IfThen(sci.comparer == SqlCompareNTType, fmt.Sprintf("(%s %s)", sci.comparer.Val(), szRight), szRight).(string)
		} else {
			return EmptySqlToken
		}
	}
	result.SetVal(exp)
	return result
}
