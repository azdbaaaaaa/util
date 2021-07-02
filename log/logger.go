package log

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Sugar *zap.SugaredLogger
var Logger *zap.Logger
var level zap.AtomicLevel
var ContextKeys = []string{"req_id", "trace_id"}

type LoggerOption struct {
	Development bool          `json:"development" toml:"development" yaml:"development"`
	Level       zapcore.Level `json:"level" toml:"level" yaml:"level"`
	StdoutPath  string        `json:"stdout_path" toml:"stdout_path" yaml:"stdout_path" mapstructure:"stdout_path"`
	StderrPath  string        `json:"stderr_path" toml:"stderr_path" yaml:"stderr_path" mapstructure:"stderr_path"`
	ContextKeys []string      `json:"context_keys" toml:"context_keys" yaml:"context_keys" mapstructure:"context_keys"`
}

func New(option LoggerOption) {
	var development, disableCaller, disableStacktrace bool
	disableCaller = false
	disableStacktrace = true
	var encoding = "json"

	stdoutPaths := []string{"stdout"}
	if option.StdoutPath != "" {
		stdoutPaths = append(stdoutPaths, option.StdoutPath)
	}

	stderrPaths := []string{"stderr"}
	if option.StderrPath != "" {
		stderrPaths = append(stderrPaths, option.StderrPath)
	}

	if option.Development {
		development = true
		disableCaller = false
		disableStacktrace = false
		encoding = "console"
	}
	if option.ContextKeys != nil {
		ContextKeys = option.ContextKeys
	}
	level = zap.NewAtomicLevelAt(option.Level)
	config := zap.Config{
		DisableCaller:     disableCaller,
		DisableStacktrace: disableStacktrace,
		Level:             level,
		Encoding:          encoding,
		EncoderConfig:     zap.NewProductionEncoderConfig(),
		OutputPaths:       stdoutPaths,
		ErrorOutputPaths:  stderrPaths,
		Development:       development,
	}
	zapLogger, err := config.Build()
	if err != nil {
		panic(fmt.Sprintf("failed to new zap log,%v", err))
	}
	Logger = zapLogger
	Sugar = zapLogger.Sugar()
	return
}

func Debugf(template string, args ...interface{}) {
	Sugar.Debugf(template, args...)
}

func Debugw(msg string, kws ...interface{}) {
	Sugar.Debugw(msg, kws...)
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

func SetLevel(lvl zapcore.Level) {
	level.SetLevel(lvl)
}

func WithContext(ctx context.Context) *zap.SugaredLogger {
	kvs := make([]interface{}, len(ContextKeys)*2, len(ContextKeys)*2)
	for _, k := range ContextKeys {
		v := ctx.Value(k)
		if v == nil {
			continue
		}
		kvs = append(kvs, k, v)
	}
	return Sugar.With(kvs...)
}

func init() {
	New(LoggerOption{Development: true, Level: zapcore.DebugLevel})
}
