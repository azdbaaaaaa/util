package device

import (
	"encoding/base64"
	"github.com/azdbaaaaaa/util/log"
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestAesEncrypt(t *testing.T) {
	data := base64.StdEncoding.EncodeToString([]byte("imei|v1.0.0"))
	res, err := AesEncrypt([]byte(data), []byte(KEY_BASE64), []byte(ENCRYPT_IV))
	assert.Equal(t, err, nil)
	log.Infow(string(res))
}
