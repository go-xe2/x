package xrsa

import (
	"github.com/go-xe2/x/encoding/xbase64"
)

func GenRsaKeyBase64(keyLen ...int) (privateKey, publicKey string, err error) {
	priData, pubData, err := GenRsaKey(keyLen...)
	if err != nil {
		return "", "", err
	}
	privateKey = xbase64.UrlEncode(priData)
	publicKey = xbase64.UrlEncode(pubData)
	return
}
