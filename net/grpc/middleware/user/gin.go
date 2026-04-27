package user

import (
	"github.com/azdbaaaaaa/util/log"
	"github.com/azdbaaaaaa/util/proto/common"
	"github.com/gin-gonic/gin"
	"strconv"
)

const (
	HeaderUserID  = "uid"
	HeaderUserKey = "ukey"
	HeaderAppID   = "appId"
)

func SetUser(c *gin.Context) {
	user := User{
		AppId:   common.AppIdType_APP_ID_LIGHTHOUSE,
		UserKey: c.GetHeader(HeaderUserKey),
	}

	uidStr := c.GetHeader(HeaderUserID)
	if uidStr != "" {
		uid, err := strconv.ParseInt(uidStr, 64, 10)
		if err != nil {
			log.Errorw("uid invalid")
			c.Next()
			return
		}
		user.UserId = uid
	}

	c.Set(ContextKeyUser, user)
	c.Next()
}
