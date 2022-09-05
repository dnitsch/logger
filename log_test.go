package log

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

var (
	TestPhrase = "Want %v, Got: %v"
)

func Test_LogDebug(t *testing.T) {
	tests := []struct {
		name        string
		level       LogLevel
		logMethod   func(msg string)
		message     string
		expectEmpty bool
	}{
		{
			name:        "debug at debug",
			level:       DebugLvl,
			message:     "write me out...",
			expectEmpty: false,
		},
		{
			name:        "debug at info",
			level:       InfoLvl,
			message:     "write me out...",
			expectEmpty: true,
		},
		{
			name:        "debug at error",
			level:       ErrorLvl,
			message:     "write me out...",
			expectEmpty: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			l := New(&buf, tt.level)

			l.Debug(tt.message)
			s := buf.String()
			if tt.expectEmpty && s != "" {
				t.Errorf(TestPhrase, "", s)
			}
			if !tt.expectEmpty && !strings.Contains(s, `"level":"debug"`) {
				t.Errorf("incorrect level or not set in msg: %s", s)
			}
		})
	}
}

func Test_LogInfo(t *testing.T) {
	tests := []struct {
		name        string
		level       LogLevel
		logMethod   func(msg string)
		message     string
		expectEmpty bool
	}{
		{
			name:        "info at debug",
			level:       DebugLvl,
			message:     "write me out...",
			expectEmpty: false,
		},
		{
			name:        "info at info",
			level:       InfoLvl,
			message:     "write me out...",
			expectEmpty: false,
		},
		{
			name:        "info at error",
			level:       ErrorLvl,
			message:     "write me out...",
			expectEmpty: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			l := New(&buf, tt.level)

			l.Info(tt.message)
			s := buf.String()
			if tt.expectEmpty && s != "" {
				t.Errorf(TestPhrase, "", s)
			}
			if !tt.expectEmpty && !strings.Contains(s, `"level":"info"`) {
				t.Errorf("incorrect level or not set in msg: %s", s)
			}
		})
	}
}

func Test_LogError(t *testing.T) {
	tests := []struct {
		name        string
		level       LogLevel
		logMethod   func(msg string)
		message     string
		expectEmpty bool
	}{
		{
			name:        "error at debug",
			level:       DebugLvl,
			message:     "write me out...",
			expectEmpty: false,
		},
		{
			name:        "error at info",
			level:       InfoLvl,
			message:     "write me out...",
			expectEmpty: false,
		},
		{
			name:        "error at error",
			level:       ErrorLvl,
			message:     "write me out...",
			expectEmpty: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			l := New(&buf, tt.level)

			l.Error(fmt.Errorf(tt.message))
			s := buf.String()
			if tt.expectEmpty && s != "" {
				t.Errorf(TestPhrase, "", s)
			}
			if !tt.expectEmpty && !strings.Contains(s, `"level":"error"`) {
				t.Errorf("incorrect level or not set in msg: %s", s)
			}
		})
	}
}
