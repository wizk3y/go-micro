package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var defaultLogger *Logger

func init() {
	InitLoggerDefault()
}

// InitLoggerDefault -- init logger default
func InitLoggerDefault() {
	// init production encoder config
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	// init production config
	cfg := zap.NewProductionConfig()
	cfg.Sampling = nil
	cfg.EncoderConfig = encoderCfg
	cfg.OutputPaths = []string{"stdout"}
	cfg.ErrorOutputPaths = []string{"stdout"}
	// build logger
	logger, _ := cfg.Build()
	logger = logger.WithOptions(
		zap.AddCallerSkip(2),
	)

	defaultLogger = &Logger{
		sl: logger.Sugar(),
	}
}

// InitLoggerDefaultDev -- init logger dev
func InitLoggerDefaultDev() {
	// init production encoder config
	encoderCfg := zap.NewDevelopmentEncoderConfig()
	// init production config
	cfg := zap.NewDevelopmentConfig()
	cfg.Sampling = nil
	cfg.EncoderConfig = encoderCfg
	cfg.OutputPaths = []string{"stdout"}
	// build logger
	logger, _ := cfg.Build()
	logger = logger.WithOptions(
		zap.AddCallerSkip(2),
	)

	defaultLogger = &Logger{
		sl: logger.Sugar(),
	}
}
