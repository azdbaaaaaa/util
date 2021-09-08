package tars

type ClientConfig struct {
	Addr string `json:"addr"`
}

type Registry struct {
	Addr string `json:"addr"`
}

//func NewTarsInParam(ctx context.Context) *common.InParam {
//	inParam := new(common.InParam)
//
//	// inparam
//	if v := ctx.Value(metadata.ContextKeyInParam); v != nil {
//		if metaParam, ok := v.(metadata.InParam); ok {
//			inParam.AppId = int32(metaParam.AppId)
//			inParam.ClientIp = metaParam.ClientIp
//			inParam.UserIp = metaParam.UserIp
//		}
//	}
//	// device
//	if v := ctx.Value(metadata.ContextKeyDevice); v != nil {
//		if metaParam, ok := v.(metadata.Device); ok {
//			inParam.DeviceId = metaParam.IMEI
//		}
//	}
//	return inParam
//}
