package xrsa

import (
	"github.com/go-xe2/x/encoding/xbase64"
	"github.com/go-xe2/x/xtest"
	"testing"
)

const (
	base64PriKey = `MIICXQIBAAKBgQCmR6E5E-uY3hw3cRjAvxn5dlWJYszoeQrjKW-Wz3Jr8EVttCymQlwuTgoj6v6xO2AJgU2IMBhN1f21fcDp21NgLFQVfAatbRFWBIrAbiUXwR9ZrG7M5yj7qLPpRr5fl9LJmBWIPjRmuVoi3SCAduKlNe5OCNp3M-4sZoRgNS8lfwIDAQABAoGBAIcN-ucPRPZewPBPDdnP3dr-wC7cbq4LUkp7Z0VIrRj6iIm4I-POVmzNufK-dauiHDjcKwsIdVUTEASPZVcMa0SMrpqTVdtUYQ31oOxh1V5zVPLnVJaz3mTKH2QIpZmr3m48nb0QRm_nI9SFqDEfbHHwM75Ij1OvOWKRZw5V7i5RAkEA1JK42s-WNlP67xsGMKWa1apc1_7towSMGqTGlQTXGseZMHl2YkndI160Rz6urV9WWE-n7xYCpO1d9Dga5O5HSwJBAMg_1MuwY1qTrTbHriVdtLlFE3iDnnkMlobzOLOVOKWOkcxJvFviiKyT3qeYoKW70hR8Z3Zvam_pG2uun2Ct9h0CQAQr1ODGTgZG45epiheOSFmE_ElowTT_s9gZ_6OQ8r-dxw3CdGY9WM-G3ja_riHMyx70gTEZ13dxCPTv3Oc9Hb0CQQCp6getAjF7tN0AI3Tv_dAQeL1pv_zi57x-K7kMIG0dhZjPCC4MpW6lSR9fhFGj73f1rA26YBWnedurhlN0HIg9AkBG2iy5rEz7RxLwssAEDSsAdHDvCTED5Lhmf45RhjVOO97Vak-Ue3BAckUNx_qwcyICHxfldoo29EReXl3ZpX6P`

	base64PubKey = `MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCmR6E5E-uY3hw3cRjAvxn5dlWJYszoeQrjKW-Wz3Jr8EVttCymQlwuTgoj6v6xO2AJgU2IMBhN1f21fcDp21NgLFQVfAatbRFWBIrAbiUXwR9ZrG7M5yj7qLPpRr5fl9LJmBWIPjRmuVoi3SCAduKlNe5OCNp3M-4sZoRgNS8lfwIDAQAB`
)

func TestUtils(t *testing.T) {
	xtest.Case(t, func() {
		oPubKey, err := ParseBase64PublicKey(base64PubKey)

		xtest.Assert(err, nil)
		xtest.AssertNE(oPubKey, nil)

		oPriKey, err := ParseBase64PrivateKey(base64PriKey)
		xtest.Assert(err, nil)
		xtest.AssertNE(oPriKey, nil)

		pemPubKey, err := Base64PublicKeyToPem(base64PubKey)
		xtest.Assert(err, nil)
		t.Log("base64 to pem public key:", pemPubKey)

		_, err = Base64PrivateKeyToPem(base64PriKey)
		xtest.Assert(err, nil)
		//t.Log("base64 to pem private key:", pemPriKey)

		data := "hello golang rsa"
		t.Log("begin rsa encrypt data:", data)

		encData, err := EncryptBase64(data, oPubKey)
		xtest.Assert(err, nil)
		t.Log("encData:")
		t.Log(encData)

		rawPubKey, err := xbase64.UrlDecode(base64PubKey)
		xtest.Assert(err, nil)

		rawPriKey, err := xbase64.UrlDecode(base64PriKey)
		xtest.Assert(err, nil)

		t.Log("base64PubKey:")
		t.Log(base64PubKey)
		s := xbase64.UrlEncode(rawPubKey)
		t.Log("urlEncode raw pub key:")
		t.Log(s)

		encData1, err := EncryptBase64ByRawKey(data, rawPubKey)
		xtest.Assert(err, nil)
		t.Log("encData1")
		t.Log(encData1)
		//xtest.Assert(encData, encData1)

		decData, err := DecryptBase64(encData, oPriKey)
		xtest.Assert(err, nil)
		//xtest.Assert(data, decData)
		t.Log("dec data:", decData)

		decData1, err := DecryptBase64ByRawKey(encData1, rawPriKey)
		xtest.Assert(err, nil)
		//xtest.Assert(data, decData1)
		t.Log("dec data1:", decData1)
		xtest.Assert(data, decData)
		xtest.Assert(data, decData1)
	})
}
