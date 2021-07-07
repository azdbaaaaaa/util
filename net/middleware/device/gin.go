package device

import (
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

func SetDevice(c *gin.Context) {
	// 获取设备信息的header
	hd := c.GetHeader(HeaderDevice)
	if len(hd) > 0 {
		dst, err := openssl.AesCBCDecrypt([]byte(hd), []byte(KEY_BASE64), []byte(ENCRYPT_IV), openssl.PKCS7_PADDING)
		if err != nil {
			log.Errorw("openssl.AesCBCDecrypt error", "err", err, "hd", hd)
			c.Next()
			return
		}
		// 获取userAgent的header
		hua := c.GetHeader(HeaderUserAgent)
		c.Set(ContextKeyDevice, New(string(dst), hua))
	}
	c.Next()
	return
}
