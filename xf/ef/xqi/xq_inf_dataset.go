package xqi

type Dataset interface {
	IsOpen() bool
	Close() bool
	DSField(index int) DSField
	DSFieldByName(name string) DSField
	// 字段值
	FieldValue(index int) interface{}
	// 字段数
	FieldCount() int
	// 记录数
	RowCount() int
	// 第一行记录，无数据时返回false
	MoveFirst() bool
	// 最后一行记录，无数据时返回false
	MoveLast() bool
	// 下一行数据，无下一行时返回false
	Next() bool
	// 上一行数据，无上一行时返回false
	Prior() bool
	IsEmpty() bool
	IsEof() bool
	IsBof() bool
	ToMap() []map[string]interface{}
	MarshalJSON() (data []byte, err error)
	UnmarshalJSON(data []byte) error
}

type MemDataset interface {
	Dataset
	Fields() DSFields
	// 添加数据
	Append(values ...interface{}) Dataset
	// 删除当前行
	Delete() bool
	// 删除指定行
	DeleteByRow(row int) bool
	// 生成数据集字段
	CreateDataSet(fields ...*TQueryColDef) Dataset
	// 从map数据复制数据创建
	CreateFromSliceMap(maps []map[string]interface{}) Dataset
	// 获取指定行数据
	Row(rowIndex int) DatasetRow
	// 遍历所有记录
	Iterator(iter func(row DatasetRow) bool)
	Current() DatasetRow
	Clear() Dataset
}
