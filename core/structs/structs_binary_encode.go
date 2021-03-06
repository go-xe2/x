package structs

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
)

// 针对基本类型进行二进制打包，支持的基本数据类型包括:
// int/8/16/32/64、uint/8/16/32/64、float32/64、bool、string、[]byte。
// 其他未知类型使用 fmt.Sprintf("%v", value) 转换为字符串之后处理。
func encode(values ...interface{}) []byte {
	buf := new(bytes.Buffer)
	for i := 0; i < len(values); i++ {
		if values[i] == nil {
			return buf.Bytes()
		}
		switch value := values[i].(type) {
		case int:
			buf.Write(encodeInt(value))
		case int8:
			buf.Write(encodeInt8(value))
		case int16:
			buf.Write(encodeInt16(value))
		case int32:
			buf.Write(encodeInt32(value))
		case int64:
			buf.Write(encodeInt64(value))
		case uint:
			buf.Write(encodeUint(value))
		case uint8:
			buf.Write(encodeUint8(value))
		case uint16:
			buf.Write(encodeUint16(value))
		case uint32:
			buf.Write(encodeUint32(value))
		case uint64:
			buf.Write(encodeUint64(value))
		case bool:
			buf.Write(encodeBool(value))
		case string:
			buf.Write(encodeString(value))
		case []byte:
			buf.Write(value)
		case float32:
			buf.Write(encodeFloat32(value))
		case float64:
			buf.Write(encodeFloat64(value))
		default:
			if err := binary.Write(buf, binary.LittleEndian, value); err != nil {
				buf.Write(encodeString(fmt.Sprintf("%v", value)))
			}
		}
	}
	return buf.Bytes()
}

// 将变量转换为二进制[]byte，并指定固定的[]byte长度返回，长度单位为字节(byte)；
// 如果转换的二进制长度超过指定长度，那么进行截断处理
func encodeByLength(length int, values ...interface{}) []byte {
	b := encode(values...)
	if len(b) < length {
		b = append(b, make([]byte, length-len(b))...)
	} else if len(b) > length {
		b = b[0:length]
	}
	return b
}

// 整形二进制解包，注意第二个及其后参数为字长确定的整形变量的指针地址，以便确定解析的[]byte长度，小端存储
// 例如：int8/16/32/64、uint8/16/32/64、float32/64等等
func decode(b []byte, values ...interface{}) error {
	buf := bytes.NewBuffer(b)
	for i := 0; i < len(values); i++ {
		err := binary.Read(buf, binary.LittleEndian, values[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func encodeString(s string) []byte {
	return []byte(s)
}

func eecodeToString(b []byte) string {
	return string(b)
}

func encodeBool(b bool) []byte {
	if b == true {
		return []byte{1}
	} else {
		return []byte{0}
	}
}

// 自动识别int类型长度，转换为[]byte
func encodeInt(i int) []byte {
	if i <= math.MaxInt8 {
		return encodeInt8(int8(i))
	} else if i <= math.MaxInt16 {
		return encodeInt16(int16(i))
	} else if i <= math.MaxInt32 {
		return encodeInt32(int32(i))
	} else {
		return encodeInt64(int64(i))
	}
}

// 自动识别uint类型长度，转换为[]byte
func encodeUint(i uint) []byte {
	if i <= math.MaxUint8 {
		return encodeUint8(uint8(i))
	} else if i <= math.MaxUint16 {
		return encodeUint16(uint16(i))
	} else if i <= math.MaxUint32 {
		return encodeUint32(uint32(i))
	} else {
		return encodeUint64(uint64(i))
	}
}

func encodeInt8(i int8) []byte {
	return []byte{byte(i)}
}

func encodeUint8(i uint8) []byte {
	return []byte{byte(i)}
}

func encodeInt16(i int16) []byte {
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, uint16(i))
	return b
}

func encodeUint16(i uint16) []byte {
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, i)
	return b
}

func encodeInt32(i int32) []byte {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, uint32(i))
	return b
}

func encodeUint32(i uint32) []byte {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, i)
	return b
}

func encodeInt64(i int64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(i))
	return b
}

func encodeUint64(i uint64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, i)
	return b
}

func encodeFloat32(f float32) []byte {
	bits := math.Float32bits(f)
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, bits)
	return b
}

func encodeFloat64(f float64) []byte {
	bits := math.Float64bits(f)
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, bits)
	return b
}

// 将二进制解析为int类型，根据[]byte的长度进行自动转换.
// 注意内部使用的是uint*，使用int会造成位丢失。
func decodeToInt(b []byte) int {
	if len(b) < 2 {
		return int(decodeToUint8(b))
	} else if len(b) < 3 {
		return int(decodeToUint16(b))
	} else if len(b) < 5 {
		return int(decodeToUint32(b))
	} else {
		return int(decodeToUint64(b))
	}
}

// 将二进制解析为uint类型，根据[]byte的长度进行自动转换
func decodeToUint(b []byte) uint {
	if len(b) < 2 {
		return uint(decodeToUint8(b))
	} else if len(b) < 3 {
		return uint(decodeToUint16(b))
	} else if len(b) < 5 {
		return uint(decodeToUint32(b))
	} else {
		return uint(decodeToUint64(b))
	}
}

// 将二进制解析为bool类型，识别标准是判断二进制中数值是否都为0，或者为空。
func decodeToBool(b []byte) bool {
	if len(b) == 0 {
		return false
	}
	if bytes.Compare(b, make([]byte, len(b))) == 0 {
		return false
	}
	return true
}

func decodeToInt8(b []byte) int8 {
	return int8(b[0])
}

func decodeToUint8(b []byte) uint8 {
	return uint8(b[0])
}

func decodeToInt16(b []byte) int16 {
	return int16(binary.LittleEndian.Uint16(LeFillUpSize(b, 2)))
}

func decodeToUint16(b []byte) uint16 {
	return binary.LittleEndian.Uint16(LeFillUpSize(b, 2))
}

func decodeToInt32(b []byte) int32 {
	return int32(binary.LittleEndian.Uint32(LeFillUpSize(b, 4)))
}

func decodeToUint32(b []byte) uint32 {
	return binary.LittleEndian.Uint32(LeFillUpSize(b, 4))
}

func decodeToInt64(b []byte) int64 {
	return int64(binary.LittleEndian.Uint64(LeFillUpSize(b, 8)))
}

func decodeToUint64(b []byte) uint64 {
	return binary.LittleEndian.Uint64(LeFillUpSize(b, 8))
}

func decodeToFloat32(b []byte) float32 {
	return math.Float32frombits(binary.LittleEndian.Uint32(LeFillUpSize(b, 4)))
}

func decodeToFloat64(b []byte) float64 {
	return math.Float64frombits(binary.LittleEndian.Uint64(LeFillUpSize(b, 8)))
}

// 当b位数不够时，进行高位补0。
// 注意这里为了不影响原有输入参数，是采用的值复制设计。
func LeFillUpSize(b []byte, l int) []byte {
	if len(b) >= l {
		return b[:l]
	}
	c := make([]byte, l)
	copy(c, b)
	return c
}
