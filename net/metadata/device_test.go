package metadata

import (
	"encoding/base64"
	"fmt"
	"github.com/azdbaaaaaa/util/net/http/gin/middleware/device"
	"github.com/forgoer/openssl"
	"github.com/go-playground/assert/v2"
	"testing"
)

var src = []byte("ffffffffd9eb0f22000000001a36c32f|1.6.2.5|720|1640|sdk_phoenix|11|1|TECNO LC7|1625|sdk_phoenix|1628672412294|0||||||GMT+08:00|Asia/Shanghai||1628672412295")

var dst = []byte("A1bT6bpU9ZyrXYbFUAKS2w==")

func TestAesCBCEncrypt(t *testing.T) {
	encrypted, err := openssl.AesCBCEncrypt(src, []byte(device.KEY_BASE64), []byte(device.ENCRYPT_IV), openssl.PKCS7_PADDING)
	assert.Equal(t, err, nil)
	assert.Equal(t, base64.StdEncoding.EncodeToString(encrypted), string(dst))
	fmt.Println(base64.StdEncoding.EncodeToString(encrypted))
}

func TestAesCBCDecrypt(t *testing.T) {
	decoded, err := base64.StdEncoding.DecodeString(string(dst))
	assert.Equal(t, err, nil)
	decrypted, err := openssl.AesCBCDecrypt(decoded, []byte(device.KEY_BASE64), []byte(device.ENCRYPT_IV), openssl.PKCS7_PADDING)
	assert.Equal(t, err, nil)
	fmt.Println(string(decrypted)) // imei|v1.0.0
}
