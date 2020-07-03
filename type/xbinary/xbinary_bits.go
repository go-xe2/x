package xbinary

import "github.com/go-xe2/x/core/exception"

// 二进制位(0|1)
type Bit int8

// 默认编码
func EncodeBits(bits []Bit, i int, l int) []Bit {
	return EncodeBitsWithUint(bits, uint(i), l)
}

// 将ui按位合并到bits数组中，并占length长度位(注意：uis数组中存放的是二进制的0|1数字)
func EncodeBitsWithUint(bits []Bit, ui uint, l int) []Bit {
	a := make([]Bit, l)
	for i := l - 1; i >= 0; i-- {
		a[i] = Bit(ui & 1)
		ui >>= 1
	}
	if bits != nil {
		return append(bits, a...)
	} else {
		return a
	}
}

// 将bits转换为[]byte，从左至右进行编码，不足1 byte按0往末尾补充
func EncodeBitsToBytes(bits []Bit) []byte {
	if len(bits)%8 != 0 {
		for i := 0; i < len(bits)%8; i++ {
			bits = append(bits, 0)
		}
	}
	b := make([]byte, 0)
	for i := 0; i < len(bits); i += 8 {
		b = append(b, byte(DecodeBitsToUint(bits[i:i+8])))
	}
	return b
}

// 解析为int
func DecodeBits(bits []Bit) int {
	v := int(0)
	for _, i := range bits {
		v = v<<1 | int(i)
	}
	return v
}

// 解析为uint
func DecodeBitsToUint(bits []Bit) uint {
	v := uint(0)
	for _, i := range bits {
		v = v<<1 | uint(i)
	}
	return v
}

// 解析[]byte为字位数组[]uint8
func DecodeBytesToBits(bs []byte) []Bit {
	bits := make([]Bit, 0)
	for _, b := range bs {
		bits = EncodeBitsWithUint(bits, uint(b), 8)
	}
	return bits
}

// 拆分为多个数组
func Split(buf []Bit, lim int) [][]Bit {
	if lim == 0 {
		panic(exception.New("数组拆分大小不能为0"))
	}
	var chunk []Bit
	bufLen := len(buf)
	chunkLen := bufLen/lim + 1
	chunks := make([][]Bit, 0, chunkLen)
	for bufLen >= lim {
		chunk, buf = buf[:lim], buf[lim:]
		chunks = append(chunks, chunk)
		bufLen = len(buf)
	}
	if bufLen > 0 {
		chunks = append(chunks, buf[:bufLen])
	}
	return chunks
}
