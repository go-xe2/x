package xqi

import (
	"github.com/go-xe2/x/encoding/xjson"
	"github.com/go-xe2/x/encoding/xxml"
)

type EntitySelectConvert interface {
	Page(index, size int) EntitySelectPage
	// 返回数据类型相关方法
	List() EntitySelectList
	// 生成sql脚本相关
	Build() EntitySelectBuilder
	First() EntitySelectSingle
}

type EntitySelectList interface {
	// 返回[]map[string]interface数据方法
	Rows() (data []map[string]interface{}, err error)
	// 返回xml字符串
	Xml() (data xxml.XmlStr, err error)
	// 返回json字符串
	Json() (data xjson.JsonStr, err error)
	// 返回数据集
	Dataset() (data Dataset, err error)
	// 返回slice数组
	Slice() (data [][]interface{}, err error)
	// 访问返回数据并构造返回数据
	Visit(visitor QueryBinderVisit) (data interface{}, err error)
	// 自定议绑定器绑定返回数据方法
	Bind(binder DbQueryBinder) (data interface{}, err error)
	Convert() EntitySelectConvert
}

type EntitySelectPage interface {
	// map数组列表
	Rows() (data []map[string]interface{}, info QueryPageInfo, err error)
	// xml 字符串
	Xml() (data xxml.XmlStr, info QueryPageInfo, err error)
	// json字符串
	Json() (data xjson.JsonStr, info QueryPageInfo, err error)
	// 返回数据集
	Dataset() (data Dataset, info QueryPageInfo, err error)
	// 返回数组列表
	Slice() (data [][]interface{}, info QueryPageInfo, err error)
	//  访问返回数据并构造返回数据
	Visit(visitor QueryBinderVisit) (data interface{}, info QueryPageInfo, err error)
	// 自定议绑定器绑定返回数据方法
	Bind(binder DbQueryBinder) (data interface{}, info QueryPageInfo, err error)

	Convert() EntitySelectConvert
}

type EntitySelectBuilder interface {
	// sql表达式
	Sql() (sql string, vars []interface{})
	Convert() EntitySelectConvert
}

type EntitySelectSingle interface {
	Map() (data map[string]interface{}, err error)
	Xml() (data xxml.XmlStr, err error)
	Json() (data xjson.JsonStr, err error)
	Dataset() (data Dataset, err error)
	Slice() (data []interface{}, err error)
	Visit(visitor QueryBinderVisit) (data interface{}, err error)
	Convert() EntitySelectConvert
}

type EntitySelect interface {
	Top(count int) EntitySelect
	Limit(rows int, offset ...int) EntitySelect

	Where(where ...SqlCondition) EntitySelect
	Order(fields ...SqlOrderField) EntitySelect
	Group(fields ...SqlField) EntitySelect
	Having(having ...SqlCondition) EntitySelect
	// 使用别名作为嵌套查询或更新的数据表源
	Alias(alias string) SqlTable
	// 返回分页数据
	Page(index, size int) EntitySelectPage
	// 返回数据类型相关方法
	List() EntitySelectList
	// 返回单条数据
	First() EntitySelectSingle
	// 生成sql脚本相关
	Build() EntitySelectBuilder
}
