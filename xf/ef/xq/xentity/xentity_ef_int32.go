package xentity

import (
	"github.com/go-xe2/x/type/t"
	"github.com/go-xe2/x/xf/ef/xqi"
)

type TEFInt32 struct {
	*baseField
}

func newEFInt32(entity xqi.Entity, defineName string, attr []xqi.XqAttribute, annotations map[string]interface{}, params ...interface{}) interface{} {
	inst := &TEFInt32{}
	base := newBaseField(entity, defineName, attr, annotations, params, inst)
	inst.baseField = base
	return inst.Constructor(inst)
}

func (ef *TEFInt32) Value() interface{} {
	return ef.Int32()
}

func (ef *TEFInt32) Int32() int32 {
	return t.Int32(ef.baseField.Value())
}

func (ef *TEFInt32) MarshalJSON() ([]byte, error) {
	return []byte(t.String(ef.Int32())), nil
}

func (ef *TEFInt32) UnmarshalJSON(data []byte) error {
	ef.Set(t.Int32(string(data)))
	return nil
}

func (ef *TEFInt32) FieldType() xqi.FieldDataType {
	return xqi.FDTInt32
}

func (ef *TEFInt32) NewInstance(alias string, inherited ...interface{}) xqi.EntField {
	inst := &TEFInt32{}
	var this interface{} = inst
	if len(inherited) > 0 {
		if _, ok := inherited[0].(xqi.EntField); ok {
			this = inherited[0]
		}
	}
	inst.baseField = ef.baseField.NewInstance(alias, this).(*baseField)
	return inst
}
