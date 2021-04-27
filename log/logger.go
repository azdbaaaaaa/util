package log

import (
	"fmt"
	"go.uber.org/zap"
)

var Sugar *zap.SugaredLogger
var Logger *zap.Logger

type LoggerOption struct {
	Development bool `json:"development" toml:"development" yaml:"development"`
}

func NewLogger(option LoggerOption) (logger *zap.Logger, sugar *zap.SugaredLogger, err error) {
	var zapLogger *zap.Logger
	if option.Development {
		zapLogger, err = zap.NewDevelopment()
	} else {
		zapLogger, err = zap.NewProduction()
	}
	if err != nil {
		panic(fmt.Sprintf("failed to new zap log,%v", err))
	}
	Sugar = zapLogger.Sugar()
	Logger = zapLogger
	return zapLogger, zapLogger.Sugar(), nil
}

func Infof(template string, args ...interface{}) {
	Sugar.Infof(template, args...)
}

func Infow(msg string, kws ...interface{}) {
	Sugar.Infow(msg, kws...)
}

func Warnf(template string, args ...interface{}) {
	Sugar.Warnf(template, args...)
}

func Warnw(msg string, kws ...interface{}) {
	Sugar.Warnw(msg, kws...)
}

func Errorf(template string, args ...interface{}) {
	Sugar.Errorf(template, args...)
}

func Errorw(msg string, kws ...interface{}) {
	Sugar.Errorw(msg, kws...)
}

func Panicf(template string, args ...interface{}) {
	Sugar.Panicf(template, args...)
}

func init() {
	var err error
	_, _, err = NewLogger(LoggerOption{Development: true})
	if err != nil {
		panic(err)
	}
}
