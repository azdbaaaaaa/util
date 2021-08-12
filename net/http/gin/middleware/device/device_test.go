package device

import (
	metadata2 "github.com/azdbaaaaaa/util/net/metadata"
	"github.com/go-playground/assert/v2"
	"testing"
)

var src = []byte("ffffffffb9dabf35000000000181c05b|1.6.2.5|720|1640|sdk|11|1|TECNO LC7|1625|sdk|1628672277370|0||||||GMT+08:00|Asia/Shanghai||1628672277371")

var dst = string("9IrOx3rJaTHSfrqqPxCY9JPJpzHqyOkKmPGQ+T2vuknR5gzfruaXITLb7AuDuuBt5l+1S5kpR/hFaZ+7qtFgMBftH6dGdPNjue3LvgQvS+ReNSleYQ7EsvnVdb78LVUu5EZ/0v46YHSTo4Dky9ONalDkqGPjCilJMBSHcU9+dVtrw3s2LwdBcZyQ9WhIxc7y")

func TestDecode(t *testing.T) {
	//dst = "A1bT6bpU9ZyrXYbFUAKS2w=="
	//wdtoken := "6owSdLMdAElY7xxSiQMVn4BFclqO9oRVuRxMrDPDtzQnErIvkBL7KXJYj3FBauaPsMH5M0eolTKcQYngjmHbsVBG93hIfLTrwTklniykUVtiD+7jF9+79lp7Iujjg0yqoyJbSzmLZBhFP78KwgdzzwU7BrkB+plhuCQNCwi6fRBg5irK8wiSmsKrRAKp52yl"
	decrypted, err := Decode(dst)
	assert.Equal(t, err, nil)
	d := metadata2.New(string(decrypted), "user_agent")
	assert.NotEqual(t, d, nil)
}