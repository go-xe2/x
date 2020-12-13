package xentity

import (
	"github.com/go-xe2/x/type/t"
	"github.com/go-xe2/x/xf/ef/xqi"
)

type TEFInt struct {
	*baseField
}

var _ xqi.EFInt = (*TEFInt)(nil)

func newEFInt(entity xqi.Entity, defineName string, attr []xqi.XqAttribute, annotations map[string]interface{}, params ...interface{}) interface{} {
	inst := &TEFInt{}
	base := newBaseField(entity, defineName, attr, annotations, params, inst)
	inst.baseField = base
	return inst.Constructor(inst)
}

func (ef *TEFInt) Value() interface{} {
	return ef.Int()
}

func (ef TEFInt) TryInt() (int, bool) {
	if ef.IsOpen() {
		return ef.Int(), true
	}
	return 0, false
}

func (ef *TEFInt) Int() int {
	return t.Int(ef.baseField.Value())
}

func (ef *TEFInt) MarshalJSON() ([]byte, error) {
	return []byte(t.String(ef.Int())), nil
}

func (ef *TEFInt) UnmarshalJSON(data []byte) error {
	ef.Set(t.Int(string(data)))
	return nil
}

func (ef *TEFInt) FieldType() xqi.FieldDataType {
	return xqi.FDTInt
}

func (ef *TEFInt) NewInstance(alias string, inherited ...interface{}) xqi.EntField {
	inst := &TEFInt{}
	var this interface{} = inst
	if len(inherited) > 0 {
		if _, ok := inherited[0].(xqi.EntField); ok {
			this = inherited[0]
		}
	}
	inst.baseField = ef.baseField.NewInstance(alias, this).(*baseField)
	return inst
}
