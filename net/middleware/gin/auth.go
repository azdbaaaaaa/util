package gin

import (
	"github.com/gin-gonic/gin"
	"strings"
)

const ContextKeyUser = "user"

//func AuthRequired(svc UserService) gin.HandlerFunc {
//	return func(ctx *gin.Context) {
//		authorization := strings.Split(ctx.GetHeader("Authorization"), "Bearer ")
//		if len(authorization) < 2 {
//			ctx.AbortWithStatusJSON(http.StatusUnauthorized, nil)
//			return
//		}
//		token := authorization[1]
//		if token == "" {
//			ctx.AbortWithStatusJSON(http.StatusUnauthorized, nil)
//			return
//		}
//		user, err := svc.GetUserByToken(token)
//		if err != nil {
//			ctx.AbortWithStatusJSON(http.StatusUnauthorized, nil)
//			return
//		}
//		ctx.Set("user", user)
//		ctx.Next()
//		return
//	}
//}

// ParseToken 用户名注入到ctx中
func ParseToken() gin.HandlerFunc {
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
