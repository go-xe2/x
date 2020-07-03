package xrsa

import "testing"

func TestNewRSA(t *testing.T) {
	rsa, err := NewRSA()
	if err != nil {
		t.Fatal(err)
	}
	mPriKey, _ := PrivateKeyToString(rsa.PrivateKey())
	mPubKey, _ := PublicKeyToString(rsa.PublicKey())
	t.Log("public key:", mPubKey)
	t.Log("private key:", mPriKey)
	t.Log("========================")
	v, err := PublicKeyToPem(rsa.PublicKey())
	t.Log("public pem key:", v, ", err:", err)

	v, err = PrivateKeyToPem(rsa.PrivateKey())
	t.Log("private pem key:", v, ", err:", err)

	data := "golang rsa test"
	encData, err := rsa.EncryptBase64(data)
	if err != nil {
		t.Error(err)
	}
	t.Log("encData:", encData)

	decData, err := rsa.DecryptBase64(encData)
	if err != nil {
		t.Error(err)
	}
	t.Log("decData:", decData)

	sign, err := rsa.SignBase64(data)
	if err != nil {
		t.Error(err)
	}
	t.Log("data sign:", sign)

	b, err := rsa.VerifyBase64(data, sign)
	if err != nil {
		t.Error(err)
	}
	t.Log("data:", data, ", sign:", sign, ", verify result:", b)
}
