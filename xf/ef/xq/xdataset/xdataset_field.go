package xdataset

import (
	"fmt"
	"github.com/go-xe2/x/encoding/xjson"
	"github.com/go-xe2/x/type/t"
	"github.com/go-xe2/x/xf/ef/xqi"
	"io"
)

type tDsField struct {
	// 字段所属的数据集
	ds *tDataset
	// 字段列序号
	colIndex int
	rowIndex int
	// golang 数据类型
	fieldType xqi.FieldDataType
	// 字段名称
	name string
}

var _ xqi.DSField = (*tDsField)(nil)

func newDsField(ds *tDataset, fieldName string, fieldIndex int, fieldType xqi.FieldDataType) *tDsField {
	return &tDsField{
		ds:        ds,
		colIndex:  fieldIndex,
		name:      fieldName,
		fieldType: fieldType,
	}
}

func (df *tDsField) FieldName() string {
	return df.name
}

func (df *tDsField) FieldType() xqi.FieldDataType {
	return df.fieldType
}

func (df *tDsField) FieldIndex() int {
	return df.colIndex
}

func (df *tDsField) Value() interface{} {
	return df.ds.colData(df.colIndex, df.rowIndex)
}

func (df *tDsField) writeDefineJson(writer io.Writer) (n int, err error) {
	s := fmt.Sprintf(`{ "N": "%s", "T": "%s" }`, df.name, df.fieldType.Json())
	return writer.Write([]byte(s))
}

func (df *tDsField) writeDataJson(writer io.Writer) (n int, err error) {
	v := df.Value()
	s := t.String(xjson.JsonFormat(v))
	return writer.Write([]byte(s))
}
