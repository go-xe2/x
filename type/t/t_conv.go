package t

import (
	"github.com/go-xe2/x/type/xconv"
)

// Convert converts the variable <i> to the type <t>, the type <t> is specified by string.
// The unnecessary parameter <params> is used for additional parameter passing.
func Convert(i interface{}, t string, params ...interface{}) interface{} {
	return xconv.Convert(i, t, params...)
}

func String(v interface{}, def ...string) string {
	return xconv.String(v, def...)
}

// Int converts <i> to int.
func Int(i interface{}, def ...int) int {
	return xconv.Int(i, def...)
}

// Int8 converts <i> to int8.
func Int8(i interface{}, def ...int8) int8 {
	return xconv.Int8(i, def...)
}

// Int16 converts <i> to int16.
func Int16(i interface{}, def ...int16) int16 {
	return xconv.Int16(i, def...)
}

// Int32 converts <i> to int32.
func Int32(i interface{}, def ...int32) int32 {
	return xconv.Int32(i, def...)
}

func Int64(v interface{}, def ...int64) int64 {
	return xconv.Int64(v, def...)
}

// Uint converts <i> to uint.
func Uint(i interface{}, def ...uint) uint {
	return xconv.Uint(i, def...)
}

// Uint8 converts <i> to uint8.
func Uint8(i interface{}, def ...uint8) uint8 {
	return xconv.Uint8(i, def...)
}

// Uint16 converts <i> to uint16.
func Uint16(i interface{}, def ...uint16) uint16 {
	return xconv.Uint16(i, def...)
}

// Uint32 converts <i> to uint32.
func Uint32(i interface{}, def ...uint32) uint32 {
	return xconv.Uint32(i, def...)
}

// Uint64 converts <i> to uint64.
func Uint64(i interface{}, def ...uint64) uint64 {
	return xconv.Uint64(i, def...)
}

func Float(i interface{}, def ...float32) float32 {
	return xconv.Float(i, def...)
}

// Float32 converts <i> to float32.
func Float32(i interface{}, def ...float32) float32 {
	return xconv.Float32(i, def...)
}

// Float64 converts <i> to float64.
func Float64(i interface{}, def ...float64) float64 {
	return xconv.Float64(i, def...)
}

// Byte converts <i> to byte.
func Byte(i interface{}) byte {
	return xconv.Byte(i)
}

// Bytes converts <i> to []byte.
func Bytes(i interface{}) []byte {
	return xconv.Bytes(i)
}

// Rune converts <i> to rune.
func Rune(i interface{}) rune {
	return xconv.Rune(i)
}

// Runes converts <i> to []rune.
func Runes(i interface{}) []rune {
	return xconv.Runes(i)
}

func Bool(i interface{}, def ...bool) bool {
	return xconv.Bool(i, def...)
}
