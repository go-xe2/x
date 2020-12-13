package xentity

import (
	"github.com/go-xe2/x/type/t"
	"github.com/go-xe2/x/xf/ef/xqi"
)

type TEFDouble struct {
	*baseField
}

func newEFDouble(entity xqi.Entity, defineName string, attr []xqi.XqAttribute, annotations map[string]interface{}, params ...interface{}) interface{} {
	inst := &TEFDouble{}
	base := newBaseField(entity, defineName, attr, annotations, params, inst)
	inst.baseField = base
	return inst.Constructor(inst)
}

func (ef *TEFDouble) Value() interface{} {
	return ef.Double()
}

func (ef *TEFDouble) Double() float64 {
	return t.Float64(ef.baseField.Value())
}

func (ef *TEFDouble) TryDouble() (float64, bool) {
	if ef.IsOpen() {
		return ef.Double(), true
	}
	return 0, false
}

func (ef *TEFDouble) MarshalJSON() ([]byte, error) {
	return []byte(t.String(ef.Double())), nil
}

func (ef *TEFDouble) UnmarshalJSON(data []byte) error {
	ef.Set(t.Float64(string(data)))
	return nil
}

func (ef *TEFDouble) FieldType() xqi.FieldDataType {
	return xqi.FDTDouble
}

func (ef *TEFDouble) NewInstance(alias string, inherited ...interface{}) xqi.EntField {
	inst := &TEFDouble{}
	var this interface{} = inst
	if len(inherited) > 0 {
		if _, ok := inherited[0].(xqi.EntField); ok {
			this = inherited[0]
		}
	}
	inst.baseField = ef.baseField.NewInstance(alias, this).(*baseField)
	return inst
}
