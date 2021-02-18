package util

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strings"
)

//type UserService interface {
//	GetUserByToken(token string) (string, error)
//}
//
///*
//用户认证校验器
//*/
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

const ContextKeyUser = "user"

// 用户名注入到ctx中
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
		user, err := getValues(token, []string{"name", "_user"})
		if err != nil {
			ctx.Next()
			return
		}
		ctx.Set(ContextKeyUser, user)
		ctx.Next()
		return
	}
}

// 从ctx中获取用户名
func GetUsername(ctx context.Context) (string, error) {
	user := ctx.Value(ContextKeyUser)
	username, ok := user.(string)
	if !ok {
		return "", CodeCtxKeyInvalid
	}
	return username, nil
}

func getValue(token string, key string) (interface{}, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, CodeTokenError
	}
	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, CodeTokenError
	}
	m := make(map[string]interface{})
	if err := json.Unmarshal(payload, &m); err != nil {
		return nil, CodeTokenError
	}
	if v, ok := m[key]; ok {
		return v, nil
	}
	return nil, CodeTokenError
}

func getValues(token string, keys []string) (interface{}, error) {
	for _, key := range keys {
		v, err := getValue(token, key)
		if err != nil {
			continue
		}
		return v, nil
	}
	return nil, CodeTokenError
}
