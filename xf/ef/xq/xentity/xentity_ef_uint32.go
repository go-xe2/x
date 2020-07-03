package xentity

import (
	"github.com/go-xe2/x/type/t"
	"github.com/go-xe2/x/xf/ef/xqi"
)

type TEFUint32 struct {
	*baseField
}

func newEFUint32(entity xqi.Entity, defineName string, attr []xqi.XqAttribute, annotations map[string]interface{}, params ...interface{}) interface{} {
	inst := &TEFUint32{}
	base := newBaseField(entity, defineName, attr, annotations, params, inst)
	inst.baseField = base
	return inst.Constructor(inst)
}

func (ef *TEFUint32) Uint32() uint32 {
	return t.Uint32(ef.baseField.Value())
}

func (ef *TEFUint32) Value() interface{} {
	return ef.Uint32()
}

func (ef *TEFUint32) MarshalJSON() ([]byte, error) {
	return []byte(t.String(ef.Uint32())), nil
}

func (ef *TEFUint32) UnmarshalJSON(data []byte) error {
	ef.Set(t.Uint32(string(data)))
	return nil
}

func (ef *TEFUint32) FieldType() xqi.FieldDataType {
	return xqi.FDTUint32
}

func (ef *TEFUint32) NewInstance(alias string, inherited ...interface{}) xqi.EntField {
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
