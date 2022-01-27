package log

import (
	"context"
	"github.com/azdbaaaaaa/util/net/metadata"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sort"
	"time"
)

type zapLogger struct {
	Logger        *zap.Logger
	SugaredLogger *zap.SugaredLogger

	contextKeys []string
	option      LoggerOption
	config      zap.Config
}

func (log *zapLogger) SetLevel(level zapcore.Level) {
	log.config.Level = zap.NewAtomicLevelAt(level)
}

func New(option LoggerOption) *zapLogger {
	zLogger := &zapLogger{
		contextKeys: []string{metadata.ContextKeyReqID},
		option:      option,
	}

	if option.ContextKeys != nil {
		zLogger.contextKeys = append(zLogger.contextKeys, option.ContextKeys...)
	}
	if option.Development {
		zLogger.config = zap.Config{
			Level:             zap.NewAtomicLevelAt(option.Level),
			Development:       true,
			DisableCaller:     false,
			DisableStacktrace: false,
			Encoding:          "console",
		}
	} else {
		zLogger.config = zap.Config{
			Level:             zap.NewAtomicLevelAt(option.Level),
			Development:       false,
			DisableCaller:     false,
			DisableStacktrace: false,
			Encoding:          "json",
		}
	}

	logger = zLogger.build()
	return logger
}

func (log *zapLogger) build() *zapLogger {
	sink, errSink, outSink, err := log.openSinks()
	if err != nil {
		return nil
	}

	var encoder zapcore.Encoder
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	switch log.config.Encoding {
	case "json":
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	case "console":
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	default:
		panic("encoder not support")
	}
	logLevel := log.config.Level.Level()
	stdoutPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= logLevel && lvl < zapcore.ErrorLevel
	})
	stderrPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, sink, stdoutPriority),
		zapcore.NewCore(encoder, errSink, stderrPriority),
		zapcore.NewCore(encoder, outSink, logLevel),
	)
	log.Logger = zap.New(core, log.buildOptions(errSink)...)
	log.SugaredLogger = zap.New(core, log.buildOptions(errSink)...).Sugar()
	return log
}

func (log *zapLogger) openSinks() (zapcore.WriteSyncer, zapcore.WriteSyncer, zapcore.WriteSyncer, error) {
	if log.option.Logger == nil {
		log.option.Logger = &lumberjack.Logger{
			MaxSize:    20,
			MaxBackups: 30,
			MaxAge:     30,
			Compress:   false,
		}
	}
	sink := zapcore.AddSync(&lumberjack.Logger{
		Filename:   log.option.StdoutPath,
		MaxSize:    log.option.MaxSize,
		MaxBackups: log.option.MaxBackups,
		MaxAge:     log.option.MaxAge,
		Compress:   log.option.Compress,
	})

	errSink := zapcore.AddSync(&lumberjack.Logger{
		Filename:   log.option.StderrPath,
		MaxSize:    log.option.MaxSize,
		MaxBackups: log.option.MaxBackups,
		MaxAge:     log.option.MaxAge,
		Compress:   log.option.Compress,
	})

	outSink, _, err := zap.Open("stdout")
	if err != nil {
		return nil, nil, nil, err
	}
	//errSink, _, err := zap.Open([]string{log.option.StderrPath, "stderr"}...)
	//if err != nil {
	//	closeOut()
	//	return nil, nil, err
	//}
	return sink, errSink, outSink, nil
}

func (log *zapLogger) buildOptions(errSink zapcore.WriteSyncer) []zap.Option {
	opts := []zap.Option{zap.ErrorOutput(errSink)}

	if log.config.Development {
		opts = append(opts, zap.Development())
	}

	if !log.config.DisableCaller {
		opts = append(opts, zap.AddCaller(), zap.AddCallerSkip(1))
	}

	stackLevel := zap.ErrorLevel
	if log.config.Development {
		stackLevel = zap.WarnLevel
	}
	if !log.config.DisableStacktrace {
		opts = append(opts, zap.AddStacktrace(stackLevel))
	}

	if log.config.Sampling != nil {
		opts = append(opts, zap.WrapCore(func(core zapcore.Core) zapcore.Core {
			return zapcore.NewSampler(core, time.Second, int(log.config.Sampling.Initial), int(log.config.Sampling.Thereafter))
		}))
	}

	if len(log.config.InitialFields) > 0 {
		fs := make([]zap.Field, 0, len(log.config.InitialFields))
		keys := make([]string, 0, len(log.config.InitialFields))
		for k := range log.config.InitialFields {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			fs = append(fs, zap.Any(k, log.config.InitialFields[k]))
		}
		opts = append(opts, zap.Fields(fs...))
	}

	return opts
}

func (log *zapLogger) WithContext(ctx context.Context) *zap.SugaredLogger {
	kvs := make([]interface{}, 0, len(log.contextKeys)*2)
	for _, k := range log.contextKeys {
		v := ctx.Value(k)
		if v == nil {
			continue
		}
		kvs = append(kvs, k, v)
	}
	return log.SugaredLogger.With(kvs...)
}
