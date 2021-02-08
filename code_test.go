package util

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestDefaultCode(t *testing.T) {
	assert.Equal(t, CodeSuccess.Code(), 8000)
	assert.Equal(t, CodeSuccess.Message(), "成功")

	assert.Equal(t, CodeErrorParam.Code(), 8010)
	assert.Equal(t, CodeErrorParam.Message(), "参数不正确")

	assert.Equal(t, CodeErrorToken.Code(), 9001)
	assert.Equal(t, CodeErrorToken.Message(), "令牌已过期，请重新获取令牌")
}

func TestNewCode(t *testing.T) {
	var CodeTest = NewCode(100000, "测试")
	assert.Equal(t, CodeTest.Code(), 100000)
	assert.Equal(t, CodeTest.Message(), "测试")
}
