package xentity

import "github.com/go-xe2/x/xf/ef/xqi"

func Fields(fields ...xqi.SqlField) []xqi.SqlField {
	return fields
}

func AllFields() []xqi.SqlField {
	return nil
}

func OrderFields(fields ...xqi.SqlOrderField) []xqi.SqlOrderField {
	return fields
}

func Values(fieldValues ...xqi.FieldValue) []xqi.FieldValue {
	return fieldValues
}
