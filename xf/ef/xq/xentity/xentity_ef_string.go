package xentity

import (
	"fmt"
	"github.com/go-xe2/x/type/t"
	"github.com/go-xe2/x/xf/ef/xqi"
)

type TEFString struct {
	*baseField
}

func newEFString(entity xqi.Entity, defineName string, attr []xqi.XqAttribute, annotations map[string]interface{}, params ...interface{}) interface{} {
	inst := &TEFString{}
	base := newBaseField(entity, defineName, attr, annotations, params, inst)
	inst.baseField = base
	return inst.Constructor(inst)
}

func (ef *TEFString) Str() string {
	return t.String(ef.baseField.Value())
}

func (ef *TEFString) Value() interface{} {
	return ef.Str()
}

func (ef *TEFString) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", ef.Str())), nil
}

func (ef *TEFString) UnmarshalJSON(data []byte) error {
	ef.Set(string(data))
	return nil
}

func (ef *TEFString) FieldType() xqi.FieldDataType {
	return xqi.FDTString
}

func (ef *TEFString) NewInstance(alias string, inherited ...interface{}) xqi.EntField {
	inst := &TEFString{}
	var this interface{} = inst
	if len(inherited) > 0 {
		if _, ok := inherited[0].(xqi.EntField); ok {
			this = inherited[0]
		}
	}
	inst.baseField = ef.baseField.NewInstance(alias, this).(*baseField)
	return inst
}
