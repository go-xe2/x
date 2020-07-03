package xdriveri

// 数据库表定义
type DbTableDefine interface {
	// 表名称
	TableName() string
	// 字段列表
	Fields() []DBFieldDefine
}
