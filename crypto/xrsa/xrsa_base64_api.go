// 使用base64编码输出

package xrsa

import (
	"crypto/rsa"
	"github.com/go-xe2/x/core/exception"
	"github.com/go-xe2/x/encoding/xbase64"
)

func EncryptBase64(data string, key *rsa.PublicKey) (string, error) {
	v, err := Encrypt(data, key)
	if err != nil {
		return "", err
	}
	return xbase64.UrlEncode([]byte(v)), nil
}

func EncryptBase64ByRawKey(data string, rawPublicKey []byte) (string, error) {
	v, err := EncryptByRawKey(data, rawPublicKey)
	if err != nil {
		return "", err
	}
	return xbase64.UrlEncode([]byte(v)), nil
}

func EncryptBase64ByPemKey(data string, pemPublicKey string) (string, error) {
	v, err := EncryptByPemKey(data, pemPublicKey)
	if err != nil {
		return "", err
	}
	return xbase64.UrlEncode([]byte(v)), nil
}

func EncryptBase64ByPemFile(data string, pemPublicKeyFile string) (string, error) {
	v, err := EncryptByPemFile(data, pemPublicKeyFile)
	if err != nil {
		return "", err
	}
	return xbase64.UrlEncode([]byte(v)), nil
}

func EncryptBase64ByBase64key(data string, base64PubKey string) (string, error) {
	key, err := ParseBase64PublicKey(base64PubKey)
	if err != nil {
		return "", err
	}
	v, err := Encrypt(data, key)
	if err != nil {
		return "", err
	}
	return xbase64.UrlEncode([]byte(v)), nil
}

func DecryptBase64(base64Data string, key *rsa.PrivateKey) (string, error) {
	data, err := xbase64.UrlDecode(base64Data)
	if err != nil {
		return "", exception.New("数据不是有效的base64编码")
	}
	return Decrypt(string(data), key)
}

func DecryptBase64ByRawKey(base64Data string, rawPrivateKey []byte) (string, error) {
	data, err := xbase64.UrlDecode(base64Data)
	if err != nil {
		return "", exception.New("数据不是有效的base64编码")
	}
	return DecryptByRawKey(string(data), rawPrivateKey)
}

func DecryptBase64ByPemKey(base64Data string, pemPrivateKey string) (string, error) {
	data, err := xbase64.UrlDecode(base64Data)
	if err != nil {
		return "", exception.New("数据不是有效的base64编码")
	}
	return DecryptByPemKey(string(data), pemPrivateKey)
}

func DecryptBase64ByPemFile(base64Data string, pemPrivateKeyFile string) (string, error) {
	data, err := xbase64.UrlDecode(base64Data)
	if err != nil {
		return "", exception.New("数据不是有效的base64编码")
	}
	return DecryptByPemFile(string(data), pemPrivateKeyFile)
}

func DecryptBase64ByBase64Key(base64Data string, base64PriKey string) (string, error) {
	key, err := ParseBase64PrivateKey(base64PriKey)
	if err != nil {
		return "", err
	}
	data, err := xbase64.DecodeString(base64Data)
	if err != nil {
		return "", exception.New("数据不是有效的base64编码")
	}
	return Decrypt(string(data), key)
}

func SignBase64(data string, key *rsa.PrivateKey) (string, error) {
	v, err := Sign(data, key)
	if err != nil {
		return "", err
	}
	return xbase64.UrlEncode([]byte(v)), nil
}

func SignBase64ByRawKey(data string, rawPriKey []byte) (string, error) {
	v, err := SignByRawKey(data, rawPriKey)
	if err != nil {
		return "", err
	}
	return xbase64.UrlEncode([]byte(v)), nil
}

func SignBase64ByPemKey(data string, pemPriKey string) (string, error) {
	v, err := SignByPemKey(data, pemPriKey)
	if err != nil {
		return "", err
	}
	return xbase64.UrlEncode([]byte(v)), nil
}

func SignBase64ByPemFile(data string, pemPriKeyFile string) (string, error) {
	v, err := SignByPemFile(data, pemPriKeyFile)
	if err != nil {
		return "", err
	}
	return xbase64.UrlEncode([]byte(v)), nil
}

func SignBase64ByBase64Key(data string, base64PriKey string) (string, error) {
	key, err := ParseBase64PrivateKey(base64PriKey)
	if err != nil {
		return "", err
	}
	v, err := Sign(data, key)
	if err != nil {
		return "", err
	}
	return xbase64.EncodeString([]byte(v)), nil
}

func decodeBase64Sign(base64Sign string) (string, error) {
	sign, err := xbase64.UrlDecode(base64Sign)
	if err != nil {
		return "", exception.New("签名不是有效的base64编码")
	}
	return string(sign), nil
}

func VerifyBase64(data string, base64Sign string, key *rsa.PublicKey) (bool, error) {
	sign, err := decodeBase64Sign(base64Sign)
	if err != nil {
		return false, err
	}
	return Verify(data, sign, key)
}

func VerifyBase64ByRawKey(data string, base64Sign string, rawPubKey []byte) (bool, error) {
	sign, err := decodeBase64Sign(base64Sign)
	if err != nil {
		return false, err
	}
	return VerifyByRawKey(data, sign, rawPubKey)
}

func VerifyBase64ByPemKey(data string, base64Sign string, pemPubKey string) (bool, error) {
	sign, err := decodeBase64Sign(base64Sign)
	if err != nil {
		return false, err
	}
	return VerifyByPemKey(data, sign, pemPubKey)
}

func VerifyBase64ByPemFile(data string, base64Sign string, pemPubKeyFile string) (bool, error) {
	sign, err := decodeBase64Sign(base64Sign)
	if err != nil {
		return false, err
	}
	return VerifyByPemFile(data, sign, pemPubKeyFile)
}

func VerifyBase64ByBase64Key(data string, base64Sign string, base64PubKey string) (bool, error) {
	sign, err := decodeBase64Sign(base64Sign)
	if err != nil {
		return false, err
	}
	key, err := ParseBase64PublicKey(base64PubKey)
	if err != nil {
		return false, err
	}
	return Verify(data, sign, key)
}
