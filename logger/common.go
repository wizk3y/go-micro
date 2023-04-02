package logger

// Debug --
func Debug(args ...interface{}) {
	defaultLogger.Debug(args...)
}

// Info --
func Info(args ...interface{}) {
	defaultLogger.Info(args...)
}

// Warn --
func Warn(args ...interface{}) {
	defaultLogger.Warn(args...)
}

// Error --
func Error(args ...interface{}) {
	defaultLogger.Error(args...)
}

// DPanic --
func DPanic(args ...interface{}) {
	defaultLogger.DPanic(args...)
}

// Panic --
func Panic(args ...interface{}) {
	defaultLogger.Panic(args...)
}

// Fatal --
func Fatal(args ...interface{}) {
	defaultLogger.Fatal(args...)
}

// Debugf --
func Debugf(template string, args ...interface{}) {
	defaultLogger.Debugf(template, args...)
}

// Infof --
func Infof(template string, args ...interface{}) {
	defaultLogger.Infof(template, args...)
}

// Warnf --
func Warnf(template string, args ...interface{}) {
	defaultLogger.Warnf(template, args...)
}

// Errorf --
func Errorf(template string, args ...interface{}) {
	defaultLogger.Errorf(template, args...)
}

// DPanicf --
func DPanicf(template string, args ...interface{}) {
	defaultLogger.DPanicf(template, args...)
}

// Panicf --
func Panicf(template string, args ...interface{}) {
	defaultLogger.Panicf(template, args...)
}

// Fatalf --
func Fatalf(template string, args ...interface{}) {
	defaultLogger.Fatalf(template, args...)
}

// Debugw --
func Debugw(msg string, keysAndValues ...interface{}) {
	defaultLogger.Debugw(msg, keysAndValues...)
}

// Infow --
func Infow(msg string, keysAndValues ...interface{}) {
	defaultLogger.Infow(msg, keysAndValues...)
}

// Warnw --
func Warnw(msg string, keysAndValues ...interface{}) {
	defaultLogger.Warnw(msg, keysAndValues...)
}

// Errorw --
func Errorw(msg string, keysAndValues ...interface{}) {
	defaultLogger.Errorw(msg, keysAndValues...)
}

// DPanicw --
func DPanicw(msg string, keysAndValues ...interface{}) {
	defaultLogger.DPanicw(msg, keysAndValues...)
}

// Panicw --
func Panicw(msg string, keysAndValues ...interface{}) {
	defaultLogger.Panicw(msg, keysAndValues...)
}

// Fatalw --
func Fatalw(msg string, keysAndValues ...interface{}) {
	defaultLogger.Fatalw(msg, keysAndValues...)
}
