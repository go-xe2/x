package xentity

import (
	"github.com/go-xe2/x/type/t"
	"github.com/go-xe2/x/xf/ef/xqi"
)

type TEFBool struct {
	*baseField
}

func newEFBool(entity xqi.Entity, defineName string, attr []xqi.XqAttribute, annotations map[string]interface{}, params ...interface{}) interface{} {
	inst := &TEFBool{}
	base := newBaseField(entity, defineName, attr, annotations, params, inst)
	inst.baseField = base
	return inst.Constructor(inst)
}

func (ef *TEFBool) Supper() xqi.EntField {
	return ef.baseField
}

func (ef *TEFBool) Value() interface{} {
	return ef.Bool()
}

func (ef *TEFBool) Bool() bool {
	return t.Bool(ef.baseField.Value())
}

func (ef *TEFBool) TryBool() (bool, bool) {
	if ef.IsOpen() {
		return ef.Bool(), true
	}
	return false, false
}

func (ef *TEFBool) MarshalJSON() ([]byte, error) {
	return []byte(t.String(ef.Bool())), nil
}

func (ef *TEFBool) UnmarshalJSON(data []byte) error {
	ef.Set(t.Bool(t.String(data)))
	return nil
}

func (ef *TEFBool) FieldType() xqi.FieldDataType {
	return xqi.FDTBool
}

func (ef *TEFBool) NewInstance(alias string, inherited ...interface{}) xqi.EntField {
	inst := &TEFBool{}
	var this interface{} = inst
	if len(inherited) > 0 {
		if _, ok := inherited[0].(xqi.EntField); ok {
			this = inherited[0]
		}
	}
	inst.baseField = ef.baseField.NewInstance(alias, this).(*baseField)
	return inst
}
