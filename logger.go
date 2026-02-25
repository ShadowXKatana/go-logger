package logger

import (
	"fmt"
	"log"
	"os"
)

// Level represents the logging level.
type Level int

const (
	DEBUG Level = iota
	INFO
	WARN
	ERROR
)

var levelNames = map[Level]string{
	DEBUG: "DEBUG",
	INFO:  "INFO",
	WARN:  "WARN",
	ERROR: "ERROR",
}

// Logger is a simple leveled logger.
type Logger struct {
	level  Level
	logger *log.Logger
}

// New creates a new Logger writing to os.Stdout at the given minimum level.
func New(level Level) *Logger {
	return &Logger{
		level:  level,
		logger: log.New(os.Stdout, "", log.LstdFlags),
	}
}

func (l *Logger) log(level Level, msg string) {
	if level >= l.level {
		l.logger.SetPrefix(fmt.Sprintf("[%s] ", levelNames[level]))
		l.logger.Print(msg)
	}
}

// Debug logs a message at DEBUG level.
func (l *Logger) Debug(msg string) { l.log(DEBUG, msg) }

// Info logs a message at INFO level.
func (l *Logger) Info(msg string) { l.log(INFO, msg) }

// Warn logs a message at WARN level.
func (l *Logger) Warn(msg string) { l.log(WARN, msg) }

// Error logs a message at ERROR level.
func (l *Logger) Error(msg string) { l.log(ERROR, msg) }

// Debugf logs a formatted message at DEBUG level.
func (l *Logger) Debugf(format string, args ...any) { l.log(DEBUG, fmt.Sprintf(format, args...)) }

// Infof logs a formatted message at INFO level.
func (l *Logger) Infof(format string, args ...any) { l.log(INFO, fmt.Sprintf(format, args...)) }

// Warnf logs a formatted message at WARN level.
func (l *Logger) Warnf(format string, args ...any) { l.log(WARN, fmt.Sprintf(format, args...)) }

// Errorf logs a formatted message at ERROR level.
func (l *Logger) Errorf(format string, args ...any) { l.log(ERROR, fmt.Sprintf(format, args...)) }
