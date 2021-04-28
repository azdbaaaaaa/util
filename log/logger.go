package log

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Sugar *zap.SugaredLogger
var Logger *zap.Logger

type LoggerOption struct {
	Development bool          `json:"development" toml:"development" yaml:"development"`
	Level       zapcore.Level `json:"level" toml:"level" yaml:"level"`
	StdoutPaths []string      `json:"stdout_paths" toml:"stdout_paths" yaml:"stdout_paths" mapstructure:"stdout_paths"`
	StderrPaths []string      `json:"stderr_paths" toml:"stderr_paths" yaml:"stderr_paths" mapstructure:"stderr_paths"`
}

func NewLogger(option LoggerOption) {
	var development bool
	stdoutPaths := []string{"stdout"}
	if len(option.StdoutPaths) > 0 {
		for _, stdout := range option.StdoutPaths {
			stdoutPaths = append(stdoutPaths, stdout)
		}
	}

	stderrPaths := []string{"stderr"}
	if len(option.StderrPaths) > 0 {
		for _, stderr := range option.StderrPaths {
			stderrPaths = append(stderrPaths, stderr)
		}
	}

	if option.Development {
		development = true
	}
	config := zap.Config{
		DisableCaller:     true,
		DisableStacktrace: true,
		Level:             zap.NewAtomicLevelAt(option.Level),
		Encoding:          "json",
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
	NewLogger(LoggerOption{Development: true, Level: zapcore.DebugLevel})
}
