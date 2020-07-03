package xutil

import (
	"github.com/go-xe2/x/type/t"
	"strings"
)

// 比较器
// a < b 返回 < 0
// a == b 返回= 0
// a > b 返回 > 0
type Comparator func(a, b interface{}) int

func ComparatorString(a, b interface{}) int {
	return strings.Compare(t.String(a), t.String(b))
}

func ComparatorInt(a, b interface{}) int {
	return t.Int(a) - t.Int(b)
}

func ComparatorInt8(a, b interface{}) int {
	return int(t.Int8(a) - t.Int8(b))
}

func ComparatorInt16(a, b interface{}) int {
	return int(t.Int16(a) - t.Int16(b))
}

func ComparatorInt32(a, b interface{}) int {
	return int(t.Int32(a) - t.Int32(b))
}

func ComparatorInt64(a, b interface{}) int {
	return int(t.Int64(a) - t.Int64(b))
}

func ComparatorUint(a, b interface{}) int {
	return int(t.Uint(a) - t.Uint(b))
}

func ComparatorUint8(a, b interface{}) int {
	return int(t.Uint8(a) - t.Uint8(b))
}

func ComparatorUint16(a, b interface{}) int {
	return int(t.Uint16(a) - t.Uint16(b))
}

func ComparatorUint32(a, b interface{}) int {
	return int(t.Uint32(a) - t.Uint32(b))
}

func ComparatorUint64(a, b interface{}) int {
	return int(t.Uint64(a) - t.Uint64(b))
}

func ComparatorFloat32(a, b interface{}) int {
	return int(t.Float32(a) - t.Float32(b))
}

func ComparatorFloat64(a, b interface{}) int {
	return int(t.Float64(a) - t.Float64(b))
}

func ComparatorByte(a, b interface{}) int {
	return int(t.Byte(a) - t.Byte(b))
}

func ComparatorRune(a, b interface{}) int {
	return int(t.Rune(a) - t.Rune(b))
}

func ComparatorTime(a, b interface{}) int {
	aTime := t.Time(a)
	bTime := t.Time(b)
	switch {
	case aTime.After(bTime):
		return 1
	case aTime.Before(bTime):
		return -1
	default:
		return 0
	}
}
