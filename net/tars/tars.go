package tars

import (
	"context"
	"github.com/azdbaaaaaa/util/net/metadata"
	"github.com/azdbaaaaaa/util/net/tars/common"
)

type ClientConfig struct {
	Addr string `json:"addr"`
}

type Registry struct {
	Addr string `json:"addr"`
}

func NewTarsInParam(ctx context.Context) *common.InParam {
	inParam := new(common.InParam)

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
	return inParam
}
