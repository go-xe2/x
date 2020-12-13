package xentity

import (
	"github.com/go-xe2/x/type/t"
	"github.com/go-xe2/x/xf/ef/xqi"
)

type TEFUint8 struct {
	*baseField
}

func newEFUint8(entity xqi.Entity, defineName string, attr []xqi.XqAttribute, annotations map[string]interface{}, params ...interface{}) interface{} {
	inst := &TEFUint8{}
	base := newBaseField(entity, defineName, attr, annotations, params, inst)
	inst.baseField = base
	return inst.Constructor(inst)
}

func (ef *TEFUint8) Uint8() uint8 {
	return t.Uint8(ef.baseField.Value())
}

func (ef *TEFUint8) Value() interface{} {
	return ef.Uint8()
}

func (ef *TEFUint8) TryUint8() (uint8, bool) {
	if ef.IsOpen() {
		return ef.Uint8(), true
	}
	return 0, false
}

func (ef *TEFUint8) MarshalJSON() ([]byte, error) {
	return []byte(t.String(ef.Uint8())), nil
}

func (ef *TEFUint8) UnmarshalJSON(data []byte) error {
	ef.Set(t.Uint8(string(data)))
	return nil
}

func (ef *TEFUint8) FieldType() xqi.FieldDataType {
	return xqi.FDTUint8
}

func (ef *TEFUint8) NewInstance(alias string, inherited ...interface{}) xqi.EntField {
	inst := &TEFUint8{}
	var this interface{} = inst
	if len(inherited) > 0 {
		if _, ok := inherited[0].(xqi.EntField); ok {
			this = inherited[0]
		}
	}
	inst.baseField = ef.baseField.NewInstance(alias, this).(*baseField)
	return inst
}
