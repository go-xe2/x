package xentity

import (
	"github.com/go-xe2/x/type/t"
	"github.com/go-xe2/x/xf/ef/xqi"
)

type TEFByte struct {
	*baseField
}

func newEFByte(entity xqi.Entity, defineName string, attr []xqi.XqAttribute, annotations map[string]interface{}, params ...interface{}) interface{} {
	inst := &TEFByte{}
	base := newBaseField(entity, defineName, attr, annotations, params, inst)
	inst.baseField = base
	return inst.Constructor(inst)
}

func (ef *TEFByte) Supper() xqi.EntField {
	return ef.baseField
}

func (ef *TEFByte) Value() interface{} {
	return ef.Byte()
}

func (ef *TEFByte) Byte() byte {
	return t.Byte(ef.baseField.Value())
}

func (ef *TEFByte) MarshalJSON() ([]byte, error) {
	return []byte(t.String(ef.Byte())), nil
}

func (ef *TEFByte) UnmarshalJSON(data []byte) error {
	if len(data) > 0 {
		ef.Set(t.Byte(string(data)))
	}
	return nil
}

func (ef *TEFByte) FieldType() xqi.FieldDataType {
	return xqi.FDTByte
}

func (ef *TEFByte) NewInstance(alias string, inherited ...interface{}) xqi.EntField {
	inst := &TEFByte{}
	var this interface{} = inst
	if len(inherited) > 0 {
		if _, ok := inherited[0].(xqi.EntField); ok {
			this = inherited[0]
		}
	}
	inst.baseField = ef.baseField.NewInstance(alias, this).(*baseField)
	return inst
}
