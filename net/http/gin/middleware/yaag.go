package middleware

import (
	"github.com/azdbaaaaaa/util/net/http/gin/middleware/yaag"
	"github.com/azdbaaaaaa/util/net/http/gin/middleware/yaag/models"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

func GinYaag() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !yaag.IsOn() {
			return
		}
		apiCall := models.ApiCall{}
		yaag.Before(&apiCall, c.Request)
		c.Next()
		if yaag.IsStatusCodeValid(c.Writer.Status()) {
			apiCall.MethodType = c.Request.Method
			apiCall.CurrentPath = strings.Split(c.Request.RequestURI, "?")[0]
			apiCall.ResponseBody = ""
			apiCall.ResponseCode = c.Writer.Status()
			headers := map[string]string{}
			for k, v := range c.Writer.Header() {
				log.Println(k, v)
				headers[k] = strings.Join(v, " ")
			}
			apiCall.ResponseHeader = headers
			go yaag.GenerateHtml(&apiCall)
		}
	}
}
