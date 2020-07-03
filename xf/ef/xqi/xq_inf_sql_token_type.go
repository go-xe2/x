package xqi

type SqlTokenType int

const (
	// 空节点
	SqlEmptyTokenType SqlTokenType = iota
	// 常量
	SqlStaticTokenType
	// sql变量
	SqlVarTokenType
	// 字段类型
	SqlFieldTokenType
	// 字段赋值表达式
	SqlFieldAssignTokenType
	// 数据表节点
	SqlTableTokenType
	// sql查询表达式表
	SqlQueryTableTokenType
	// 序列字段
	SqlOrderFieldTokenType
	// 算术运算表达式
	SqlArithmeticTokenType
	// 关系表达式
	SqlConditionTokenType
	// 关系表达式项
	SqlConditionItemTokenType
	// 函数表达式
	SqlFunExpressTokenType
	// 聚合函数表达式
	SqlAggregateFunTokenType
	// 表达式
	SqlExpressTokenType
	// 查询语句节点
	SqlQueryTokenType
	// 插入语句
	SqlInsertTokenType
	// 更新语句
	SqlUpdateTokenType
)

var SqlTokenTypeNames = map[SqlTokenType]string{
	SqlStaticTokenType:        "staticToken",
	SqlVarTokenType:           "varToken",
	SqlFieldTokenType:         "fieldToken",
	SqlFieldAssignTokenType:   "fieldAssignToken",
	SqlTableTokenType:         "tableToken",
	SqlOrderFieldTokenType:    "orderFieldToken",
	SqlArithmeticTokenType:    "arithmeticToken",
	SqlConditionTokenType:     "conditionToken",
	SqlConditionItemTokenType: "conditionItemToken",
	SqlFunExpressTokenType:    "funExpressToken",
	SqlAggregateFunTokenType:  "aggregateFunToken",
	SqlExpressTokenType:       "sqlExpressToken",
	SqlQueryTokenType:         "queryToken",
	SqlInsertTokenType:        "insertToken",
	SqlUpdateTokenType:        "updateToken",
}

func (stt SqlTokenType) Name() string {
	if s, ok := SqlTokenTypeNames[stt]; ok {
		return s
	}
	return "unknown token type"
}
