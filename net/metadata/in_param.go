package metadata

import (
	"context"
	"github.com/azdbaaaaaa/util/proto/common"
	"github.com/azdbaaaaaa/util/xutil/xerror"
)

const (
	ContextKeyInParam = "in_param"
)

type InParam struct {
	AppId    common.AppIdType `json:"app_id,omitempty"`
	UserIp   string           `json:"user_ip,omitempty"`
	UserId   int64            `json:"user_id,omitempty"`
	ClientIp string           `json:"client_ip,omitempty"`
	Token    string           `json:"token,omitempty"`
}

// InParamFromContext 只有在app网关层才会注入，如果直接内部调用会返回：ErrNoInParamError
func InParamFromContext(ctx context.Context) (i InParam, err error) {
	i = InParam{
		AppId:    common.AppIdType_APP_ID_LIGHTHOUSE,
		UserIp:   "",
		UserId:   0,
		ClientIp: "",
		Token:    "",
	}
	if v := ctx.Value(ContextKeyInParam); v != nil {
		if i, ok := v.(InParam); ok {
			return i, nil
		}
	}
	return i, xerror.ErrNoInParamError
}

