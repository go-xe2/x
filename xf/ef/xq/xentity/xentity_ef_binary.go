package xentity

import (
	"fmt"
	"github.com/go-xe2/x/type/t"
	"github.com/go-xe2/x/xf/ef/xqi"
)

type TEFBinary struct {
	*baseField
}

func newEFBinary(entity xqi.Entity, defineName string, attr []xqi.XqAttribute, annotations map[string]interface{}, params ...interface{}) interface{} {
	inst := &TEFBinary{}
	base := newBaseField(entity, defineName, attr, annotations, params, inst)
	inst.baseField = base
	return inst.Constructor(inst)
}

func (ef *TEFBinary) Supper() xqi.EntField {
	return ef.baseField
}

func (ef *TEFBinary) Value() interface{} {
	return ef.Bytes()
}

func (ef *TEFBinary) Bytes() []byte {
	return t.Bytes(ef.baseField.Value())
}

func (ef *TEFBinary) TryBytes() ([]byte, bool) {
	if ef.IsOpen() {
		return ef.Bytes(), true
	}
	return nil, false
}

func (ef *TEFBinary) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", t.String(ef.Bytes()))), nil
}

func (ef *TEFBinary) UnmarshalJSON(data []byte) error {
	ef.Set(t.Bytes(string(data)))
	return nil
}

func (ef *TEFBinary) FieldType() xqi.FieldDataType {
	return xqi.FDTBinary
}

func (ef *TEFBinary) NewInstance(alias string, inherited ...interface{}) xqi.EntField {
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
