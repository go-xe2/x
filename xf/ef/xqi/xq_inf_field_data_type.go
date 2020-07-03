package xqi

type FieldDataType uint8

const (
	FDTUnknown FieldDataType = iota
	FDTString
	FDTBool
	FDTInt
	FDTInt8
	FDTInt16
	FDTInt32
	FDTInt64
	FDTUint
	FDTUint8
	FDTUint16
	FDTUint32
	FDTUint64
	// float32
	FDTFloat
	// float64
	FDTDouble
	FDTByte
	FDTDatetime
	FDTBinary
	// 用户自定义类型起始值, 不允许用户定义类型与内值类型重复
	FDTDefineBase
)

var fieldDataTypeJson = map[FieldDataType]string{
	FDTString:   "S",
	FDTInt:      "I",
	FDTInt8:     "I8",
	FDTInt16:    "I1",
	FDTInt32:    "I3",
	FDTInt64:    "I6",
	FDTUint:     "U",
	FDTUint8:    "U8",
	FDTUint16:   "U1",
	FDTUint32:   "U3",
	FDTUint64:   "U6",
	FDTBool:     "BL",
	FDTByte:     "BE",
	FDTFloat:    "F",
	FDTDouble:   "D",
	FDTDatetime: "T",
	FDTBinary:   "BY",
	FDTUnknown:  "UK",
}

var jsFieldDataType2Type = map[string]FieldDataType{
	"S":  FDTString,
	"I":  FDTInt,
	"I8": FDTInt8,
	"I1": FDTInt16,
	"I3": FDTInt32,
	"I6": FDTInt64,
	"U":  FDTUint,
	"U8": FDTUint8,
	"U1": FDTUint16,
	"U3": FDTUint32,
	"U6": FDTUint64,
	"BL": FDTBool,
	"BE": FDTByte,
	"F":  FDTFloat,
	"D":  FDTDouble,
	"T":  FDTDatetime,
	"BY": FDTBinary,
	"UK": FDTUnknown,
}

func JsFieldDataTypeToType(szType string) (FieldDataType, bool) {
	if s, ok := jsFieldDataType2Type[szType]; ok {
		return s, true
	}
	if len(szType) > 2 && szType[:2] == "UT" {
		return FDTDefineBase, true
	}
	return FDTUnknown, false
}

func (fdt FieldDataType) Json(ut ...string) string {
	if fdt >= FDTDefineBase && len(ut) > 0 {
		return "UT:" + ut[0]
	}
	if s, ok := fieldDataTypeJson[fdt]; ok {
		return s
	}
	return "UK"
}

func (fdt FieldDataType) String() string {
	switch fdt {
	case FDTString:
		return "string"
	case FDTBool:
		return "bool"
	case FDTInt:
		return "int"
	case FDTInt8:
		return "int8"
	case FDTInt16:
		return "int16"
	case FDTInt32:
		return "int32"
	case FDTInt64:
		return "int64"
	case FDTUint:
		return "uint"
	case FDTUint8:
		return "uint8"
	case FDTUint16:
		return "uint16"
	case FDTUint32:
		return "uint32"
	case FDTUint64:
		return "uint64"
	case FDTFloat:
		return "float"
	case FDTDouble:
		return "double"
	case FDTByte:
		return "byte"
	case FDTDatetime:
		return "datetime"
	case FDTBinary:
		return "binary"
	case FDTDefineBase:
		return "userDefineType"
	}
	return "unknown"
}
