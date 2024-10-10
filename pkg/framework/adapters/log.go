package adapters

import (
	"context"
	"github.com/armineyvazi/jsonmap/pkg/framework/ports"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
)

type logger struct {
	zap *zap.Logger
}

func NewLogger(logLevel string, stackTrace bool) ports.Logger {
	zaplogLevel := getLogLevelFromEnv(logLevel)
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		MessageKey:     "msg",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // Ensure ISO8601 format
	stdoutEncoder := zapcore.NewJSONEncoder(encoderConfig)

	stdoutCore := zapcore.NewCore(
		stdoutEncoder,
		zapcore.Lock(os.Stdout),
		zaplogLevel,
	)

	core := zapcore.NewTee(stdoutCore)

	// Create the logger instance
	z := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	if stackTrace {
		z = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zap.ErrorLevel))
	}

	return &logger{
		zap: z,
	}
}

func (l *logger) Info(ctx context.Context, msg string, fields ...zap.Field) {
	l.zap.Info(msg, fields...)
}

func (l *logger) Error(ctx context.Context, msg string, fields ...zap.Field) {
	l.zap.Error(msg, fields...)
}

func (l *logger) Panic(ctx context.Context, msg string, fields ...zap.Field) {
	l.zap.Panic(msg, fields...)
}

func (l *logger) Fatal(ctx context.Context, msg string, fields ...zap.Field) {
	l.zap.Fatal(msg, fields...)
}

func (l *logger) Warn(ctx context.Context, msg string, fields ...zap.Field) {
	l.zap.Warn(msg, fields...)
}

func (l *logger) DPanic(ctx context.Context, msg string, fields ...zap.Field) {
	l.zap.DPanic(msg, fields...)
}

func (l *logger) Sync() {
	l.zap.Sync()
}

// getLogLevelFromEnv returns the zap log level based on the LOG_LEVEL environment variable.
func getLogLevelFromEnv(logLevel string) zapcore.Level {
	logLevelEnv := strings.ToLower(logLevel)

	switch logLevelEnv {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	case "dpanic":
		return zap.DPanicLevel
	case "panic":
		return zap.PanicLevel
	case "fatal":
		return zap.FatalLevel
	default:
		// Default to info level if no valid log level is found
		return zap.InfoLevel
	}
}
