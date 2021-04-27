package log

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Sugar *zap.SugaredLogger
var Logger *zap.Logger

type LoggerOption struct {
	Level     zapcore.Level `json:"level" toml:"level" yaml:"level"`
	Output    string        `json:"output" toml:"output" yaml:"output" mapstructure:"output"`
	ErrOutput string        `json:"err_output" toml:"err_output" yaml:"err_output" mapstructure:"err_outputx"`
}

func NewLogger(option LoggerOption) (logger *zap.Logger, sugar *zap.SugaredLogger, err error) {
	var zapLogger *zap.Logger
	output := []string{"stdout"}
	errOutput := []string{"stderr"}
	if option.Output != "" {
		output = append(output, option.Output)
	}
	if option.ErrOutput != "" {
		errOutput = append(errOutput, option.ErrOutput)
	}
	config := zap.Config{
		Level:             zap.NewAtomicLevelAt(option.Level),
		Encoding:          "json",
		EncoderConfig:     zap.NewProductionEncoderConfig(),
		OutputPaths:       output,
		ErrorOutputPaths:  errOutput,
		Development:       false,
	}
	zapLogger, err = config.Build()
	//zapLogger, err = zap.NewProduction()
	if err != nil {
		panic(fmt.Sprintf("failed to new zap log,%v", err))
	}
	Logger = zapLogger
	Sugar = zapLogger.Sugar()
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
	_, _, err = NewLogger(LoggerOption{Level: zapcore.DebugLevel})
	if err != nil {
		panic(err)
	}
}
