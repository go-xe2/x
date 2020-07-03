package xentity

import (
	"github.com/go-xe2/x/type/t"
	"github.com/go-xe2/x/xf/ef/xqi"
)

type TEFFloat struct {
	*baseField
}

func newEFFloat(entity xqi.Entity, defineName string, attr []xqi.XqAttribute, annotations map[string]interface{}, params ...interface{}) interface{} {
	inst := &TEFFloat{}
	base := newBaseField(entity, defineName, attr, annotations, params, inst)
	inst.baseField = base
	return inst.Constructor(inst)
}

func (ef *TEFFloat) Value() interface{} {
	return ef.Float()
}

func (ef *TEFFloat) Float() float32 {
	return t.Float32(ef.baseField.Value())
}

func (ef *TEFFloat) MarshalJSON() ([]byte, error) {
	return []byte(t.String(ef.Float())), nil
}

func (ef *TEFFloat) UnmarshalJSON(data []byte) error {
	ef.Set(t.Float32(string(data)))
	return nil
}

func (ef *TEFFloat) FieldType() xqi.FieldDataType {
	return xqi.FDTFloat
}

func (ef *TEFFloat) NewInstance(alias string, inherited ...interface{}) xqi.EntField {
	inst := &TEFFloat{}
	var this interface{} = inst
	if len(inherited) > 0 {
		if _, ok := inherited[0].(xqi.EntField); ok {
			this = inherited[0]
		}
	}
	inst.baseField = ef.baseField.NewInstance(alias, this).(*baseField)
	return inst
}
