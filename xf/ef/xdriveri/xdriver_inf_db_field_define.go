package xdriveri

type DBFieldDefine interface {
	// 字段名称
	FieldName() string
	// 数据类型
	Type() DbDataType
	// 字段大小
	Size() int
	// 小数点位数
	Decimal() int
	// 是否允许为空
	AllowNull() bool
	// 是否自动增长
	AutoIncrement() bool
	// 是否主键
	IsPrimary() bool
	// 是否外键
	IsForeign() bool
	Default() interface{}
}
