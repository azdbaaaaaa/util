package common

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
	"time"
)

var Logger *logrus.Logger

type MyFormatter struct{}

func (s *MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := time.Now().Local().Format("2006/01/02 - 15:04:05")
	msg := fmt.Sprintf("[Logger] %s [%s] %s-%d: %s\n", timestamp, strings.ToUpper(entry.Level.String()), entry.Caller.Function, entry.Caller.Line, entry.Message)
	return []byte(msg), nil
}

func init() {
	var logger = logrus.New()

	logger.Out = os.Stdout
	logger.SetReportCaller(true)
	logger.Formatter = &MyFormatter{}
	Logger = logger
}
