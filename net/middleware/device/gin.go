package device

import (
	"encoding/base64"
	"github.com/azdbaaaaaa/util/log"
	"github.com/forgoer/openssl"
	"github.com/gin-gonic/gin"
)

const (
	HeaderDevice    = "wdtoken"
	HeaderUserAgent = "user-agent"

	KEY_BASE64 = "ZXN1b2h0aGdpbEZDSU9PTA=="
	ENCRYPT_IV = "ficoollighthouse"
)

func Decode(dst string) (decrypted []byte, err error) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorw("AES Decode error", "err", err)
		}
	}()
	decoded, err := base64.StdEncoding.DecodeString(dst)
	if err != nil {
		log.Errorw("base64.StdEncoding.DecodeString error", "err", err, "dst", dst)
		return
	}
	decrypted, err = openssl.AesCBCDecrypt(decoded, []byte(KEY_BASE64), []byte(ENCRYPT_IV), openssl.PKCS7_PADDING)
	if err != nil {
		log.Errorw("openssl.AesCBCDecrypt error", "err", err, "dst", dst)
		return
	}
	return
}

func SetDevice(c *gin.Context) {
	// 获取设备信息的header
	hd := c.GetHeader(HeaderDevice)
	if len(hd) > 0 {
		decrypted, err := Decode(hd)
		if err != nil {
			c.Next()
			return
		}
		// 获取userAgent的header
		hua := c.GetHeader(HeaderUserAgent)
		c.Set(ContextKeyDevice, New(string(decrypted), hua))
	}
	c.Next()
	return
}
