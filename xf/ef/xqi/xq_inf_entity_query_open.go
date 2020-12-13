package xqi

type EntityQueryOpen interface {
	// 打开数据, 以自身字段绑定数据
	// 返回所有满足条件的数据
	All() (int, error)
	// 最多返回count条数据
	Top(count int) (int, error)

	Page(pageIndex int, pageSize int) (pageInfo QueryPageInfo, err error)

	// 查询条件
	Where(where ...SqlCondition) EntityQueryOpen
	// 查询分组
	Group(fields ...SqlField) EntityQueryOpen
	// 分组过滤条件
	Having(fields ...SqlCondition) EntityQueryOpen
	// 排序
	Order(fields ...SqlOrderField) EntityQueryOpen
	// 数据返回限制
	Limit(rows int, offset ...int) EntityQueryOpen
	// 获取生成的sql语句表达式
	Sql() (sql string, vars []interface{})
}
