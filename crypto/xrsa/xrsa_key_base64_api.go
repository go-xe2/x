package xrsa

import (
	"crypto/rsa"
	"errors"
	"github.com/go-xe2/x/core/exception"
	"github.com/go-xe2/x/encoding/xbase64"
)

func ParseBase64PublicKey(base64Key string) (*rsa.PublicKey, error) {
	key, err := xbase64.UrlDecode(base64Key)
	if err != nil {
		return nil, errors.New("公钥不是有效的base64编码")
	}
	return ParsePublicKey(key)
}

func ParseBase64PrivateKey(base64Key string) (*rsa.PrivateKey, error) {
	key, err := xbase64.UrlDecode(base64Key)
	if err != nil {
		return nil, exception.New("私钥不是有效的base64编码")
	}
	return ParsePrivateKey(key)
}
