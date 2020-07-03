package xqcomm

import . "github.com/go-xe2/x/xf/ef/xqi"

func SqlFunSubstring(exp interface{}, from int, l ...int) SqlField {
	fun := NewSqlFunSubString(exp, from, l...)
	return NewSqlField(fun, "")
}

func SqlFunStrConcat(val1 interface{}, val2 interface{}) SqlField {
	fun := NewSqlFunConcat(val1, val2)
	return NewSqlField(fun, "")
}
