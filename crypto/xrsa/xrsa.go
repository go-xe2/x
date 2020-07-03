package xrsa

import (
	"crypto/rsa"
	"github.com/go-xe2/x/core/exception"
)

type TRSA struct {
	publicKey  *rsa.PublicKey
	privateKey *rsa.PrivateKey
}

// 创建rsa加解密对角实例
func NewRSA(priPubKey ...interface{}) (*TRSA, error) {
	var priKey *rsa.PrivateKey
	var pubKey *rsa.PublicKey
	for i := 0; i < len(priPubKey); i++ {
		if k, ok := priPubKey[i].(*rsa.PublicKey); pubKey == nil && ok {
			pubKey = k
		} else if k, ok := priPubKey[i].(*rsa.PrivateKey); priKey == nil && ok {
			priKey = k
		}
		if priKey != nil && pubKey != nil {
			break
		}
	}
	instance := &TRSA{}
	if priKey != nil && pubKey != nil {
		instance.privateKey = priKey
		instance.publicKey = pubKey
	} else {
		if len(priPubKey) > 0 {
			return nil, exception.New("请传入*rsa.PublicKey及*rsa.PrivateKey密钥对")
		}
		// 未传入时，创建新的密钥对
		_, _, err := instance.NewKeys()
		if err != nil {
			return nil, err
		}
	}
	return instance, nil
}

func NewRSAByPemFile(publicKeyFile, privateKeyFile string) (*TRSA, error) {
	pub, err := PublicKeyPemFromFile(publicKeyFile)
	if err != nil {
		return nil, err
	}
	pri, err := PrivateKeyPemFromFile(privateKeyFile)
	if err != nil {
		return nil, err
	}
	return &TRSA{
		privateKey: pri,
		publicKey:  pub,
	}, nil
}

// 从pem编码的密钥对创建rsa实例
func NewRSAByPem(publicKey []byte, privateKey []byte) (*TRSA, error) {
	pub, err := ParsePublicKeyPem(publicKey)
	if err != nil {
		return nil, err
	}
	pri, err := ParsePrivateKeyPem(privateKey)
	if err != nil {
		return nil, err
	}
	return &TRSA{
		publicKey:  pub,
		privateKey: pri,
	}, nil
}

// 从原始密码对数据创建rsa实例
func NewRSAByRaw(publicKey []byte, privateKey []byte) (*TRSA, error) {
	pub, err := ParsePublicKey(publicKey)
	if err != nil {
		return nil, err
	}
	pri, err := ParsePrivateKey(privateKey)
	if err != nil {
		return nil, err
	}
	return &TRSA{
		publicKey:  pub,
		privateKey: pri,
	}, nil
}

// 从base64编码的密钥对中创建rsa实例
func NewRSAByBase64(publicKey string, privateKey string) (*TRSA, error) {
	pub, err := ParseBase64PublicKey(publicKey)
	if err != nil {
		return nil, err
	}
	pri, err := ParseBase64PrivateKey(privateKey)
	if err != nil {
		return nil, err
	}
	return &TRSA{
		publicKey:  pub,
		privateKey: pri,
	}, nil
}

// 使用公钥加密数据
func (rsa *TRSA) Encrypt(data string) (string, error) {
	return Encrypt(data, rsa.publicKey)
}

// 使用公钥加密数据并以base64 urlEncode编码后返回
func (rsa *TRSA) EncryptBase64(data string) (string, error) {
	return EncryptBase64(data, rsa.publicKey)
}

// 使用私钥解密数据
func (rsa *TRSA) Decrypt(encData string) (string, error) {
	return Decrypt(encData, rsa.privateKey)
}

// 使用私钥解密base64编码的数据
func (rsa *TRSA) DecryptBase64(base64Data string) (string, error) {
	return DecryptBase64(base64Data, rsa.privateKey)
}

// 生成数据签名
func (rsa *TRSA) Sign(data string) (string, error) {
	return Sign(data, rsa.privateKey)
}

// 验证数据签名
func (rsa *TRSA) Verify(data string, sign string) (bool, error) {
	return Verify(data, sign, rsa.publicKey)
}

// 生成以base64编码的数据签名
func (rsa *TRSA) SignBase64(data string) (string, error) {
	return SignBase64(data, rsa.privateKey)
}

// 验证base64编码的数据签名
func (rsa *TRSA) VerifyBase64(data string, base64Sign string) (bool, error) {
	return VerifyBase64(data, base64Sign, rsa.publicKey)
}

// 获取公钥
func (rsa *TRSA) PublicKey() *rsa.PublicKey {
	return rsa.publicKey
}

// 获取私钥
func (rsa *TRSA) PrivateKey() *rsa.PrivateKey {
	return rsa.privateKey
}

// 生成新的密钥
func (rsa *TRSA) NewKeys() (privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey, err error) {
	rawPriKey, rawPubKey, err := GenRsaKey()
	if err != nil {
		return nil, nil, err
	}
	privateKey, err = ParsePrivateKey(rawPriKey)
	if err != nil {
		return nil, nil, err
	}
	publicKey, err = ParsePublicKey(rawPubKey)
	if err != nil {
		return nil, nil, err
	}
	rsa.publicKey = publicKey
	rsa.privateKey = privateKey
	return
}
