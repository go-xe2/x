package xentity

import (
	"github.com/go-xe2/x/type/t"
	"github.com/go-xe2/x/xf/ef/xqi"
)

type TEFUint16 struct {
	*baseField
}

func newEFUint16(entity xqi.Entity, defineName string, attr []xqi.XqAttribute, annotations map[string]interface{}, params ...interface{}) interface{} {
	inst := &TEFUint16{}
	base := newBaseField(entity, defineName, attr, annotations, params, inst)
	inst.baseField = base
	return inst.Constructor(inst)
}

func (ef *TEFUint16) Uint16() uint16 {
	return t.Uint16(ef.baseField.Value())
}

func (ef *TEFUint16) Value() interface{} {
	return ef.Uint16()
}

func (ef *TEFUint16) MarshalJSON() ([]byte, error) {
	return []byte(t.String(ef.Uint16())), nil
}

func (ef *TEFUint16) UnmarshalJSON(data []byte) error {
	ef.Set(t.Uint16(string(data)))
	return nil
}

func (ef *TEFUint16) FieldType() xqi.FieldDataType {
	return xqi.FDTUint16
}

func (ef *TEFUint16) NewInstance(alias string, inherited ...interface{}) xqi.EntField {
	inst := &TEFBinary{}
	var this interface{} = inst
	if len(inherited) > 0 {
		if _, ok := inherited[0].(xqi.EntField); ok {
			this = inherited[0]
		}
	}
	inst.baseField = ef.baseField.NewInstance(alias, this).(*baseField)
	return inst
}
