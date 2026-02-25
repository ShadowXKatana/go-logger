package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Level represents the logging level (alias for zapcore.Level).
type Level = zapcore.Level

const (
	DEBUG = zapcore.DebugLevel
	INFO  = zapcore.InfoLevel
	WARN  = zapcore.WarnLevel
	ERROR = zapcore.ErrorLevel
)

// Logger wraps zap.Logger and emits structured JSON logs.
type Logger struct {
	z *zap.Logger
}

// New creates a new Logger writing JSON to stdout at the given minimum level.
func New(level Level) *Logger {
	cfg := zap.NewProductionConfig()
	cfg.Level = zap.NewAtomicLevelAt(level)
	z, _ := cfg.Build()
	return &Logger{z: z}
}

// Debug logs a message at DEBUG level.
func (l *Logger) Debug(msg string) { l.z.Debug(msg) }

// Info logs a message at INFO level.
func (l *Logger) Info(msg string) { l.z.Info(msg) }

// Warn logs a message at WARN level.
func (l *Logger) Warn(msg string) { l.z.Warn(msg) }

// Error logs a message at ERROR level.
func (l *Logger) Error(msg string) { l.z.Error(msg) }

// Debugf logs a formatted message at DEBUG level.
func (l *Logger) Debugf(format string, args ...any) { l.z.Sugar().Debugf(format, args...) }

// Infof logs a formatted message at INFO level.
func (l *Logger) Infof(format string, args ...any) { l.z.Sugar().Infof(format, args...) }

// Warnf logs a formatted message at WARN level.
func (l *Logger) Warnf(format string, args ...any) { l.z.Sugar().Warnf(format, args...) }

// Errorf logs a formatted message at ERROR level.
func (l *Logger) Errorf(format string, args ...any) { l.z.Sugar().Errorf(format, args...) }
