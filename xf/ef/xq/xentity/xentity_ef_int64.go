package xentity

import (
	"github.com/go-xe2/x/type/t"
	"github.com/go-xe2/x/xf/ef/xqi"
)

type TEFInt64 struct {
	*baseField
}

func newEFInt64(entity xqi.Entity, defineName string, attr []xqi.XqAttribute, annotations map[string]interface{}, params ...interface{}) interface{} {
	inst := &TEFInt64{}
	base := newBaseField(entity, defineName, attr, annotations, params, inst)
	inst.baseField = base
	return inst.Constructor(inst)
}

func (ef *TEFInt64) Int64() int64 {
	return t.Int64(ef.baseField.Value())
}

func (ef *TEFInt64) Value() interface{} {
	return ef.Int64()
}

func (ef *TEFInt64) MarshalJSON() ([]byte, error) {
	return []byte(t.String(ef.Int64())), nil
}

func (ef *TEFInt64) UnmarshalJSON(data []byte) error {
	ef.Set(t.Int64(string(data)))
	return nil
}

func (ef *TEFInt64) FieldType() xqi.FieldDataType {
	return xqi.FDTInt64
}

func (ef *TEFInt64) NewInstance(alias string, inherited ...interface{}) xqi.EntField {
	inst := &TEFInt64{}
	var this interface{} = inst
	if len(inherited) > 0 {
		if _, ok := inherited[0].(xqi.EntField); ok {
			this = inherited[0]
		}
	}
	inst.baseField = ef.baseField.NewInstance(alias, this).(*baseField)
	return inst
}
