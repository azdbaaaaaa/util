package tars

import (
	"context"
	"github.com/azdbaaaaaa/util/net/metadata"
	"github.com/azdbaaaaaa/util/net/tars/common"
	tars_common "github.com/azdbaaaaaa/util/net/tars/common"
	common3 "github.com/azdbaaaaaa/util/proto/common"
)

type ClientConfig struct {
	Addr string `json:"addr"`
}

type Registry struct {
	Addr string `json:"addr"`
}

func NewTarsInParam(ctx context.Context) *common.InParam {
	inParam := &common.InParam{
		AppId:  tars_common.AppIdEnum_LIGHT_HOUSE,
		AreaId: 91,
	}

	// inparam
	if v := ctx.Value(metadata.ContextKeyInParam); v != nil {
		if metaParam, ok := v.(metadata.InParam); ok {
			inParam.AppId = int32(metaParam.AppId)
			inParam.ClientIp = metaParam.ClientIp
			inParam.UserIp = metaParam.UserIp
			inParam.UserId = metaParam.UserId
		}
	}
	// device
	if v := ctx.Value(metadata.ContextKeyDevice); v != nil {
		if device, ok := v.(metadata.Device); ok {
			inParam.DeviceId = device.IMEI
			inParam.Version = device.VersionCode
			inParam.Channel = device.Channel
			inParam.Source = device.Source
			inParam.Country = device.CountryCode
		}
	}
	if inParam.AppId == 0 {
		inParam.AppId = int32(common3.AppIdType_APP_ID_LIGHTHOUSE)
	}
	return inParam
}
