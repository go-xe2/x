package xentity

import (
	"github.com/go-xe2/x/type/t"
	"github.com/go-xe2/x/xf/ef/xqi"
)

type TEFUint64 struct {
	*baseField
}

func newEFUint64(entity xqi.Entity, defineName string, attr []xqi.XqAttribute, annotations map[string]interface{}, params ...interface{}) interface{} {
	inst := &TEFUint64{}
	base := newBaseField(entity, defineName, attr, annotations, params, inst)
	inst.baseField = base
	return inst.Constructor(inst)
}

func (ef *TEFUint64) Uint64() uint64 {
	return t.Uint64(ef.baseField.Value())
}

func (ef *TEFUint64) Value() interface{} {
	return ef.Uint64()
}

func (ef *TEFUint64) MarshalJSON() ([]byte, error) {
	return []byte(t.String(ef.Uint64())), nil
}

func (ef *TEFUint64) UnmarshalJSON(data []byte) error {
	ef.Set(t.Uint64(string(data)))
	return nil
}

func (ef *TEFUint64) FieldType() xqi.FieldDataType {
	return xqi.FDTUint64
}

func (ef *TEFUint64) NewInstance(alias string, inherited ...interface{}) xqi.EntField {
	inst := &TEFUint64{}
	var this interface{} = inst
	if len(inherited) > 0 {
		if _, ok := inherited[0].(xqi.EntField); ok {
			this = inherited[0]
		}
	}
	inst.baseField = ef.baseField.NewInstance(alias, this).(*baseField)
	return inst
}
