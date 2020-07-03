package xclass

// class继承字段类
type classVTExtendField struct {
	// 字段所在序号
	fieldIndex []int
	// 字段所属类型
	extendType *classVT
	offset     uintptr
}
