package middleware

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func ErrHandler(err error) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorization := strings.Split(ctx.GetHeader("Authorization"), "Bearer ")
		if len(authorization) != 2 {
			ctx.Next()
			return
		}
		token := authorization[1]
		if token == "" {
			ctx.Next()
			return
		}
		//user, err := getValues(token, []string{"name", "_user", "sub"})
		//if err != nil {
		//	ctx.Next()
		//	return
		//}
		//ctx.Set(ContextKeyUser, user)
		ctx.Next()
		return
	}
}
