package log

import (
	"fmt"
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestZapLogger_SetLevelDebugf(t *testing.T) {
	fmt.Println(123)
	//client := NewZapLogger(LoggerOption{
	//	Development: false,
	//	Level:       -1,
	//	StdoutPath:  "/tmp/test-out.log",
	//	StderrPath:  "/tmp/test-err.log",
	//})
	//client.SugaredLogger.Debugf("name:%s", "jimmy")
	//client.SugaredLogger.Errorf("err name:%s", "jimmy")
	fmt.Println(123)
	assert.Equal(t, 1, 0)
}

func TestMain(m *testing.M) {
	New(LoggerOption{
		Development: false,
		Level:       -1,
		StdoutPath:  "/tmp/test-out.log",
		StderrPath:  "/tmp/test-err.log",
	})
}
