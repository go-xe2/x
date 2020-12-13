package xentity

import (
	"fmt"
	"github.com/go-xe2/x/type/t"
	"github.com/go-xe2/x/type/xtime"
	"github.com/go-xe2/x/xf/ef/xq/xqcomm"
	"github.com/go-xe2/x/xf/ef/xqi"
	"time"
)

type TEFDatetime struct {
	*baseField
}

var _ xqi.EFDate = (*TEFDatetime)(nil)

func newEFDatetime(entity xqi.Entity, defineName string, attr []xqi.XqAttribute, annotations map[string]interface{}, params ...interface{}) interface{} {
	inst := &TEFDatetime{}
	base := newBaseField(entity, defineName, attr, annotations, params, inst)
	inst.baseField = base
	return inst.Constructor(inst)
}

func (ef *TEFDatetime) Supper() xqi.EntField {
	return ef.baseField
}

func (ef *TEFDatetime) Value() interface{} {
	return ef.Date()
}

func (ef *TEFDatetime) Date() time.Time {
	n := t.Int64(ef.baseField.Value())
	return time.Unix(n, 0)
}

func (ef *TEFDatetime) TryDate() (time.Time, bool) {
	if ef.IsOpen() {
		return ef.Date(), true
	}
	return time.Now(), false
}

func (ef *TEFDatetime) Set(val interface{}) xqi.FieldValue {
	tm := t.XTime(val)
	if tm == nil {
		tm = xtime.New(time.Unix(0, 0))
	}
	return xqcomm.NewFieldValue(ef, tm.Unix())
}

func (ef *TEFDatetime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", t.XTime(ef.Date()).String())), nil
}

func (ef *TEFDatetime) UnmarshalJSON(data []byte) error {
	ef.Set(t.XTime(string(data)).Time)
	return nil
}

func (ef *TEFDatetime) FieldType() xqi.FieldDataType {
	return xqi.FDTDatetime
}

func (ef *TEFDatetime) Formatter() string {
	if ef.formatter != "" {
		return ef.formatter
	}
	// 默认返回日志格式化
	return FormatDate
}

func (ef *TEFDatetime) NewInstance(alias string, inherited ...interface{}) xqi.EntField {
	inst := &TEFDatetime{}
	var this interface{} = inst
	if len(inherited) > 0 {
		if _, ok := inherited[0].(xqi.EntField); ok {
			this = inherited[0]
		}
	}
	inst.baseField = ef.baseField.NewInstance(alias, this).(*baseField)
	return inst
}
