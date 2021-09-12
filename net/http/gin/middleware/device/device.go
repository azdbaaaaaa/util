package device

import (
	"encoding/base64"
	"github.com/azdbaaaaaa/util/log"
	metadata2 "github.com/azdbaaaaaa/util/net/metadata"
	"github.com/azdbaaaaaa/util/xutil/xerror"
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
	key, err := base64.StdEncoding.DecodeString(KEY_BASE64)
	if err != nil {
		log.Errorw("base64.StdEncoding.DecodeString KEY_BASE64 error", "err", err, "dst", dst)
		return
	}
	decrypted, err = openssl.AesCBCDecrypt(decoded, key, []byte(ENCRYPT_IV), openssl.PKCS5_PADDING)
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
			log.Errorw("SetDevice.Decode", "err", err, "device_header", hd)
			err = xerror.ErrDeviceInvalidError
			c.Abort()
		}
		d := metadata2.New(string(decrypted))
		// 获取userAgent的header
		userAgent := c.GetHeader(HeaderUserAgent)
		d.UserAgent = userAgent
		c.Set(metadata2.ContextKeyDevice, *d)
	}
	c.Next()
}
