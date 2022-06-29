package v1

import (
	"fmt"
	"github.com/azdbaaaaaa/util/log"
	"github.com/go-playground/assert/v2"
	"go.uber.org/zap/zapcore"
	"testing"
)

func TestPaginationReq_GetIndex(t *testing.T) {
	log.New(log.LoggerOption{
		Development: false,
		Level:       -1,
		StdoutPath:  "/tmp/test-out.log",
		StderrPath:  "/tmp/test-err.log",
	})
	fmt.Println(log.Level())
	log.Debugf("xxx")
	log.Infof("eee")
	log.SetLevel(zapcore.InfoLevel)
	fmt.Println(log.Level())
	log.Debugf("xxx222")
	log.Infof("eee222")
	log.SetLevel(zapcore.DebugLevel)
	fmt.Println(log.Level())
	log.Debugf("xxx333")
	log.Infof("eee333")
	m := &PaginationReq{
		Page:     1,
		PageSize: 10,
	}
	from, to, err := m.GetIndex(5)
	assert.Equal(t, err, nil)
	assert.Equal(t, from, 0)
	assert.Equal(t, to, 5)
}
