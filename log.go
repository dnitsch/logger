package log

import (
	"fmt"
	"io"
	"strings"

	"github.com/rs/zerolog"
)

type LogLevel string

const (
	DebugLvl LogLevel = "debug"
	InfoLvl  LogLevel = "info"
	ErrorLvl LogLevel = "error"
)

var parseLvl = map[string]LogLevel{
	"debug": DebugLvl,
	"info":  InfoLvl,
	"error": ErrorLvl,
}

type Loggeriface interface {
	Debugf(format string, args ...any)
	Debug(msg string)
	Infof(format string, args ...any)
	Info(msg string)
	Errorf(format string, args ...any)
	Error(err error)
}

type Logger struct {
	Loggeriface
	logger zerolog.Logger
	lvl    LogLevel
	writer io.Writer
}

func New(writer io.Writer, lvl LogLevel) Logger {
	locallvl, zloglvl := lvl, zerolog.ErrorLevel

	zloglvl, err := zerolog.ParseLevel(string(lvl))
	if err != nil {
		fmt.Printf("loggerFailed: %v. setting error level", err)
		locallvl = ErrorLvl
	}

	zerolog.CallerMarshalFunc = lShortMarshall()
	zl := zerolog.New(writer).With().Timestamp().CallerWithSkipFrameCount(zerolog.CallerSkipFrameCount + 1).Logger().Level(zloglvl)

	return Logger{
		logger: zl,
		lvl:    locallvl,
		writer: writer,
	}
}

// Writer returns the current writer of the logging instance
func (l Logger) Writer() io.Writer {
	return l.writer
}

// Level returns the current logging level of the instance
func (l Logger) Level() LogLevel {
	return l.lvl
}

// ParseLevel returns a LogLevel if found by string
// If not found will default to error
// possible values are `error`, `debug`, `info`
func ParseLevel(lvl string) LogLevel {
	if flvl, found := parseLvl[lvl]; found {
		return flvl
	}
	return ErrorLvl
}

func (l Logger) Debugf(format string, args ...any) {
	l.logger.Debug().Msg(fmt.Sprintf(format, args...))
}

func (l Logger) Debug(msg string) {
	l.logger.Debug().Msg(msg)
}

func (l Logger) Infof(format string, args ...any) {
	l.logger.Info().Msg(fmt.Sprintf(format, args...))
}

func (l Logger) Info(msg string) {
	l.logger.Info().Msg(msg)
}

func (l Logger) Errorf(format string, args ...any) {
	l.logger.Error().Err(fmt.Errorf(format, args...)).Msg("")
}

func (l Logger) Error(err error) {
	l.logger.Error().Err(err).Msg("")
}

func lShortMarshall() func(pc uintptr, file string, line int) string {
	return func(pc uintptr, file string, line int) string {
		short := strings.Split(file, "/")
		if len(short) > 1 {
			return fmt.Sprintf("%s:%v", short[len(short)-1], line)
		}
		return fmt.Sprintf("%s:%v", file, line)
	}
}
