package device

import (
	"encoding/base64"
	"fmt"
	"github.com/azdbaaaaaa/util/log"
	metadata2 "github.com/azdbaaaaaa/util/net/metadata"
	"github.com/forgoer/openssl"
	"github.com/go-playground/assert/v2"
	"testing"
)

var src = []byte("000000003fe4c136ffffffffe5f22f86|1.6.2.5|720|1640|sdk_phoenix|11|1|TECNO LC7|1625|sdk_phoenix|1628672277370|0||||||GMT+08:00|Asia/Shanghai||1628672277371")

var dst = string("OX7SLZjHiuc+bb8qMElUk1me8Uyr9R63k+fni3dvLGMrQbrh24JFz2KO98mmvuGnCmBCLDGUmHELJ17HXESSGSvbuAumIbW5fmIomeCtZUQ1v8vEs0bCUuGJnaqC4Osd4LxTGOm+/sa6W6/aD8GR/B5vP0bE9VYXuiyXx2J6YWJHyvJhRlEk17MnXVKh242yaMdMMcFrxcdM6IS95/IXaIRGJtMKiMvsvYzbwJHrcDnKD8Ub9ZIZvE/Yn8KAxw+9")

func TestDecode(t *testing.T) {
	//dst = "A1bT6bpU9ZyrXYbFUAKS2w=="
	//wdtoken := "6owSdLMdAElY7xxSiQMVn4BFclqO9oRVuRxMrDPDtzQnErIvkBL7KXJYj3FBauaPsMH5M0eolTKcQYngjmHbsVBG93hIfLTrwTklniykUVtiD+7jF9+79lp7Iujjg0yqoyJbSzmLZBhFP78KwgdzzwU7BrkB+plhuCQNCwi6fRBg5irK8wiSmsKrRAKp52yl"
	decrypted, err := Decode(dst)
	assert.Equal(t, err, nil)
	d := metadata2.New(string(decrypted))
	assert.NotEqual(t, d, nil)
}

func TestAesCBCEncrypt(t *testing.T) {
	key, err := base64.StdEncoding.DecodeString(KEY_BASE64)
	if err != nil {
		log.Errorw("base64.StdEncoding.DecodeString KEY_BASE64 error", "err", err, "dst", dst)
		return
	}
	encrypted, err := openssl.AesCBCEncrypt(src, key, []byte(ENCRYPT_IV), openssl.PKCS5_PADDING)
	assert.Equal(t, err, nil)
	token := base64.StdEncoding.EncodeToString(encrypted)
	fmt.Println(token)
	assert.Equal(t, token, string(dst))
}