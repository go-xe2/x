package xqi

import "io"

type DatasetRow interface {
	RowIndex() int
	Field(index int) DSField
	FieldByName(name string) DSField
	Values() []interface{}
	Names() []string
	// 删除当前行
	Delete()
	ToMap() map[string]interface{}
	WriteJson(writer io.Writer) (n int, err error)
}
