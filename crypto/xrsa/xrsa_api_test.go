package xrsa

import "testing"

func TestRsa(t *testing.T) {
	pri, pub, err := GenRSAKeyPemString(1024)
	if err != nil {
		t.Fatal(err)
	}
	//t.Log("pub key:", pub)
	//t.Log("pri key:", pri)
	encryptData, err := EncryptBase64ByPemKey("go rsa demo", pub)
	if err != nil {
		t.Error("encrypt err:", err)
	}
	t.Log("encrypt data:", encryptData)

	decryptData, err := DecryptBase64ByPemKey(encryptData, pri)
	if err != nil {
		t.Error("decrypt err:", err)
	}
	t.Log("decrypt data:", decryptData)

	data := "this is needSign data"
	sign, err := SignByPemKey(data, pri)
	if err != nil {
		t.Error("data sign err:", err)
	}
	t.Log("data sign:", sign)

	b, err := VerifyByPemKey(data, sign, pub)
	if err != nil {
		t.Error("data verify sign err:", err)
	}
	t.Log("data verify sign result:", b)

	sign1, err := SignBase64ByPemKey(data, pri)
	if err != nil {
		t.Error("sign1 err:", err)
	}
	t.Log("sign1:", sign1)

	b, err = VerifyBase64ByPemKey(data, sign1, pub)
	if err != nil {
		t.Error("sign1 verify err:", err)
	}
	t.Log("sign1 verify result:", b)
}
