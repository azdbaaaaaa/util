package in_param

import (
	"github.com/azdbaaaaaa/util/proto"
	"github.com/gin-gonic/gin"
)

const (
	HeaderAppID  = "appId"
	HeaderAreaID = "areaId"
)

func SetInParam(c *gin.Context) {
	inParam := InParam{
		AppId:    proto.AppIdType_APP_ID_LIGHTHOUSE,
		UserIp:   c.ClientIP(),
		ClientIp: c.Request.RemoteAddr,
		//Language: "EN",
		//UserId: c.Get(ContextKeyUser),
	}
	//dev, exists := c.Get(device.ContextKeyDevice)
	//if exists {
	//	if d, ok := dev.(device.Device); ok {
	//		inParam.AreaId = d.GetAreaID()
	//		inParam.Version = int32(d.VersionCode)
	//		inParam.DeviceId = d.IMEI
	//		inParam.ClientIp = c.Request.RemoteAddr
	//		inParam.Channel = d.Channel
	//		inParam.Source = d.Source
	//	}
	//}
	c.Set(ContextKeyInParam, inParam)
	c.Next()
}
