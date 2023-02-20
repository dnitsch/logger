package log

import (
	"fmt"
	"io"

	"github.com/go-logr/logr"
	"github.com/go-logr/zerologr"
	"github.com/rs/zerolog"
)

func NewLogr(writer io.Writer, lvl LogLevel) logr.Logger {
	zloglvl := zerolog.ErrorLevel

	zloglvl, err := zerolog.ParseLevel(string(lvl))
	if err != nil {
		fmt.Printf("loggerFailed: %v. setting error level", err)
		zloglvl = zerolog.ErrorLevel
	}

	zerolog.CallerMarshalFunc = lShortMarshall()
	zl := zerolog.New(writer).With().Timestamp().CallerWithSkipFrameCount(zerolog.CallerSkipFrameCount + 1).Logger().Level(zloglvl)
	return zerologr.New(&zl)
}
