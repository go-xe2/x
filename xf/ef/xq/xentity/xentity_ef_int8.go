package xentity

import (
	"github.com/go-xe2/x/type/t"
	"github.com/go-xe2/x/xf/ef/xqi"
)

type TEFInt8 struct {
	*baseField
}

func newEFInt8(entity xqi.Entity, defineName string, attr []xqi.XqAttribute, annotations map[string]interface{}, params ...interface{}) interface{} {
	inst := &TEFInt8{}
	base := newBaseField(entity, defineName, attr, annotations, params, inst)
	inst.baseField = base
	return inst.Constructor(inst)
}

func (ef *TEFInt8) Value() interface{} {
	return ef.Int8()
}

func (ef *TEFInt8) Int8() int8 {
	return t.Int8(ef.baseField.Value())
}

func (ef *TEFInt8) MarshalJSON() ([]byte, error) {
	return []byte(t.String(ef.Int8())), nil
}

func (ef *TEFInt8) UnmarshalJSON(data []byte) error {
	ef.Set(t.Int8(string(data)))
	return nil
}

func (ef *TEFInt8) FieldType() xqi.FieldDataType {
	return xqi.FDTInt8
}

func (ef *TEFInt8) NewInstance(alias string, inherited ...interface{}) xqi.EntField {
	inst := &TEFInt8{}
	var this interface{} = inst
	if len(inherited) > 0 {
		if _, ok := inherited[0].(xqi.EntField); ok {
			this = inherited[0]
		}
	}
	inst.baseField = ef.baseField.NewInstance(alias, this).(*baseField)
	return inst
}
