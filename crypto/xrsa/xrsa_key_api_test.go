package xrsa

import (
	"testing"
)

const priKey = `-----BEGIN RSA Private key-----
MIICdAIBADALBgkqhkiG9w0BAQEEggJgMIICXAIBAAKBgQDvjXf7QHJ4MLFjWZYZ
+55Wqg55l0iwU6eiewTLmz4Bs/lgZYguV8mGBGmOHQQpUL3c6+/yLNjcRngkaEp+
CGUg5Q5dI/h3tdBjrmCd9YWrfpEUFDbJ966L6+y8Xwn5ofHu4wvvg2bMBlqZ3AW+
Vp82FjAWTwcoZMXdpo1Fr908GwIDAQABAoGANvxO9QIHBGWQ4hAXLyGdZ6EjL2+1
RJv/F4GLSTz6D6QC8sh7aL/HxxZOMEuiv+UDn6kca4c1w+H9A8ZYivNcWq3iCrT1
l32X8o5bUKzXCA5mChFAVQW1LbEPREclVAccTlvAAd22Xrc3w2xB301DbN2XujYf
AeqvP9We2LQ2ErkCQQD9VlkwBbs/hms2mZN332LfmsC7v38ay5cKVOqTPmJMVdxW
9p8G/qBn+qx9HHPGxZZ3Lq3Z62DxBs4gJxpnoUldAkEA8hIHrgidD5h4tZ9n4d9y
0WbOS3l7q2DMCMTX1JVNGE45hGbnl668e1xcDgDVoTaFnhORx4oWn+NMxkom8Q4r
1wJBAJRaPnw6vv5iTuJ4aL0n2ZSr4PWRHHOqiVaJh5yWSaX+GbvrTBEihvic+OmD
AeGCz6wXb8NPbUN4ArkdbP8GmhkCQCt9LcAIcB8jJ1yJ3OHpgPk53QoWMy+g3kcd
BiF9CTK6qv6sdiL8E4SeHTOu1rJ11x+FKIWOu23SKjLdk41vHFUCQD6quvgD/y8+
kLp77GqhvsVMLT8JpbhADd3GD2xjGoFoQGrt0C7NUPvjLMIVR/zfIepu+0D03TJf
yFuj8odqvp8=
-----END RSA Private key-----
`

const pubKey = `-----BEGIN RSA Public key-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDvjXf7QHJ4MLFjWZYZ+55Wqg55
l0iwU6eiewTLmz4Bs/lgZYguV8mGBGmOHQQpUL3c6+/yLNjcRngkaEp+CGUg5Q5d
I/h3tdBjrmCd9YWrfpEUFDbJ966L6+y8Xwn5ofHu4wvvg2bMBlqZ3AW+Vp82FjAW
TwcoZMXdpo1Fr908GwIDAQAB
-----END RSA Public key-----
`

func TestParseKey(t *testing.T) {
	sPriKey, err := DecodePem([]byte(priKey))
	if err != nil {
		t.Error(err)
	}
	t.Log("parse private key:", sPriKey)

	sPubKey, err := DecodePem([]byte(pubKey))
	if err != nil {
		t.Error(err)
	}
	t.Log("parse public key:", sPubKey)
}
