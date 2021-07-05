package in_param

import (
	"github.com/azdbaaaaaa/util/proto"
	"github.com/gin-gonic/gin"
)

func SetInParam(c *gin.Context) {
	c.Set(ContextKeyInParam, proto.InParam{
		AreaId:   int32(proto.AppIdEnum_LIGHTHOUSE),
		ClientIp: c.ClientIP(),
		//UserId: c.Get(ContextKeyUser),
	})
	c.Next()
}
