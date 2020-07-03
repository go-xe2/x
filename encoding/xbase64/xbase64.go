package xbase64

import "encoding/base64"

// base64编码
func Encode(src []byte) []byte {
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(src)))
	base64.StdEncoding.Encode(dst, src)
	return dst
}

// base64解码
func Decode(dst []byte) ([]byte, error) {
	src := make([]byte, base64.StdEncoding.DecodedLen(len(dst)))
	n, err := base64.StdEncoding.Decode(src, dst)
	return src[:n], err
}

// base64编码字符串
func EncodeString(src []byte) string {
	return string(Encode(src))
}

// base64字符串解码
func DecodeString(str string) ([]byte, error) {
	return Decode([]byte(str))
}

func UrlEncode(src []byte) string {
	return base64.URLEncoding.EncodeToString(src)
}

func UrlDecode(str string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(str)
}
