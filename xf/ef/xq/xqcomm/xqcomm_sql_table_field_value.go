package xqcomm

import (
	"fmt"
	"github.com/go-xe2/x/type/t"
	"github.com/go-xe2/x/xf/ef/xqi"
)

type tSqlFieldValue struct {
	field xqi.SqlTableField
	val   interface{}
	// 是否是主键,为-1时表达未检查,0:非主键，1:主键
	isPrimaryKey int
}

var _ xqi.FieldValue = (*tSqlFieldValue)(nil)

// 创建字段赋值表达式
func NewFieldValue(field xqi.SqlTableField, val interface{}) xqi.FieldValue {
	return &tSqlFieldValue{
		field:        field,
		val:          val,
		isPrimaryKey: -1,
	}
}

func (tfa *tSqlFieldValue) Field() xqi.SqlTableField {
	return tfa.field
}

func (tfa *tSqlFieldValue) Value() interface{} {
	return tfa.val
}

func (tfa *tSqlFieldValue) IsPrimaryKey() bool {
	if tfa.isPrimaryKey == -1 {
		if tfa.field == nil {
			tfa.isPrimaryKey = 0
		} else {
			if efd, ok := tfa.field.This().(xqi.EntField); ok {
				if efd.IsPrimary() {
					tfa.isPrimaryKey = 1
				} else {
					tfa.isPrimaryKey = 0
				}
			} else {
				tfa.isPrimaryKey = 0
			}
		}
	}
	return tfa.isPrimaryKey == 1
}

func (tfa *tSqlFieldValue) String() string {
	szFieldName := ""
	szValue := ""
	if tfa.field != nil {
		szFieldName = tfa.field.FieldName()
	}
	if tfa.val != nil {
		if sr, ok := tfa.val.(fmt.Stringer); ok {
			szValue = sr.String()
		} else {
			szValue = t.String(tfa.val)
		}
	}
	return fmt.Sprintf("{ %s=%s }", szFieldName, szValue)
}
