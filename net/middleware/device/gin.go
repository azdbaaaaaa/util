package device

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"github.com/azdbaaaaaa/util/log"
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

	sDec, err := base64.StdEncoding.DecodeString(hd)
	if err != nil {
		log.Errorw("base64 decode error", "err", err)
		c.Next()
		return
	}
	text, err := AesDecrypt(sDec, []byte(KEY_BASE64), []byte(ENCRYPT_IV))
	if err != nil {
		log.Errorw("AesDecrypt error", "err", err)
		c.Next()
		return
	}
	// 获取userAgent的header
	hua := c.GetHeader(HeaderUserAgent)
	c.Set(ContextKeyDevice, New(string(text), hua))
	c.Next()
	return
}

func PKCS5Padding(plaintext []byte, blockSize int) []byte {
	padding := blockSize - len(plaintext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(plaintext, padtext...)
}

//@brief:去除填充数据
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

//@brief:AES加密
func AesEncrypt(origData, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	//AES分组长度为128位，所以blockSize=16，单位字节
	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv[:blockSize]) //初始向量的长度必须等于块block的长度16字节
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

//@brief:AES解密
func AesDecrypt(crypted, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	//AES分组长度为128位，所以blockSize=16，单位字节
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, iv[:blockSize]) //初始向量的长度必须等于块block的长度16字节
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}
