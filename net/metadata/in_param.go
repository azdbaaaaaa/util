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
	AppId common.AppIdType `json:"app_id,omitempty"`
	//AreaId   proto.AreaIdType `json:"area_id,omitempty"`
	//Version  int32            `json:"version,omitempty"`
	UserId int64 `json:"user_id,omitempty"`
	//DeviceId string           `json:"device_id,omitempty"`
	ClientIp string `json:"client_ip,omitempty"`
	//Channel  string           `json:"channel,omitempty"`
	//Source   string           `json:"source,omitempty"`
	UserIp string `json:"user_ip,omitempty"`
	//Language string           `json:"language,omitempty"`
	//Country  string           `json:"country,omitempty"`
	//AuthorId int64            `json:"author_id,omitempty"`
}

func InParamFromContext(ctx context.Context) (i InParam, err error) {
	if v := ctx.Value(ContextKeyInParam); v != nil {
		if i, ok := v.(InParam); ok {
			return i, nil
		}
	}
	return i, xerror.ErrNoInParamError
}
