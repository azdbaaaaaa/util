package in_param

import (
	metadata2 "github.com/azdbaaaaaa/util/net/metadata"
	"github.com/azdbaaaaaa/util/proto/common"
	"github.com/gin-gonic/gin"
)

const (
	HeaderUid  = "uid"
	HeaderUkey = "ukey"
)

func SetInParam(c *gin.Context) {
	inParam := metadata2.InParam{
		AppId:    common.AppIdType_APP_ID_LIGHTHOUSE,
		UserIp:   c.ClientIP(),
		ClientIp: c.Request.RemoteAddr,
	}
	c.Set(metadata2.ContextKeyInParam, inParam)
	c.Next()
}
