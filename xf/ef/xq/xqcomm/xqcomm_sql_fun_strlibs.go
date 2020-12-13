package xqcomm

import . "github.com/go-xe2/x/xf/ef/xqi"

func SqlFunSubstring(exp interface{}, from int, l ...int) SqlField {
	fun := NewSqlFunSubString(exp, from, l...)
	return NewSqlField(fun, "")
}

func SqlFunStrConcat(val1 interface{}, val2 interface{}, more ...interface{}) SqlField {
	fun := NewSqlFunConcat(val1, val2, more...)
	return NewSqlField(fun, "")
}

func SqlFunFromBase64(val1 interface{}) SqlField {
	fun := NewSqlFunFromBase64(val1)
	return NewSqlField(fun, "")
}

func SqlFunToBase64(val1 interface{}) SqlField {
	fun := NewSqlFunToBase64(val1)
	return NewSqlField(fun, "")
}
