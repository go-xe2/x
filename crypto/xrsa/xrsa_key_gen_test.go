package xrsa

import (
	"github.com/go-xe2/x/xtest"
	"testing"
)

func TestGenRSAKeyPenString(t *testing.T) {
	xtest.Case(t, func() {
		//pri, pub, err := GenRSAKeyPemString(1024)
		//if err != nil {
		//	t.Error(err)
		//}
		//t.Log("privateKey:", pri)
		//t.Log("publicKey:", pub)

		pri1, pub1, err1 := GenRsaKeyBase64(1024)
		if err1 != nil {
			t.Error("genRsaKey error:", err1)
		}
		t.Log("base64 private key:", pri1)
		t.Log("base64 public key:", pub1)
	})
}
