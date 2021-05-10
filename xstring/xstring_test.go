package xstring

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestJoin1(t *testing.T) {
	scopes := []string{"signature", "impersonation"}
	assert.Equal(t, Join(scopes, ","), "signature,impersonation")
}

func TestJoin2(t *testing.T) {
	scopes := []string{}
	assert.Equal(t, Join(scopes, ","), "")
}

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		scopes := []string{"signature", "impersonation"}
		Join(scopes, ",")
	}
}
