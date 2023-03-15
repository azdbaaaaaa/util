package auth

import (
	"context"
	"github.com/gin-gonic/gin"
	"strings"
)

func parseToken(token string) (struct{}, error) {
	return struct{}{}, nil
}

func userClaimFromToken(struct{}) string {
	return "foobar"
}

// AuthFunc is used by a middleware to authenticate requests
func AuthFunc(ctx context.Context) (context.Context, error) {
	return ctx, nil
	//token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	//if err != nil {
	//	return nil, err
	//}
	//tokenInfo, err := parseToken(token)
	//if err != nil {
	//	return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
	//}
	//
	//grpc_ctxtags.Extract(ctx).Set("auth.sub", userClaimFromToken(tokenInfo))
	//
	//// WARNING: in production define your own type to avoid context collisions
	//newCtx := context.WithValue(ctx, metadata2.ContextKeyUserID, tokenInfo)
	//
	//return newCtx, nil
}

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
