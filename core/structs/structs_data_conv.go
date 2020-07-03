package structs

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

var (
	emptyStrMap = map[string]struct{}{
		"":      struct{}{},
		"0":     struct{}{},
		"off":   struct{}{},
		"false": struct{}{},
		"False": struct{}{},
	}
)

type iString interface {
	String() string
}

func anyToString(v interface{}, def ...string) string {
	var defValue = ""
	if len(def) > 0 {
		defValue = def[0]
	}
	if v == nil {
		return defValue
	}
	switch value := v.(type) {
	case int:
		return strconv.FormatInt(int64(value), 10)
	case int8:
		return strconv.Itoa(int(value))
	case int16:
		return strconv.Itoa(int(value))
	case int32:
		return strconv.Itoa(int(value))
	case int64:
		return strconv.FormatInt(int64(value), 10)
	case uint:
		return strconv.FormatUint(uint64(value), 10)
	case uint8:
		return strconv.FormatUint(uint64(value), 10)
	case uint16:
		return strconv.FormatUint(uint64(value), 10)
	case uint32:
		return strconv.FormatUint(uint64(value), 10)
	case uint64:
		return strconv.FormatUint(uint64(value), 10)
	case float32:
		return strconv.FormatFloat(float64(value), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(value, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(value)
	case string:
		return value
	case []byte:
		return string(value)
	case []rune:
		return string(value)
	default:
		if f, ok := value.(iString); ok {
			return f.String()
		} else if f, ok := value.(error); ok {
			return f.Error()
		} else if f, ok := value.(error); ok {
			return fmt.Sprintf("error:%s", f.Error())
		} else {
			if jsonContent, err := json.Marshal(value); err != nil {
				return fmt.Sprint(value)
			} else {
				return string(jsonContent)
			}
		}
	}
}

// Int converts <i> to int.
func anyToInt(i interface{}, def ...int) int {
	var defValue = 0
	if len(def) > 0 {
		defValue = def[0]
	}
	if i == nil {
		return defValue
	}
	if v, ok := i.(int); ok {
		return v
	}
	return int(anyToInt64(i, anyToInt64(defValue)))
}

// Int8 converts <i> to int8.
func anyToInt8(i interface{}, def ...int8) int8 {
	var defV int8 = 0
	if len(def) > 0 {
		defV = def[0]
	}
	if i == nil {
		return defV
	}
	if v, ok := i.(int8); ok {
		return v
	}
	return int8(anyToInt64(i, anyToInt64(defV)))
}

// Int16 converts <i> to int16.
func anyToInt16(i interface{}, def ...int16) int16 {
	var defV int16 = 0
	if len(def) > 0 {
		defV = def[0]
	}
	if i == nil {
		return defV
	}
	if v, ok := i.(int16); ok {
		return v
	}
	return int16(anyToInt64(i, anyToInt64(defV)))
}

// Int32 converts <i> to int32.
func anyToInt32(i interface{}, def ...int32) int32 {
	var defV int32 = 0
	if len(def) > 0 {
		defV = def[0]
	}
	if i == nil {
		return defV
	}
	if v, ok := i.(int32); ok {
		return v
	}
	return int32(anyToInt64(i, anyToInt64(defV)))
}

func anyToInt64(v interface{}, def ...int64) int64 {
	var defVaue int64 = 0
	if len(def) > 0 {
		defVaue = def[0]
	}
	if v == nil {
		return defVaue
	}
	switch value := v.(type) {
	case int:
		return int64(value)
	case int8:
		return int64(value)
	case int16:
		return int64(value)
	case int32:
		return int64(value)
	case int64:
		return value
	case uint:
		return int64(value)
	case uint8:
		return int64(value)
	case uint16:
		return int64(value)
	case uint32:
		return int64(value)
	case uint64:
		return int64(value)
	case float32:
		return int64(value)
	case float64:
		return int64(value)
	case bool:
		if value {
			return 1
		}
		return 0
	case []byte:
		return decodeToInt64(value)
	default:
		s := anyToString(value)
		// Hexadecimal
		if len(s) > 2 && s[0] == '0' && (s[1] == 'x' || s[1] == 'X') {
			if v, e := strconv.ParseInt(s[2:], 16, 64); e == nil {
				return v
			}
		}
		// Octal
		if len(s) > 1 && s[0] == '0' {
			if v, e := strconv.ParseInt(s[1:], 8, 64); e == nil {
				return v
			}
		}
		// Decimal
		if v, e := strconv.ParseInt(s, 10, 64); e == nil {
			return v
		}
		// Float64
		return int64(anyToFloat64(value))
	}
}

// Uint converts <i> to uint.
func anyToUint(i interface{}, def ...uint) uint {
	var defV uint = 0
	if len(def) > 0 {
		defV = def[0]
	}
	if i == nil {
		return defV
	}
	if v, ok := i.(uint); ok {
		return v
	}
	return uint(anyToUint64(i, anyToUint64(defV)))
}

// Uint8 converts <i> to uint8.
func anyToUint8(i interface{}, def ...uint8) uint8 {
	var defV uint8 = 0
	if len(def) > 0 {
		defV = def[0]
	}
	if i == nil {
		return defV
	}
	if v, ok := i.(uint8); ok {
		return v
	}
	return uint8(anyToUint64(i, anyToUint64(defV)))
}

// Uint16 converts <i> to uint16.
func anyToUint16(i interface{}, def ...uint16) uint16 {
	var defV uint16 = 0
	if len(def) > 0 {
		defV = def[0]
	}
	if i == nil {
		return defV
	}
	if v, ok := i.(uint16); ok {
		return v
	}
	return uint16(anyToUint64(i, uint64(defV)))
}

// Uint32 converts <i> to uint32.
func anyToUint32(i interface{}, def ...uint32) uint32 {
	var defV uint32 = 0
	if len(def) > 0 {
		defV = def[0]
	}
	if i == nil {
		return defV
	}
	if v, ok := i.(uint32); ok {
		return v
	}
	return uint32(anyToUint64(i, uint64(defV)))
}

// Uint64 converts <i> to uint64.
func anyToUint64(i interface{}, def ...uint64) uint64 {
	var defV uint64 = 0
	if len(def) > 0 {
		defV = def[0]
	}
	if i == nil {
		return defV
	}
	switch value := i.(type) {
	case int:
		return uint64(value)
	case int8:
		return uint64(value)
	case int16:
		return uint64(value)
	case int32:
		return uint64(value)
	case int64:
		return uint64(value)
	case uint:
		return uint64(value)
	case uint8:
		return uint64(value)
	case uint16:
		return uint64(value)
	case uint32:
		return uint64(value)
	case uint64:
		return value
	case float32:
		return uint64(value)
	case float64:
		return uint64(value)
	case bool:
		if value {
			return 1
		}
		return 0
	case []byte:
		return decodeToUint64(value)
	default:
		s := anyToString(value)
		// Hexadecimal
		if len(s) > 2 && s[0] == '0' && (s[1] == 'x' || s[1] == 'X') {
			if v, e := strconv.ParseUint(s[2:], 16, 64); e == nil {
				return v
			}
		}
		// Octal
		if len(s) > 1 && s[0] == '0' {
			if v, e := strconv.ParseUint(s[1:], 8, 64); e == nil {
				return v
			}
		}
		// Decimal
		if v, e := strconv.ParseUint(s, 10, 64); e == nil {
			return v
		}
		// Float64
		return uint64(anyToFloat64(value))
	}
}

func anyToFloat(i interface{}, def ...float32) float32 {
	return anyToFloat32(i, def...)
}

// Float32 converts <i> to float32.
func anyToFloat32(i interface{}, def ...float32) float32 {
	var defV float32 = 0
	if len(def) > 0 {
		defV = def[0]
	}
	if i == nil {
		return defV
	}
	switch value := i.(type) {
	case float32:
		return value
	case float64:
		return float32(value)
	case []byte:
		return decodeToFloat32(value)
	default:
		v, _ := strconv.ParseFloat(anyToString(i), 64)
		return float32(v)
	}
}

// Float64 converts <i> to float64.
func anyToFloat64(i interface{}, def ...float64) float64 {
	var defValue float64 = 0
	if len(def) > 0 {
		defValue = def[0]
	}
	if i == nil {
		return defValue
	}
	switch value := i.(type) {
	case float32:
		return float64(value)
	case float64:
		return value
	case []byte:
		return decodeToFloat64(value)
	default:
		v, _ := strconv.ParseFloat(anyToString(i), 64)
		return v
	}
}

// Byte converts <i> to byte.
func anyToByte(i interface{}) byte {
	if v, ok := i.(byte); ok {
		return v
	}
	return byte(anyToUint8(i))
}

// Bytes converts <i> to []byte.
func anyToBytes(i interface{}) []byte {
	if i == nil {
		return nil
	}
	switch value := i.(type) {
	case string:
		return []byte(value)
	case []byte:
		return value
	default:
		return encode(i)
	}
}

// Rune converts <i> to rune.
func anyToRune(i interface{}) rune {
	if v, ok := i.(rune); ok {
		return v
	}
	return rune(anyToInt32(i))
}

// Runes converts <i> to []rune.
func anyToRunes(i interface{}) []rune {
	if v, ok := i.([]rune); ok {
		return v
	}
	return []rune(anyToString(i))
}

func anyToBool(i interface{}, def ...bool) bool {
	bDef := false
	if len(def) > 0 {
		bDef = def[0]
	}
	if i == nil {
		return bDef
	}
	switch value := i.(type) {
	case bool:
		return value
	case string:
		if _, ok := emptyStrMap[value]; ok {
			return bDef
		}
		return true
	default:
		rv := reflect.ValueOf(i)
		switch rv.Kind() {
		case reflect.Ptr:
			return !rv.IsNil()
		case reflect.Map:
			fallthrough
		case reflect.Array:
			fallthrough
		case reflect.Slice:
			return rv.Len() != 0
		case reflect.Struct:
			return true
		default:
			s := anyToString(i)
			if _, ok := emptyStrMap[s]; ok {
				return bDef
			}
			return true
		}
	}
}
