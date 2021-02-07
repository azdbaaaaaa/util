package util

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm/logger"
	"time"
)

var Logger *zapLog

type zapLog struct {
	Log *zap.SugaredLogger
	LoggerOption
}

type LoggerOption struct {
	Development bool `json:"development" toml:"development" yaml:"development"`
}

func NewLogger(option LoggerOption) (Log *zap.SugaredLogger, err error) {
	var zapLogger *zap.Logger
	if option.Development {
		zapLogger, err = zap.NewDevelopment()
	} else {
		zapLogger, err = zap.NewProduction()
	}
	if err != nil {
		panic(fmt.Sprintf("failed to new zap log,%v", err))
	}
	Logger = &zapLog{Log: zapLogger.Sugar(), LoggerOption: option}
	return zapLogger.Sugar(), nil
}

func (l *zapLog) LogMode(level logger.LogLevel) *zap.SugaredLogger {
	return l.Log
}

func (l *zapLog) Info(ctx context.Context, template string, args ...interface{}) {
	l.Log.Infof(template, args...)
}

func (l *zapLog) Warn(ctx context.Context, template string, args ...interface{}) {
	l.Log.Warnf(template, args...)
}

func (l *zapLog) Error(ctx context.Context, template string, args ...interface{}) {
	l.Log.Errorf(template, args...)
}

func (l *zapLog) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	//sql, rows := fc()
	//if rows == -1 {
	//	l.logger.Infof(template, args...)
	//	l.Printf(l.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
	//} else {
	//	l.Printf(l.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
	//}
}

func (l *zapLog) Panic(ctx context.Context, template string, args ...interface{}) {
	l.Log.Panicf(template, args...)
}

func init() {
	var err error
	_, err = NewLogger(LoggerOption{Development: true})
	if err != nil {
		panic(err)
	}
}
