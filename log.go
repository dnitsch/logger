package log

import (
	"fmt"
	"io"

	"github.com/rs/zerolog"
)

type LogLevel string

const (
	DebugLvl LogLevel = "debug"
	InfoLvl  LogLevel = "info"
	ErrorLvl LogLevel = "error"
)

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
}

func New(writer io.Writer, lvl LogLevel) Logger {
	locallvl, zloglvl := lvl, zerolog.ErrorLevel

	zloglvl, err := zerolog.ParseLevel(string(lvl))
	if err != nil {
		fmt.Printf("loggerFailed: %v. setting error level", err)
		locallvl = ErrorLvl
	}
	return Logger{
		logger: zerolog.New(writer).With().Timestamp().Logger().Level(zloglvl),
		lvl:    locallvl,
	}
}

func (l Logger) Debugf(format string, args ...any) {
	l.Debug(fmt.Sprintf(format, args...))
}

func (l Logger) Debug(msg string) {
	l.logger.Debug().Msg(msg)
}

func (l Logger) Infof(format string, args ...any) {
	l.Info(fmt.Sprintf(format, args...))
}

func (l Logger) Info(msg string) {
	l.logger.Info().Msg(msg)
}

func (l Logger) Errorf(format string, args ...any) {
	l.Error(fmt.Errorf(format, args...))
}

func (l Logger) Error(err error) {
	l.logger.Error().Err(err).Msg("")
}
