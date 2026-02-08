package logger

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func SetupZapLogger(env string) *zap.Logger {

	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeDuration = zapcore.StringDurationEncoder
	cfg.Encoding = "json"
	cfg.OutputPaths = []string{"stdout"}
	cfg.ErrorOutputPaths = []string{"stderr"}
	cfg.EncoderConfig.TimeKey = "timestamp"
	cfg.EncoderConfig.LevelKey = "level"
	cfg.EncoderConfig.MessageKey = "msg"
	cfg.EncoderConfig.CallerKey = "caller"

	switch env {
	case envLocal:
		cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	case envDev:
		cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	case envProd:
		cfg.Level = zap.NewAtomicLevelAt(zap.WarnLevel)

	}

	logger, err := cfg.Build(
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)
	if err != nil {
		log.Fatal(err)
	}

	logger = logger.With(
		zap.String("service", "auth-service"),
		zap.String("stage", env),
	)
	return logger
}
