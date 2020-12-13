package xentity

import (
	"github.com/go-xe2/x/type/t"
	"github.com/go-xe2/x/xf/ef/xqi"
)

type TEFInt16 struct {
	*baseField
}

func newEFInt16(entity xqi.Entity, defineName string, attr []xqi.XqAttribute, annotations map[string]interface{}, params ...interface{}) interface{} {
	inst := &TEFInt16{}
	base := newBaseField(entity, defineName, attr, annotations, params, inst)
	inst.baseField = base
	return inst.Constructor(inst)
}

func (ef *TEFInt16) Value() interface{} {
	return ef.Int16()
}

func (ef *TEFInt16) Int16() int16 {
	return t.Int16(ef.baseField.Value())
}

func (ef *TEFInt16) TryInt16() (int16, bool) {
	if ef.IsOpen() {
		return ef.Int16(), true
	}
	return 0, false
}

func (ef *TEFInt16) MarshalJSON() ([]byte, error) {
	return []byte(t.String(ef.Int16())), nil
}

func (ef *TEFInt16) UnmarshalJSON(data []byte) error {
	ef.Set(t.Int16(string(data)))
	return nil
}

func (ef *TEFInt16) FieldType() xqi.FieldDataType {
	return xqi.FDTInt16
}

func (ef *TEFInt16) NewInstance(alias string, inherited ...interface{}) xqi.EntField {
	inst := &TEFInt16{}
	var this interface{} = inst
	if len(inherited) > 0 {
		if _, ok := inherited[0].(xqi.EntField); ok {
			this = inherited[0]
		}
	}
	inst.baseField = ef.baseField.NewInstance(alias, this).(*baseField)
	return inst
}
