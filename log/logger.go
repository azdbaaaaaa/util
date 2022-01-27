package log

import (
	"context"
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
	// MaxSize is the maximum size in megabytes of the log file before it gets
	// rotated. It defaults to 100 megabytes.
	MaxSize int `json:"maxsize" yaml:"maxsize"`

	// MaxAge is the maximum number of days to retain old log files based on the
	// timestamp encoded in their filename.  Note that a day is defined as 24
	// hours and may not exactly correspond to calendar days due to daylight
	// savings, leap seconds, etc. The default is not to remove old log files
	// based on age.
	MaxAge int `json:"maxage" yaml:"maxage"`

	// MaxBackups is the maximum number of old log files to retain.  The default
	// is to retain all old log files (though MaxAge may still cause them to get
	// deleted.)
	MaxBackups int `json:"maxbackups" yaml:"maxbackups"`

	// LocalTime determines if the time used for formatting the timestamps in
	// backup files is the computer's local time.  The default is to use UTC
	// time.
	LocalTime bool `json:"localtime" yaml:"localtime"`

	// Compress determines if the rotated log files should be compressed
	// using gzip. The default is not to perform compression.
	Compress bool `json:"compress" yaml:"compress"`
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
