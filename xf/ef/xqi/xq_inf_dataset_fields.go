package xqi

import "io"

type DSFields interface {
	Add(fieldName string, fieldType FieldDataType) DSFields
	AddDef(fieldDef *TQueryColDef) DSFields
	Field(index int) DSField
	FieldByName(name string) DSField
	Clear() DSFields
	WriteJson(writer io.Writer) (n int, err error)
}
