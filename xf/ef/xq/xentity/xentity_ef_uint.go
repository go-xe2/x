package xentity

import (
	"github.com/go-xe2/x/type/t"
	"github.com/go-xe2/x/xf/ef/xqi"
)

type TEFUint struct {
	*baseField
}

func newEFUint(entity xqi.Entity, defineName string, attr []xqi.XqAttribute, annotations map[string]interface{}, params ...interface{}) interface{} {
	inst := &TEFUint{}
	base := newBaseField(entity, defineName, attr, annotations, params, inst)
	inst.baseField = base
	return inst.Constructor(inst)
}

func (ef *TEFUint) Value() interface{} {
	return ef.Uint()
}

func (ef *TEFUint) Uint() uint {
	return t.Uint(ef.baseField.Value())
}

func (ef *TEFUint) MarshalJSON() ([]byte, error) {
	return []byte(t.String(ef.Uint())), nil
}

func (ef *TEFUint) UnmarshalJSON(data []byte) error {
	ef.Set(t.Uint(string(data)))
	return nil
}

func (ef *TEFUint) FieldType() xqi.FieldDataType {
	return xqi.FDTUint
}

func (ef *TEFUint) NewInstance(alias string, inherited ...interface{}) xqi.EntField {
	inst := &TEFUint{}
	var this interface{} = inst
	if len(inherited) > 0 {
		if _, ok := inherited[0].(xqi.EntField); ok {
			this = inherited[0]
		}
	}
	inst.baseField = ef.baseField.NewInstance(alias, this).(*baseField)
	return inst
}
