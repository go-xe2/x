package xqi

type DSField interface {
	FieldName() string
	FieldType() FieldDataType
	FieldIndex() int
	Value() interface{}
	//AsInt() int
	//AsInt8() int8
	//AsInt16() int16
	//AsInt32() int32
	//AsInt64() int64
	//AsFloat32() float32
	//AsFloat64() float64
	//AsBool() bool
	//AsDatetime() xtime.Time
	//AsString() string
}
