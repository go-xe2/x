package xdriveri

type DbDriverSqlBuilder interface {
	Driver() DbDriver
	// 生成数据库表名、字段名等，防止表名或字段名与数据库关键字相同时出错
	QuotesName(name string) string
	// 生成sql脚本的实参值，根据数据类型过滤特殊字符后返回，以防止非参数化查询时sql注入攻击
	MakeRealValue(val interface{}) string
	// 生成数据类型定义
	// ddType：数据类型
	// size: 数据大小, 传0时使用默认大小
	// decimal: 小数点位数
	MakeDataDefine(ddType DbDataType, size int, decimal ...int) string
	// 查询占位符
	PlaceHolder(varName string) string
	MakeQueryParams(vars []SqlParam) []interface{}

	// 生成select脚本
	BuildQuery(table, fields, joins, where, group, having, order string, rows, offset int) string
	// 生成更新update脚本
	BuildUpdate(table, fields, joins, where string) string
	// 生成插入脚本
	BuildInsert(table, fields, values, fromTable string) string
	// 生成创建数据库表脚本
	BuildTableCreate(tableDefine DbTableDefine) string

	DbDriverSqlBuilderStrLibs
	DbDriverSqlBuilderDateLibs
	DbDriverSqlBuildFunLibs
}

type DbDriverSqlBuildFunLibs interface {
	// 数据库IfNull映射
	IfNull(v1, v2 string) string
	If(exp string, v1, v2 string) string
	ISNULL(exp string) string
	CAST(exp string, ddType DbDataType, size int, decimal ...int) string
	Case(exp string, whenThen [][]string, elseValue ...string) string
}

// 数据库字符串操作映射函数接口
type DbDriverSqlBuilderStrLibs interface {
	// 数据库subString函数映射
	Substring(field string, from int, len int) string
	Concat(field string, v string) string
}

// 数据库日期函数映射接口
type DbDriverSqlBuilderDateLibs interface {
	// 数据库时间增加函数映射
	DateAdd(field string, interval int, part DatePart) string
	// 数据库时间减少函数映射
	DateSub(field string, interval int, part DatePart) string
	// 格式化时间函数映射, field为DbField可以换转成日期的值
	DateFormat(field string, format string) string
	// 计算时间差函数映射
	DateDiff(field1 string, field2 string, part DatePart) string
	// 时间转时间戳函数映射
	DateToUnix(field string) string
	// 时间戳转时间函数映射，field为DbField或实参变量
	UnixToDate(field string) string
}
