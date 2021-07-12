package pprof

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func Register(on bool, r *gin.Engine) {
	if on {
		pprof.Register(r)
	}
}
