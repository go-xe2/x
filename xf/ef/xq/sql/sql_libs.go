package sql

import (
	"github.com/go-xe2/x/xf/ef/xdriveri"
	"github.com/go-xe2/x/xf/ef/xq/xqcomm"
	. "github.com/go-xe2/x/xf/ef/xqi"
)

func Sum(expr interface{}) SqlField {
	return xqcomm.SqlFunSum(Static(expr))
}

func Max(expr interface{}) SqlField {
	return xqcomm.SqlFunMax(Static(expr))
}

func Min(expr interface{}) SqlField {
	return xqcomm.SqlFunMin(Static(expr))
}

func Avg(expr interface{}) SqlField {
	return xqcomm.SqlFunAvg(Static(expr))
}

func Count(expr interface{}) SqlField {
	return xqcomm.SqlFunCount(Static(expr))
}

func IfNull(expr interface{}, nullVal interface{}) SqlField {
	return xqcomm.SqlFunIfNull(Static(expr), Static(nullVal))
}

func If(expr interface{}, trueValue, falseValue interface{}) SqlField {
	return xqcomm.SqlFunIf(Static(expr), Static(trueValue), Static(falseValue))
}

func Cast(expr interface{}, asType DbType) SqlField {
	return xqcomm.SqlFunCast(Static(expr), asType)
}

func Case(expr interface{}, cases map[interface{}]interface{}, caseElse ...interface{}) SqlField {
	var caseExpr SqlCase = xqcomm.NewSqlFunCase(Static(expr))
	for k, v := range cases {
		caseExpr = caseExpr.When(Static(k), Static(v))
	}
	if len(caseElse) > 0 {
		return caseExpr.ElseEnd(Static(caseElse[0]))
	}
	return caseExpr.End()
}

func CaseThen(condition interface{}, caseTrue, caseFalse interface{}) SqlField {
	return xqcomm.NewSqlFunCase(Static(condition)).ThenElse(Static(caseTrue), Static(caseFalse))
}

func DateFormat(expr interface{}, format string) SqlField {
	return xqcomm.SqlFunDateFormat(Static(expr), format)
}

func DateDiff(date1, date2 interface{}, datePart xdriveri.DatePart) SqlField {
	return xqcomm.SqlFunDateDiff(Static(date1), Static(date2), datePart)
}

func DateAdd(date interface{}, interval int, datePart xdriveri.DatePart) SqlField {
	return xqcomm.SqlFunDateAdd(Static(date), interval, datePart)
}

func DateSub(date interface{}, interval int, datePart xdriveri.DatePart) SqlField {
	return xqcomm.SqlFunDateSub(Static(date), interval, datePart)
}

func DateToUnix(date interface{}) SqlField {
	return xqcomm.SqlFunDateToUnix(Static(date))
}

func UnixToDate(unix interface{}) SqlField {
	return xqcomm.SqlFunUnixToDate(Static(unix))
}

// 字符串截取
func Substring(str interface{}, from, len int) SqlField {
	return xqcomm.SqlFunSubstring(Static(str), from, len)
}

// 字符串连接
func Concat(str interface{}, str1 interface{}, more ...interface{}) SqlField {
	return xqcomm.SqlFunStrConcat(str, str1, more...)
}

// 加法运算
func OpPlus(v1, v2 interface{}) SqlField {
	return xqcomm.NewSqlField(xqcomm.NewSqlAriExp(Static(v1), Static(v2), SqlAriPlusType), "")
}

// 减法运算
func OpSub(v1, v2 interface{}) SqlField {
	return xqcomm.NewSqlField(xqcomm.NewSqlAriExp(Static(v1), Static(v2), SqlAriSubType), "")
}

// 乘法运算
func OpMul(v1, v2 interface{}) SqlField {
	return xqcomm.NewSqlField(xqcomm.NewSqlAriExp(Static(v1), Static(v2), SqlAriMulType), "")
}

// 除法运算
func OpDiv(v1, v2 interface{}) SqlField {
	return xqcomm.NewSqlField(xqcomm.NewSqlAriExp(Static(v1), Static(v2), SqlAriDivType), "")
}
