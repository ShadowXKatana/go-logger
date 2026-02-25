package logger

import (
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

// newTestLogger returns a Logger backed by an in-memory observer core so tests
// can inspect emitted entries without touching stdout.
func newTestLogger(level Level) (*Logger, *observer.ObservedLogs) {
	core, logs := observer.New(level)
	return &Logger{z: zap.New(core)}, logs
}

func TestLevels(t *testing.T) {
	tests := []struct {
		name    string
		logFunc func(l *Logger)
		wantMsg string
		wantLvl zapcore.Level
	}{
		{"debug", func(l *Logger) { l.Debug("hello") }, "hello", DEBUG},
		{"info", func(l *Logger) { l.Info("hello") }, "hello", INFO},
		{"warn", func(l *Logger) { l.Warn("hello") }, "hello", WARN},
		{"error", func(l *Logger) { l.Error("hello") }, "hello", ERROR},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l, logs := newTestLogger(DEBUG)
			tt.logFunc(l)
			if logs.Len() != 1 {
				t.Fatalf("expected 1 log entry, got %d", logs.Len())
			}
			entry := logs.All()[0]
			if entry.Message != tt.wantMsg {
				t.Errorf("message: got %q, want %q", entry.Message, tt.wantMsg)
			}
			if entry.Level != tt.wantLvl {
				t.Errorf("level: got %v, want %v", entry.Level, tt.wantLvl)
			}
		})
	}
}

func TestLevelFiltering(t *testing.T) {
	l, logs := newTestLogger(WARN)
	l.Debug("should not appear")
	l.Info("should not appear")
	l.Warn("should appear")
	l.Error("should also appear")

	for _, e := range logs.All() {
		if e.Message == "should not appear" {
			t.Errorf("expected messages below WARN to be filtered, got: %v", e.Message)
		}
	}
	if logs.Len() != 2 {
		t.Errorf("expected 2 log entries (WARN + ERROR), got %d", logs.Len())
	}
}

func TestFormatted(t *testing.T) {
	l, logs := newTestLogger(DEBUG)
	l.Infof("hello %s %d", "world", 42)
	if logs.Len() != 1 {
		t.Fatalf("expected 1 log entry, got %d", logs.Len())
	}
	want := "hello world 42"
	if logs.All()[0].Message != want {
		t.Errorf("got %q, want %q", logs.All()[0].Message, want)
	}
}

func TestFormattedAllLevels(t *testing.T) {
	tests := []struct {
		name    string
		logFunc func(l *Logger)
		wantMsg string
		wantLvl zapcore.Level
	}{
		{"debugf", func(l *Logger) { l.Debugf("val=%d", 1) }, "val=1", DEBUG},
		{"warnf", func(l *Logger) { l.Warnf("val=%d", 2) }, "val=2", WARN},
		{"errorf", func(l *Logger) { l.Errorf("val=%d", 3) }, "val=3", ERROR},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l, logs := newTestLogger(DEBUG)
			tt.logFunc(l)
			if logs.Len() != 1 {
				t.Fatalf("expected 1 log entry, got %d", logs.Len())
			}
			entry := logs.All()[0]
			if entry.Message != tt.wantMsg {
				t.Errorf("message: got %q, want %q", entry.Message, tt.wantMsg)
			}
			if entry.Level != tt.wantLvl {
				t.Errorf("level: got %v, want %v", entry.Level, tt.wantLvl)
			}
		})
	}
}

func TestNew(t *testing.T) {
	l := New(INFO)
	if l == nil {
		t.Fatal("New returned nil")
	}
	if l.z == nil {
		t.Fatal("New returned Logger with nil zap instance")
	}
}
