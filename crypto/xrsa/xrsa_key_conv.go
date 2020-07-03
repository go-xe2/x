package xrsa

import (
	"bytes"
	"encoding/pem"
	"github.com/go-xe2/x/core/exception"
	"github.com/go-xe2/x/encoding/xbase64"
	"github.com/go-xe2/x/os/xfile"
	"io"
	"io/ioutil"
	"os"
)

// base64编码的公钥转成pem编码
func Base64PublicKeyPemWriter(base64Key string, writer io.Writer) error {
	keyData, err := xbase64.UrlDecode(base64Key)
	if err != nil {
		return exception.New("公钥不是有效的base64编码")
	}
	block := &pem.Block{
		Type:  "RSA Public key",
		Bytes: keyData,
	}
	return pem.Encode(writer, block)
}

// base64编码的公钥转码成pem的字符串
func Base64PublicKeyToPem(base64Key string) (string, error) {
	writer := bytes.NewBufferString("")
	err := Base64PublicKeyPemWriter(base64Key, writer)
	if err != nil {
		return "", err
	}
	return writer.String(), nil
}

// base64编码的公钥转存为pem文件
func Base64PublicKeyToPemFile(base64Key string, publicKeyFile string) error {
	writer, err := xfile.OpenFile(publicKeyFile, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	return Base64PublicKeyPemWriter(base64Key, writer)
}

func Base64PrivateKeyPemWriter(base64Key string, writer io.Writer) error {
	keyData, err := xbase64.UrlDecode(base64Key)
	if err != nil {
		return err
	}
	block := &pem.Block{
		Type:  "RSA Private key",
		Bytes: keyData,
	}
	return pem.Encode(writer, block)
}

func Base64PrivateKeyToPem(base64Key string) (string, error) {
	writer := bytes.NewBufferString("")
	err := Base64PrivateKeyPemWriter(base64Key, writer)
	if err != nil {
		return "", err
	}
	return writer.String(), nil
}

func Base64PrivateKeyToPemFile(base64Key string, pemPrivateFile string) error {
	w, err := xfile.OpenFile(pemPrivateFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	return Base64PrivateKeyPemWriter(base64Key, w)
}

func PemPublicKeyToBase64(pemKey string) (string, error) {
	rawKey, err := DecodePem([]byte(pemKey))
	if err != nil {
		return "", err
	}
	return xbase64.UrlEncode(rawKey), nil
}

func PemPublicKeyFileToBase64(pemPublicKeyFile string) (string, error) {
	if !xfile.Exists(pemPublicKeyFile) {
		return "", exception.Newf("公钥文件%s不存在", pemPublicKeyFile)
	}
	f, err := xfile.Open(pemPublicKeyFile)
	if err != nil {
		return "", err
	}
	pemKey, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}
	rawKey, err := DecodePem(pemKey)
	if err != nil {
		return "", err
	}
	return xbase64.UrlEncode(rawKey), nil
}

func PemPrivateKeyToBase64(pemKey string) (string, error) {
	rawKey, err := DecodePem([]byte(pemKey))
	if err != nil {
		return "", err
	}
	return xbase64.UrlEncode(rawKey), nil
}

func PemPrivateKeyFileToBase64(pemPrivateKeyFile string) (string, error) {
	if !xfile.Exists(pemPrivateKeyFile) {
		return "", exception.Newf("私钥文件%s不存在", pemPrivateKeyFile)
	}
	f, err := xfile.Open(pemPrivateKeyFile)
	if err != nil {
		return "", err
	}
	pemKey, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}
	rawKey, err := DecodePem(pemKey)
	if err != nil {
		return "", err
	}
	return xbase64.UrlEncode(rawKey), nil
}
