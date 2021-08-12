package device

import (
	"encoding/base64"
	"fmt"
	metadata2 "github.com/azdbaaaaaa/util/net/metadata"
	"github.com/forgoer/openssl"
	"github.com/go-playground/assert/v2"
	"testing"
)

var src = []byte("ffffffffd9eb0f22000000001a36c32f|1.6.2.5|720|1640|sdk|11|1|TECNO LC7|1625|sdk|1628672277370|0||||||GMT+08:00|Asia/Shanghai||1628672277371")

var dst = string("e17mw+tAoqZxUmnnyzFw0zzzaI2a7inFFjIb2zMt1OJrO1yGZVpcpOFuDSF+fETEUpWUDFvzu3YQbWuaZn2DkeaLEtPdiG7+NliL7LV9Ui2BnnRe506YGT5fv60O/bogXGwvF+kqhbyKWeIC1Idg/+R+P+iLwNCNocb6k9ZfzswMwoxcC1jsYQFrJ3H1bRH6 does not equal 9IrOx3rJaTHSfrqqPxCY9JPJpzHqyOkKmPGQ+T2vuknR5gzfruaXITLb7AuDuuBt5l+1S5kpR/hFaZ+7qtFgMBftH6dGdPNjue3LvgQvS+ReNSleYQ7EsvnVdb78LVUu5EZ/0v46YHSTo4Dky9ONalDkqGPjCilJMBSHcU9+dVtrw3s2LwdBcZyQ9WhIxc7y")

func TestDecode(t *testing.T) {
	//dst = "A1bT6bpU9ZyrXYbFUAKS2w=="
	//wdtoken := "6owSdLMdAElY7xxSiQMVn4BFclqO9oRVuRxMrDPDtzQnErIvkBL7KXJYj3FBauaPsMH5M0eolTKcQYngjmHbsVBG93hIfLTrwTklniykUVtiD+7jF9+79lp7Iujjg0yqoyJbSzmLZBhFP78KwgdzzwU7BrkB+plhuCQNCwi6fRBg5irK8wiSmsKrRAKp52yl"
	decrypted, err := Decode(dst)
	assert.Equal(t, err, nil)
	d := metadata2.New(string(decrypted), "user_agent")
	assert.NotEqual(t, d, nil)
}

func TestAesCBCEncrypt(t *testing.T) {
	encrypted, err := openssl.AesCBCEncrypt(src, []byte(KEY_BASE64), []byte(ENCRYPT_IV), openssl.PKCS7_PADDING)
	assert.Equal(t, err, nil)
	assert.Equal(t, base64.StdEncoding.EncodeToString(encrypted), string(dst))
	fmt.Println(base64.StdEncoding.EncodeToString(encrypted))
}