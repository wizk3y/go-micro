package logger

import (
	"go.uber.org/zap"
)

// Logger is a utility struct for logging data in an extremely high performance system.
// We can use both Logger and SugarLog for logging. For more information,
// just visit https://godoc.org/go.uber.org/zap
type Logger struct {
	sl *zap.SugaredLogger
}

// Debug -- uses fmt.Sprint to construct and log a message.
func (l *Logger) Debug(args ...interface{}) {
	l.sl.Debug(args...)
}

// Info -- uses fmt.Sprint to construct and log a message.
func (l *Logger) Info(args ...interface{}) {
	l.sl.Info(args...)
}

// Warn -- uses fmt.Sprint to construct and log a message.
func (l *Logger) Warn(args ...interface{}) {
	l.sl.Warn(args...)
}

// Error -- uses fmt.Sprint to construct and log a message.
func (l *Logger) Error(args ...interface{}) {
	l.sl.Error(args...)
}

// DPanic -- uses fmt.Sprint to construct and log a message. In development, the
// logger then panics. (See DPanicLevel for details.)
func (l *Logger) DPanic(args ...interface{}) {
	l.sl.DPanic(args...)
}

// Panic -- uses fmt.Sprint to construct and log a message, then panics.
func (l *Logger) Panic(args ...interface{}) {
	l.sl.Panic(args...)
}

// Fatal -- uses fmt.Sprint to construct and log a message, then calls os.Exit.
func (l *Logger) Fatal(args ...interface{}) {
	l.sl.Fatal(args...)
}

// Debugf -- uses fmt.Sprintf to log a templated message.
func (l *Logger) Debugf(template string, args ...interface{}) {
	l.sl.Debugf(template, args...)
}

// Infof -- uses fmt.Sprintf to log a templated message.
func (l *Logger) Infof(template string, args ...interface{}) {
	l.sl.Infof(template, args...)
}

// Warnf -- uses fmt.Sprintf to log a templated message.
func (l *Logger) Warnf(template string, args ...interface{}) {
	l.sl.Warnf(template, args...)
}

// Errorf -- uses fmt.Sprintf to log a templated message.
func (l *Logger) Errorf(template string, args ...interface{}) {
	l.sl.Errorf(template, args...)
}

// DPanicf -- uses fmt.Sprintf to log a templated message. In development, the
// logger then panics. (See DPanicLevel for details.)
func (l *Logger) DPanicf(template string, args ...interface{}) {
	l.sl.DPanicf(template, args...)
}

// Panicf -- uses fmt.Sprintf to log a templated message, then panics.
func (l *Logger) Panicf(template string, args ...interface{}) {
	l.sl.Panicf(template, args...)
}

// Fatalf -- uses fmt.Sprintf to log a templated message, then calls os.Exit.
func (l *Logger) Fatalf(template string, args ...interface{}) {
	l.sl.Fatalf(template, args...)
}

// Debugw -- logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
//
// When debug-level logging is disabled, this is much faster than
//  s.With(keysAndValues).Debug(msg)
func (l *Logger) Debugw(msg string, keysAndValues ...interface{}) {
	l.sl.Debugw(msg, keysAndValues...)
}

// Infow -- logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func (l *Logger) Infow(msg string, keysAndValues ...interface{}) {
	l.sl.Infow(msg, keysAndValues...)
}

// Warnw -- logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func (l *Logger) Warnw(msg string, keysAndValues ...interface{}) {
	l.sl.Warnw(msg, keysAndValues...)
}

// Errorw -- logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func (l *Logger) Errorw(msg string, keysAndValues ...interface{}) {
	l.sl.Errorw(msg, keysAndValues...)
}

// DPanicw -- logs a message with some additional context. In development, the
// logger then panics. (See DPanicLevel for details.) The variadic key-value
// pairs are treated as they are in With.
func (l *Logger) DPanicw(msg string, keysAndValues ...interface{}) {
	l.sl.DPanicw(msg, keysAndValues...)
}

// Panicw -- logs a message with some additional context, then panics. The
// variadic key-value pairs are treated as they are in With.
func (l *Logger) Panicw(msg string, keysAndValues ...interface{}) {
	l.sl.Panicw(msg, keysAndValues...)
}

// Fatalw -- logs a message with some additional context, then calls os.Exit. The
// variadic key-value pairs are treated as they are in With.
func (l *Logger) Fatalw(msg string, keysAndValues ...interface{}) {
	l.sl.Fatalw(msg, keysAndValues...)
}
