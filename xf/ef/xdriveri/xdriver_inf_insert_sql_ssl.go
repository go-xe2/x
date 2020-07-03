package xdriveri

// 插入语句执行方法
type ExprInsertSqlExecute func(sql string, vars []interface{}) int

// 准备插入方法，如果返回false则取消插入
type ExprInsertSqlPrepare func(sql string) bool

// 生成insert sql脚本的中间数据包
type TExprInsertSqlSSL struct {
	// 数据源类型
	ValueMode ExprInsertValueType
	// 表名
	TableName string
	// 要插入的字段列表
	Fields []string
	// 要插入的数据或者是QuerySql的查询参数实参
	Values   [][]interface{}
	QuerySql string
	// 插入sql语句执行的方法，一次sql插入如果values数据过大可能调用几次execute, execute由具体的数据库执行
	Execute        ExprInsertSqlExecute
	ExecutePrepare ExprInsertSqlPrepare
}
