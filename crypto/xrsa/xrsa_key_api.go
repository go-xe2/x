package xrsa

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/go-xe2/x/core/exception"
	"github.com/go-xe2/x/os/xfile"
	"io"
	"io/ioutil"
)

// 解码pem字符串数据
func DecodePem(pemData []byte) ([]byte, error) {
	block, _ := pem.Decode(pemData)
	if block == nil {
		return nil, exception.New("pem文件已损坏:\n" + string(pemData))
	}
	return block.Bytes, nil
}

// 解析rsa公钥, key为已经解码的pem数据
func ParsePublicKey(key []byte) (*rsa.PublicKey, error) {
	v, err := x509.ParsePKIXPublicKey(key)
	if err != nil {
		return nil, exception.Wrap(err, "公钥文件无效")
	}
	if pub, ok := v.(*rsa.PublicKey); ok {
		return pub, nil
	} else {
		return nil, exception.New("公钥文件无效")
	}
}

// 解析rsa公钥,key为pem编码的公钥字符串数据
func ParsePublicKeyPem(key []byte) (*rsa.PublicKey, error) {
	data, err := DecodePem(key)
	if err != nil {
		return nil, exception.Wrap(err, "公钥文件已经损坏")
	}
	return ParsePublicKey(data)
}

// 从编码为pem的阅读器中解析rsa公钥
func PublicKeyPemFromReader(reader io.Reader) (*rsa.PublicKey, error) {
	if reader == nil {
		return nil, exception.New("公钥阅读器不可用")
	}
	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, exception.Wrap(err, "公钥阅读器读取失败")
	}
	return ParsePublicKeyPem(bytes)
}

// 从编码为pem的文件解析rsa公钥
func PublicKeyPemFromFile(fileName string) (*rsa.PublicKey, error) {
	if !xfile.Exists(fileName) {
		return nil, exception.Newf("公钥文件%s不存在", fileName)
	}
	f, err := xfile.Open(fileName)
	if err != nil {
		return nil, exception.Wrap(err, "打开公钥文件失败")
	}
	return PublicKeyPemFromReader(f)
}

// 从编码为pem的阅读器解析rsa私钥
func PrivateKeyPemFromReader(reader io.Reader) (*rsa.PrivateKey, error) {
	if reader == nil {
		return nil, exception.New("私钥阅读器实例不可用")
	}
	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, exception.Wrap(err, "私钥阅读器读取失败")
	}
	return ParsePrivateKeyPem(bytes)
}

// 从编码为pem的文件解析rsa私钥
func PrivateKeyPemFromFile(fileName string) (*rsa.PrivateKey, error) {
	if !xfile.Exists(fileName) {
		return nil, exception.Newf("私钥文件%s不存在", fileName)
	}
	f, err := xfile.Open(fileName)
	if err != nil {
		return nil, exception.Wrap(err, "打开私钥文件失败")
	}
	return PrivateKeyPemFromReader(f)
}

// 从解码的pem数据解析rsa私钥
func ParsePrivateKey(key []byte) (*rsa.PrivateKey, error) {
	v, err := x509.ParsePKCS1PrivateKey(key)
	if err != nil {
		return nil, exception.Wrap(err, "私钥文件无效")
	}
	return v, nil
}

// 从pem解码rsa私钥
func ParsePrivateKeyPem(key []byte) (*rsa.PrivateKey, error) {
	data, err := DecodePem(key)
	if err != nil {
		return nil, exception.Wrap(err, "私钥文件已经损坏")
	}
	return ParsePrivateKey(data)
}
