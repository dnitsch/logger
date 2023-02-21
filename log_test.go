package log_test

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"

	log "github.com/dnitsch/simplelog"
)

var (
	TestPhrase = `%s 
	got: %v
	want %v`
)

func TestParseLeve(t *testing.T) {
	ttests := map[string]struct {
		levelStr string
		expect   log.LogLevel
	}{
		"info": {
			"info", log.InfoLvl,
		},
		"debug": {
			"debug", log.DebugLvl,
		},
		"error": {
			"error", log.ErrorLvl,
		},
		"wrong": {
			"wrong", log.ErrorLvl,
		},
	}
	for name, tt := range ttests {
		t.Run(name, func(t *testing.T) {
			got := log.ParseLevel(tt.levelStr)
			if got != tt.expect {
				t.Errorf(TestPhrase, "log level incorrectly parsed", got, tt.expect)
			}
		})
	}
}

func TestLogNoFormat(t *testing.T) {
	tests := map[string]struct {
		level     log.LogLevel
		logMethod func(logger log.Logger) func(msg string)
		message   string
		expect    string
	}{
		"info at debug": {
			log.DebugLvl,
			func(logger log.Logger) func(msg string) {
				return logger.Info
			},
			"write me out...",
			`"message":"write me out..."`,
		},
		"info at info": {
			log.InfoLvl,
			func(logger log.Logger) func(msg string) {
				return logger.Info
			},
			"write me out...",
			`"message":"write me out..."`,
		},
		"info at error": {
			log.ErrorLvl,
			func(logger log.Logger) func(msg string) {
				return logger.Info
			},
			"write me out...",
			``,
		},
		"debug at error": {
			log.ErrorLvl,
			func(logger log.Logger) func(msg string) {
				return logger.Debug
			},
			"write me out...",
			``,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			var buf bytes.Buffer
			l := log.New(&buf, tt.level)
			if l.Writer() == nil {
				t.Errorf("expected Writer to not be <nil>")
			}

			if l.Level() == "" {
				t.Errorf("expected Level to not be \"\"")
			}

			(tt.logMethod(l))(tt.message)

			b, _ := io.ReadAll(&buf)
			if len(tt.expect) > 0 && string(b) != "" {
				if !strings.Contains(string(b), tt.expect) {
					t.Errorf(TestPhrase, "non formatted input", string(b), tt.expect)
				}
			}
		})
	}
}

func TestLogFormatted(t *testing.T) {
	tests := map[string]struct {
		level     log.LogLevel
		logMethod func(logger log.Logger) func(format string, args ...any)
		message   string
		expect    string
	}{
		"info at debug": {
			log.DebugLvl,
			func(logger log.Logger) func(format string, args ...any) {
				return logger.Infof
			},
			"write me out...",
			`"message":"write me out..."`,
		},
		"info at info": {
			log.InfoLvl,
			func(logger log.Logger) func(format string, args ...any) {
				return logger.Infof
			},
			"write me out...",
			`"message":"write me out..."`,
		},
		"info at error": {
			log.ErrorLvl,
			func(logger log.Logger) func(format string, args ...any) {
				return logger.Infof
			},
			"write me out...",
			``,
		},
		"debug at error": {
			log.ErrorLvl,
			func(logger log.Logger) func(format string, args ...any) {
				return logger.Debugf
			},
			"write me out...",
			``,
		},
		"error at debug": {
			log.DebugLvl,
			func(logger log.Logger) func(format string, args ...any) {
				return logger.Errorf
			},
			"write me out...",
			``,
		},
		"error at info": {
			log.InfoLvl,
			func(logger log.Logger) func(format string, args ...any) {
				return logger.Errorf
			},
			"write me out...",
			``,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			var buf bytes.Buffer
			l := log.New(&buf, tt.level)
			if l.Writer() == nil {
				t.Errorf("expected Writer to not be <nil>")
			}

			if l.Level() == "" {
				t.Errorf("expected Level to not be \"\"")
			}

			(tt.logMethod(l))(tt.message)

			b, _ := io.ReadAll(&buf)
			if len(tt.expect) > 0 && !strings.Contains(string(b), tt.expect) {
				t.Errorf(TestPhrase, "non formatted input", string(b), tt.expect)
			}
		})
	}
}

func TestLogError(t *testing.T) {
	tests := map[string]struct {
		level   log.LogLevel
		message string
		expect  string
	}{
		"error at debug": {
			level:   log.DebugLvl,
			message: "write me out...",
			expect:  `"message":"write me out..."`,
		},
		"error at info": {
			level:   log.InfoLvl,
			message: "write me out...",
			expect:  `"message":"write me out..."`,
		},
		"error at error": {
			level:   log.ErrorLvl,
			message: "write me out...",
			expect:  `"message":"write me out..."`,
		},
		"unknown w/LogR": {
			log.LogLevel("UNKNOWN"), "log me out", "",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			var buf bytes.Buffer
			l := log.New(&buf, tt.level)

			l.Error(fmt.Errorf(tt.message))
			s := buf.String()
			if len(tt.expect) < 1 && s != "" {
				t.Errorf(TestPhrase, "not empty", s, "")
			}
			if len(tt.expect) > 0 && !strings.Contains(s, `"level":"error"`) {
				t.Errorf(TestPhrase, "error not computed correctly", s, `"level":"error"`)

			}
		})
	}
}

func TestLogrImp(t *testing.T) {
	ttests := map[string]struct {
		level         log.LogLevel
		logrLevel     int
		message       string
		keysAndValues []any
		expect        string
	}{
		"info w/LogR": {
			log.InfoLvl, 0, "log me out", []any{"bar", "pair"}, `"message":"log me out"`,
		},
		"err w/LogR": {
			log.ErrorLvl, 1, "log me out", []any{"bar", "pair"}, ``,
		},
		"unknown w/LogR": {
			log.LogLevel("UNKNOWN"), 0, "log me out", []any{"bar", "pair"}, "",
		},
	}
	for name, tt := range ttests {
		t.Run(name, func(t *testing.T) {
			buf := &bytes.Buffer{}
			logr := log.NewLogr(buf, tt.level).WithCallDepth(-1)
			logr.V(tt.logrLevel).Info(tt.message, tt.keysAndValues...)

			b, _ := io.ReadAll(buf)

			if len(tt.expect) > 0 && !strings.Contains(string(b), tt.expect) {
				t.Errorf(TestPhrase, "non formatted input", string(b), tt.expect)
			}
		})
	}
}
