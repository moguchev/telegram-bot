package logger

import (
	"log"

	"go.uber.org/zap"
)

var (
	engineLogger *zap.Logger
	sugarLogger  *zap.SugaredLogger
)

func init() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalln(err)
	}
	engineLogger = logger

	// defer  // flushes buffer, if any
	sugarLogger = logger.Sugar()

	sugarLogger.Debug()

	log.Println("log level: ", sugarLogger.Level().String())
}

// Close - flushes log buffer, if any
func Close() error {
	return engineLogger.Sync()
}

// Debug uses fmt.Sprint to construct and log a message.
func Debug(args ...interface{}) {
	sugarLogger.Debug(args...)
}

// Debugf uses fmt.Sprintf to log a templated message.
func Debugf(msg string, args ...interface{}) {
	sugarLogger.Debugf(msg, args...)
}

// DebugKV uses fmt.Sprintf to log a templated message.
func DebugKV(msg string, args ...interface{}) {
	sugarLogger.Debugw(msg, args...)
}

// Info uses fmt.Sprint to construct and log a message.
func Info(args ...interface{}) {
	sugarLogger.Info(args...)
}

// Infof uses fmt.Sprintf to log a templated message.
func Infof(msg string, args ...interface{}) {
	sugarLogger.Infof(msg, args...)
}

// InfoKV uses fmt.Sprintf to log a templated message.
func InfoKV(msg string, args ...interface{}) {
	sugarLogger.Infow(msg, args...)
}

// Warn uses fmt.Sprint to construct and log a message.
func Warn(args ...interface{}) {
	sugarLogger.Warn(args...)
}

// Warnf uses fmt.Sprintf to log a templated message.
func Warnf(msg string, args ...interface{}) {
	sugarLogger.Warnf(msg, args...)
}

// WarnKV uses fmt.Sprintf to log a templated message.
func WarnKV(msg string, args ...interface{}) {
	sugarLogger.Warnw(msg, args...)
}

// Error uses fmt.Sprint to construct and log a message.
func Error(args ...interface{}) {
	sugarLogger.Error(args...)
}

// Errorf uses fmt.Sprintf to log a templated message.
func Errorf(msg string, args ...interface{}) {
	sugarLogger.Errorf(msg, args...)
}

// ErrorKV uses fmt.Sprintf to log a templated message.
func ErrorKV(msg string, args ...interface{}) {
	sugarLogger.Errorw(msg, args...)
}

// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit.
func Fatal(args ...interface{}) {
	sugarLogger.Fatal(args...)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func Fatalf(msg string, args ...interface{}) {
	sugarLogger.Fatalf(msg, args...)
}

// FatalKV uses fmt.Sprintf to log a templated message.
func FatalKV(msg string, args ...interface{}) {
	sugarLogger.Fatalw(msg, args...)
}
