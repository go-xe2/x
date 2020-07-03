package xmd5

import (
	"crypto/md5"
	"fmt"
	"github.com/go-xe2/x/core/exception"
	"github.com/go-xe2/x/type/t"
	"io"
	"os"
)

func Encrypt(v interface{}) (encrypt string, err error) {
	h := md5.New()
	if _, err = h.Write([]byte(t.Bytes(v))); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func EncryptString(v string) (encrypt string, err error) {
	return Encrypt(v)
}

func EncryptFile(path string) (encrypt string, err error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer func() {
		err = exception.Wrap(f.Close(), "file closing error")
	}()
	h := md5.New()
	_, err = io.Copy(h, f)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
