package xrsa

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"github.com/go-xe2/x/type/xbytes"
)

// 使用公钥加密data数据
func Encrypt(data string, publicKey *rsa.PublicKey) (string, error) {
	partLen := publicKey.N.BitLen()/8 - 11
	chunks := xbytes.Split([]byte(data), partLen)
	buffer := bytes.NewBufferString("")
	for _, chunk := range chunks {
		bytes, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, chunk)
		if err != nil {
			return "", err
		}
		buffer.Write(bytes)
	}
	return buffer.String(), nil
}

// 使用公钥字符串加密data数据
func EncryptByRawKey(data string, rawPublicKey []byte) (string, error) {
	pub, err := ParsePublicKey(rawPublicKey)
	if err != nil {
		return "", err
	}
	return Encrypt(data, pub)
}

// 使用pem编码的公钥字符串加密data数据
func EncryptByPemKey(data string, pemPublicKey string) (string, error) {
	key, err := DecodePem([]byte(pemPublicKey))
	if err != nil {
		return "", err
	}
	return EncryptByRawKey(data, key)
}

// 使用pem编码的公钥文件加密data数据
func EncryptByPemFile(data string, pemPublicFile string) (string, error) {
	key, err := PublicKeyPemFromFile(pemPublicFile)
	if err != nil {
		return "", err
	}
	return Encrypt(data, key)
}

// 使用私钥解密data数据
func Decrypt(data string, privateKey *rsa.PrivateKey) (string, error) {
	partLen := privateKey.N.BitLen() / 8
	chunks := xbytes.Split([]byte(data), partLen)
	buffer := bytes.NewBufferString("")
	for _, chunk := range chunks {
		decrypted, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, chunk)
		if err != nil {
			return "", err
		}
		buffer.Write(decrypted)
	}
	return buffer.String(), nil
}

// 使用私钥字符串解密data数据
func DecryptByRawKey(data string, rawKey []byte) (string, error) {
	pri, err := ParsePrivateKey(rawKey)
	if err != nil {
		return "", err
	}
	return Decrypt(data, pri)
}

// 使用pem编码的私钥字符串解密data数据
func DecryptByPemKey(data string, pemPrivateKey string) (string, error) {
	key, err := DecodePem([]byte(pemPrivateKey))
	if err != nil {
		return "", err
	}
	return DecryptByRawKey(data, key)
}

// 使用pem编码的私钥文件解密data数据
func DecryptByPemFile(data string, pemPrivateFile string) (string, error) {
	pri, err := PrivateKeyPemFromFile(pemPrivateFile)
	if err != nil {
		return "", err
	}
	return Decrypt(data, pri)
}

// 使用私钥对数据进行签名
func Sign(data string, pemPrivateKey *rsa.PrivateKey) (string, error) {
	h := crypto.SHA256.New()
	h.Write([]byte(data))
	hashed := h.Sum(nil)
	sign, err := rsa.SignPKCS1v15(rand.Reader, pemPrivateKey, crypto.SHA256, hashed)
	if err != nil {
		return "", err
	}
	return string(sign), nil
}

func SignByRawKey(data string, rawPrivateKey []byte) (string, error) {
	pri, err := ParsePrivateKey(rawPrivateKey)
	if err != nil {
		return "", err
	}
	return Sign(data, pri)
}

func SignByPemKey(data string, pemPrivateKey string) (string, error) {
	key, err := DecodePem([]byte(pemPrivateKey))
	if err != nil {
		return "", err
	}
	return SignByRawKey(data, key)
}

func SignByPemFile(data string, pemPrivateKeyFile string) (string, error) {
	key, err := PrivateKeyPemFromFile(pemPrivateKeyFile)
	if err != nil {
		return "", err
	}
	return Sign(data, key)
}

func Verify(data string, sign string, pemPublicKey *rsa.PublicKey) (bool, error) {
	h := crypto.SHA256.New()
	h.Write([]byte(data))
	hashed := h.Sum(nil)
	//decodedSign, err := base64.RawURLEncoding.DecodeString(sign)
	//if err != nil {
	//	return err
	//}
	e := rsa.VerifyPKCS1v15(pemPublicKey, crypto.SHA256, hashed, []byte(sign))
	if e == nil {
		return true, nil
	}
	return false, nil
}

func VerifyByRawKey(data string, sign string, rawPublicKey []byte) (bool, error) {
	pubKey, err := ParsePublicKey(rawPublicKey)
	if err != nil {
		return false, err
	}
	return Verify(data, sign, pubKey)
}

func VerifyByPemKey(data string, sign string, pemPublicKey string) (bool, error) {
	rawKey, err := DecodePem([]byte(pemPublicKey))
	if err != nil {
		return false, err
	}
	return VerifyByRawKey(data, sign, rawKey)
}

func VerifyByPemFile(data string, sign string, pemPublicKeyFile string) (bool, error) {
	pubKey, err := PublicKeyPemFromFile(pemPublicKeyFile)
	if err != nil {
		return false, err
	}
	return Verify(data, sign, pubKey)
}
