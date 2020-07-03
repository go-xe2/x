package xqcomm

import (
	. "github.com/go-xe2/x/xf/ef/xdriveri"
	. "github.com/go-xe2/x/xf/ef/xqi"
)

// 日期格式化函数
func SqlFunDateFormat(val interface{}, format string) SqlField {
	fun := NewSqlFunDateFormat(val, format)
	return NewSqlField(fun, "")
}

// 时间增加函数
func SqlFunDateAdd(val interface{}, interval int, datePart DatePart) SqlField {
	fun := NewSqlFunDateAdd(val, interval, datePart)
	return NewSqlField(fun, "")
}

// 时间减少函数
func SqlFunDateSub(val interface{}, interval int, datePart DatePart) SqlField {
	fun := NewSqlFunDateSub(val, interval, datePart)
	return NewSqlField(fun, "")
}

// 计算时间差值
func SqlFunDateDiff(val1, val2 interface{}, datePart DatePart) SqlField {
	fun := NewSqlFunDateDiff(val1, val2, datePart)
	return NewSqlField(fun, "")
}

// date转unix函数
func SqlFunDateToUnix(val interface{}) SqlField {
	fun := NewSqlFunDateToUnix(val)
	return NewSqlField(fun, "")
}

// unix转date函数
func SqlFunUnixToDate(val interface{}) SqlField {
	fun := NewSqlFunUnixToDate(val)
	return NewSqlField(fun, "")
}
