package xcrc32

import (
	"github.com/go-xe2/x/type/t"
	"hash/crc32"
)

func Encrypt(v interface{}) uint32 {
	return crc32.ChecksumIEEE(t.Bytes(v))
}

func EncryptString(v string) uint32 {
	return crc32.ChecksumIEEE([]byte(v))
}

func EncryptBytes(v []byte) uint32 {
	return crc32.ChecksumIEEE(v)
}
