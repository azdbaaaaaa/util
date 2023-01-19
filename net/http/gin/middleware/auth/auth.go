package auth

import (
	"github.com/azdbaaaaaa/util/net/metadata"
	"github.com/gin-gonic/gin"
)

func LoginRequired(c *gin.Context) {
	inParam, err := metadata.InParamFromContext(c)
	if err != nil {
		c.AbortWithStatus(500)
		return
	}
	if inParam.UserId == 0 {
		c.AbortWithStatus(401)
		return
	}
}
