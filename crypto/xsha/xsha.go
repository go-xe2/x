package xsha

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"github.com/go-xe2/x/core/exception"
	"github.com/go-xe2/x/type/t"
	"io"
	"os"
)

// sha1加密
func SHA1Encrypt(v interface{}) string {
	r := sha1.Sum(t.Bytes(v))
	return hex.EncodeToString(r[:])
}

// sha1加密字符串
func SHA1EncryptString(s string) string {
	return SHA1Encrypt(s)
}

// sha1加密文件
func SHA1EncryptFile(path string) (encrypt string, err error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer func() {
		err = exception.Wrap(f.Close(), "file closing error")
	}()
	h := sha1.New()
	_, err = io.Copy(h, f)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

// sha256验证
func SHA256Str(src string) string {
	h := sha256.New()
	h.Write([]byte(src)) // 需要加密的字符串为
	return hex.EncodeToString(h.Sum(nil))
}

// sha512验证
func SHA512Str(src string) string {
	h := sha512.New()
	h.Write([]byte(src)) // 需要加密的字符串为
	return hex.EncodeToString(h.Sum(nil))
}
