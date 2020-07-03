package xrsa

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/asn1"
	"encoding/pem"
	"fmt"
	"github.com/go-xe2/x/core/exception"
	"github.com/go-xe2/x/encoding/xbase64"
	"github.com/go-xe2/x/os/xfile"
	"io"
	"os"
)

func MarshalPKCS8PrivateKey(key *rsa.PrivateKey) []byte {
	info := struct {
		Version             int
		PrivateKeyAlgorithm []asn1.ObjectIdentifier
		PrivateKey          []byte
	}{}
	info.Version = 0
	info.PrivateKeyAlgorithm = make([]asn1.ObjectIdentifier, 1)
	info.PrivateKeyAlgorithm[0] = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 1, 1}
	info.PrivateKey = x509.MarshalPKCS1PrivateKey(key)
	k, _ := asn1.Marshal(info)
	return k
}

func GenRsaKey(keyLen ...int) (privateKey, publicKey []byte, err error) {
	nLen := 1024
	if len(keyLen) > 0 {
		nLen = keyLen[0]
	}
	priKey, err := rsa.GenerateKey(rand.Reader, nLen)
	if err != nil {
		return nil, nil, exception.New("私钥文件生成失败")
	}
	privateKey = x509.MarshalPKCS1PrivateKey(priKey)
	pubKey := &priKey.PublicKey
	publicKey, err = x509.MarshalPKIXPublicKey(pubKey)
	return
}

func PublicKeyToString(key *rsa.PublicKey) (string, error) {
	if key == nil {
		return "", exception.New("公钥对象实例无效")
	}
	v, err := x509.MarshalPKIXPublicKey(key)
	if err != nil {
		return "", err
	}
	return xbase64.UrlEncode(v), nil
}

func PrivateKeyToString(key *rsa.PrivateKey) (string, error) {
	if key == nil {
		return "", exception.New("私钥对象实例无效")
	}
	v := x509.MarshalPKCS1PrivateKey(key)
	return xbase64.UrlEncode(v), nil
}

// 公钥转码为pem格式
func PublicKeyToPem(key *rsa.PublicKey) (string, error) {
	if key == nil {
		return "", exception.New("公钥对象实例无效")
	}
	v, err := x509.MarshalPKIXPublicKey(key)
	if err != nil {
		return "", err
	}
	block := &pem.Block{
		Type:  "RSA Public key",
		Bytes: v,
	}
	out := bytes.NewBufferString("")
	err = pem.Encode(out, block)
	if err != nil {
		return "", err
	}
	return out.String(), nil
}

// 私钥转码为pem格式
func PrivateKeyToPem(key *rsa.PrivateKey) (string, error) {
	if key == nil {
		return "", exception.New("私钥对象实例无效")
	}
	v := x509.MarshalPKCS1PrivateKey(key)
	block := &pem.Block{
		Type:  "RSA Private key",
		Bytes: v,
	}
	out := bytes.NewBufferString("")
	err := pem.Encode(out, block)
	if err != nil {
		return "", err
	}
	return out.String(), nil
}

// 生成rsa公私钥pem
func GenRsaKeyPem(privateKeyOut io.Writer, publicKeyOut io.Writer, keyLen ...int) error {
	privateKey, publicKey, err := GenRsaKey(keyLen...)
	if err != nil {
		return err
	}
	block := &pem.Block{
		Type:  "RSA Private key",
		Bytes: privateKey,
	}
	err = pem.Encode(privateKeyOut, block)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	block = &pem.Block{
		Type:  "RSA Public key",
		Bytes: publicKey,
	}
	err = pem.Encode(publicKeyOut, block)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func GenRSAKeyPemString(bits ...int) (priKey, pubKey string, err error) {
	priWriter := bytes.NewBufferString("")
	pubWriter := bytes.NewBufferString("")
	err = GenRsaKeyPem(priWriter, pubWriter, bits...)
	priKey = priWriter.String()
	pubKey = pubWriter.String()
	return
}

func GenRSAKeyPemFile(dir string, bits ...int) error {
	if dir == "" {
		dir = "."
	}
	dir = xfile.RealPath(dir)
	if !xfile.Exists(dir) {
		xfile.Mkdir(dir)
	}
	privateFileName := xfile.Join(dir, "private_key.pem")
	privateFile, err := xfile.OpenFile(privateFileName, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	defer privateFile.Close()
	if err != nil {
		fmt.Println(err.Error())
		return err

	}
	publicFileName := xfile.Join(dir, "public_key.pem")
	publicKeyFile, err := xfile.OpenFile(publicFileName, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return GenRsaKeyPem(privateFile, publicKeyFile, bits...)
}
