package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var defaultLogger *Logger

func init() {
	InitLoggerDefault()
}

// Logger is a utility struct for logging data in an extremely high performance system.
// We can use both Logger and SugarLog for logging. For more information,
// just visit https://godoc.org/go.uber.org/zap
type Logger struct {
	// configuration
	config   map[string]interface{}
	logLevel zap.AtomicLevel
	// Logger for logging
	Logger *zap.Logger
	// Sugar for logging
	*zap.SugaredLogger
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

	sugarLog := logger.Sugar()
	cfgParams := make(map[string]interface{})
	defaultLogger = &Logger{cfgParams, cfg.Level, logger, sugarLog}
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

	sugarLog := logger.Sugar()
	cfgParams := make(map[string]interface{})
	defaultLogger = &Logger{cfgParams, cfg.Level, logger, sugarLog}
}

// Debug -- uses fmt.Sprint to construct and log a message.
func (b *Logger) Debug(args ...interface{}) {
	b.SugaredLogger.Debug(args...)
}

// Info -- uses fmt.Sprint to construct and log a message.
func (b *Logger) Info(args ...interface{}) {
	b.SugaredLogger.Info(args...)
}

// Warn -- uses fmt.Sprint to construct and log a message.
func (b *Logger) Warn(args ...interface{}) {
	b.SugaredLogger.Warn(args...)
}

// Error -- uses fmt.Sprint to construct and log a message.
func (b *Logger) Error(args ...interface{}) {
	b.SugaredLogger.Error(args...)
}

// DPanic -- uses fmt.Sprint to construct and log a message. In development, the
// logger then panics. (See DPanicLevel for details.)
func (b *Logger) DPanic(args ...interface{}) {
	b.SugaredLogger.DPanic(args...)
}

// Panic -- uses fmt.Sprint to construct and log a message, then panics.
func (b *Logger) Panic(args ...interface{}) {
	b.SugaredLogger.Panic(args...)
}

// Fatal -- uses fmt.Sprint to construct and log a message, then calls os.Exit.
func (b *Logger) Fatal(args ...interface{}) {
	b.SugaredLogger.Fatal(args...)
}

// Debugf -- uses fmt.Sprintf to log a templated message.
func (b *Logger) Debugf(template string, args ...interface{}) {
	b.SugaredLogger.Debugf(template, args...)
}

// Infof -- uses fmt.Sprintf to log a templated message.
func (b *Logger) Infof(template string, args ...interface{}) {
	b.SugaredLogger.Infof(template, args...)
}

// Warnf -- uses fmt.Sprintf to log a templated message.
func (b *Logger) Warnf(template string, args ...interface{}) {
	b.SugaredLogger.Warnf(template, args...)
}

// Errorf -- uses fmt.Sprintf to log a templated message.
func (b *Logger) Errorf(template string, args ...interface{}) {
	b.SugaredLogger.Errorf(template, args...)
}

// DPanicf -- uses fmt.Sprintf to log a templated message. In development, the
// logger then panics. (See DPanicLevel for details.)
func (b *Logger) DPanicf(template string, args ...interface{}) {
	b.SugaredLogger.DPanicf(template, args...)
}

// Panicf -- uses fmt.Sprintf to log a templated message, then panics.
func (b *Logger) Panicf(template string, args ...interface{}) {
	b.SugaredLogger.Panicf(template, args...)
}

// Fatalf -- uses fmt.Sprintf to log a templated message, then calls os.Exit.
func (b *Logger) Fatalf(template string, args ...interface{}) {
	b.SugaredLogger.Fatalf(template, args...)
}

// Debugw -- logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
//
// When debug-level logging is disabled, this is much faster than
//  s.With(keysAndValues).Debug(msg)
func (b *Logger) Debugw(msg string, keysAndValues ...interface{}) {
	b.SugaredLogger.Debugw(msg, keysAndValues...)
}

// Infow -- logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func (b *Logger) Infow(msg string, keysAndValues ...interface{}) {
	b.SugaredLogger.Infow(msg, keysAndValues...)
}

// Warnw -- logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func (b *Logger) Warnw(msg string, keysAndValues ...interface{}) {
	b.SugaredLogger.Warnw(msg, keysAndValues...)
}

// Errorw -- logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func (b *Logger) Errorw(msg string, keysAndValues ...interface{}) {
	b.SugaredLogger.Errorw(msg, keysAndValues...)
}

// DPanicw -- logs a message with some additional context. In development, the
// logger then panics. (See DPanicLevel for details.) The variadic key-value
// pairs are treated as they are in With.
func (b *Logger) DPanicw(msg string, keysAndValues ...interface{}) {
	b.SugaredLogger.DPanicw(msg, keysAndValues...)
}

// Panicw -- logs a message with some additional context, then panics. The
// variadic key-value pairs are treated as they are in With.
func (b *Logger) Panicw(msg string, keysAndValues ...interface{}) {
	b.SugaredLogger.Panicw(msg, keysAndValues...)
}

// Fatalw -- logs a message with some additional context, then calls os.Exit. The
// variadic key-value pairs are treated as they are in With.
func (b *Logger) Fatalw(msg string, keysAndValues ...interface{}) {
	b.SugaredLogger.Fatalw(msg, keysAndValues...)
}

// SetLevel --
func (b *Logger) SetLevel(l zapcore.Level) {
	b.logLevel.SetLevel(l)
}

// GetLevel --
func (b *Logger) GetLevel() string {
	return b.logLevel.Level().String()
}
