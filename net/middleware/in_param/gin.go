package in_param

import (
	"github.com/azdbaaaaaa/util/net/middleware/device"
	"github.com/azdbaaaaaa/util/proto"
	"github.com/gin-gonic/gin"
)

func SetInParam(c *gin.Context) {
	inParam := proto.InParam{
		AppId:  int32(proto.AppIdEnum_LIGHTHOUSE), // TODO 默认设置为LIGHTHOUSE， SDK到时候要不要换
		UserIp: c.ClientIP(),
	}

	// 获取设备信息
	d, exists := c.Get(device.ContextKeyDevice)
	if exists {
		if dev, ok := d.(*device.Device); ok {
			inParam.AreaId = int32(dev.GetAreaID())
			inParam.ClientIp = c.Request.RemoteAddr
		}

	}
	c.Set(ContextKeyInParam, inParam)
	c.Next()
}
