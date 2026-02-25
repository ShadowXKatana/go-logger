package logger

import (
	"bytes"
	"log"
	"strings"
	"testing"
)

func newTestLogger(level Level) (*Logger, *bytes.Buffer) {
	buf := &bytes.Buffer{}
	l := &Logger{
		level:  level,
		logger: log.New(buf, "", 0),
	}
	return l, buf
}

func TestLevels(t *testing.T) {
	tests := []struct {
		name    string
		level   Level
		logFunc func(l *Logger)
		want    string
	}{
		{"debug", DEBUG, func(l *Logger) { l.Debug("hello") }, "[DEBUG] hello"},
		{"info", INFO, func(l *Logger) { l.Info("hello") }, "[INFO] hello"},
		{"warn", WARN, func(l *Logger) { l.Warn("hello") }, "[WARN] hello"},
		{"error", ERROR, func(l *Logger) { l.Error("hello") }, "[ERROR] hello"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l, buf := newTestLogger(DEBUG)
			tt.logFunc(l)
			got := strings.TrimSpace(buf.String())
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestLevelFiltering(t *testing.T) {
	l, buf := newTestLogger(WARN)
	l.Debug("should not appear")
	l.Info("should not appear")
	l.Warn("should appear")
	l.Error("should also appear")

	out := buf.String()
	if strings.Contains(out, "should not appear") {
		t.Errorf("expected messages below WARN to be filtered, got: %s", out)
	}
	if !strings.Contains(out, "should appear") {
		t.Errorf("expected WARN message to appear, got: %s", out)
	}
	if !strings.Contains(out, "should also appear") {
		t.Errorf("expected ERROR message to appear, got: %s", out)
	}
}

func TestFormatted(t *testing.T) {
	l, buf := newTestLogger(DEBUG)
	l.Infof("hello %s %d", "world", 42)
	got := strings.TrimSpace(buf.String())
	want := "[INFO] hello world 42"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestFormattedAllLevels(t *testing.T) {
	tests := []struct {
		name    string
		logFunc func(l *Logger)
		want    string
	}{
		{"debugf", func(l *Logger) { l.Debugf("val=%d", 1) }, "[DEBUG] val=1"},
		{"warnf", func(l *Logger) { l.Warnf("val=%d", 2) }, "[WARN] val=2"},
		{"errorf", func(l *Logger) { l.Errorf("val=%d", 3) }, "[ERROR] val=3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l, buf := newTestLogger(DEBUG)
			tt.logFunc(l)
			got := strings.TrimSpace(buf.String())
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	l := New(INFO)
	if l == nil {
		t.Fatal("New returned nil")
	}
	if l.level != INFO {
		t.Errorf("expected level INFO, got %v", l.level)
	}
}
