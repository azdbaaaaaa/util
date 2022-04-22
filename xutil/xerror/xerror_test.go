package xerror

import (
	proto_common "github.com/azdbaaaaaa/util/proto/common"
	"github.com/go-playground/assert/v2"
	"testing"
)

type User struct {
}

func (u User) String() string {
	return ""
}

func TestNewProtoError(t *testing.T) {

	err := NewProtoError(proto_common.AppIdType_APP_ID_LIGHTHOUSE)
	err2 := NewProtoError(User{})
	assert.Equal(t, err, nil)
	assert.Equal(t, err2, nil)
}
