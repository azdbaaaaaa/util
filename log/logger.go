package log

import (
	"context"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zapLogger

type LoggerOption struct {
	Development        bool          `json:"development"  yaml:"development"`
	Level              zapcore.Level `json:"level"  yaml:"level"`
	StdoutPath         string        `json:"stdout_path"  yaml:"stdout_path" mapstructure:"stdout_path"`
	StderrPath         string        `json:"stderr_path"  yaml:"stderr_path" mapstructure:"stderr_path"`
	ContextKeys        []string      `json:"context_keys"  yaml:"context_keys" mapstructure:"context_keys"`
	*lumberjack.Logger `json:",inline"`
}

func Logger() *zap.Logger {
	return logger.Logger
}

func SugaredLogger() *zap.SugaredLogger {
	return logger.SugaredLogger
}

func Debugf(template string, args ...interface{}) {
	logger.SugaredLogger.Debugf(template, args...)
}

func Debugw(msg string, kws ...interface{}) {
	logger.SugaredLogger.Debugw(msg, kws...)
}

func Infof(template string, args ...interface{}) {
	logger.SugaredLogger.Infof(template, args...)
}

func Infow(msg string, kws ...interface{}) {
	logger.SugaredLogger.Infow(msg, kws...)
}

func Warnf(template string, args ...interface{}) {
	logger.SugaredLogger.Warnf(template, args...)
}

func Warnw(msg string, kws ...interface{}) {
	logger.SugaredLogger.Warnw(msg, kws...)
}

func Errorf(template string, args ...interface{}) {
	logger.SugaredLogger.Errorf(template, args...)
}

func Errorw(msg string, kws ...interface{}) {
	logger.SugaredLogger.Errorw(msg, kws...)
}

func Panicf(template string, args ...interface{}) {
	logger.SugaredLogger.Panicf(template, args...)
}

func SetLevel(lvl zapcore.Level) {
	logger.SetLevel(lvl)
}

func WithContext(ctx context.Context) *zap.SugaredLogger {
	return logger.WithContext(ctx)
}

func init() {
	New(LoggerOption{Development: true, Level: zapcore.DebugLevel})
}
