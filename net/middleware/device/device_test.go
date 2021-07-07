package device

import (
	"encoding/base64"
	"fmt"
	"github.com/forgoer/openssl"
	"github.com/go-playground/assert/v2"
	"testing"
)

var src = []byte("imei|v1.0.0")

var dst = []byte("A1bT6bpU9ZyrXYbFUAKS2w==")

func TestAesCBCEncrypt(t *testing.T) {
	encrypted, err := openssl.AesCBCEncrypt(src, []byte(KEY_BASE64), []byte(ENCRYPT_IV), openssl.PKCS7_PADDING)
	assert.Equal(t, err, nil)
	assert.Equal(t, base64.StdEncoding.EncodeToString(encrypted), string(dst))
}

func TestAesCBCDecrypt(t *testing.T) {
	decoded, err := base64.StdEncoding.DecodeString(string(dst))
	assert.Equal(t, err, nil)
	decrypted, err := openssl.AesCBCDecrypt(decoded, []byte(KEY_BASE64), []byte(ENCRYPT_IV), openssl.PKCS7_PADDING)
	assert.Equal(t, err, nil)
	fmt.Println(string(decrypted)) // imei|v1.0.0
}
