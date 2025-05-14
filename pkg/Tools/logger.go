package tools

import (
	"os"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

/*
		NOTE: Level to logger

	    panic (zerolog.PanicLevel, 5)
	    fatal (zerolog.FatalLevel, 4)
	    error (zerolog.ErrorLevel, 3)
	    warn (zerolog.WarnLevel, 2)
	    info (zerolog.InfoLevel, 1)
	    debug (zerolog.DebugLevel, 0)
	    trace (zerolog.TraceLevel, -1)
*/
func InitLoggerByService(name string, level zerolog.Level) *zerolog.Logger {
	var logger zerolog.Logger

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = func(err error) interface{} {
		return pkgerrors.MarshalStack(errors.WithStack(err))
	}
	zerolog.SetGlobalLevel(level)

	if level == zerolog.InfoLevel {
		logger = loggerInfo(name)
	}
	switch level {
	default:
		logger = loggerDebug(name)
	case zerolog.DebugLevel:
		logger = loggerDebug(name)
	case zerolog.InfoLevel:
		logger = loggerInfo(name)
	case zerolog.ErrorLevel:
		logger = loggerError(name)
	}

	logger.Debug().Msg("Start application in DEBUG mode !")

	return &logger
}

func loggerDebug(name string) zerolog.Logger {
	return zerolog.New(zerolog.ConsoleWriter{
		Out:        os.Stdout,
		NoColor:    true,
		TimeFormat: time.RFC3339,
	}).
		With().
		Timestamp().
		Str("Service", name).
		Logger()
}

func loggerInfo(name string) zerolog.Logger {
	return zerolog.New(os.Stdout).
		With().
		Timestamp().
		Str("Service", name).
		Logger()
}

func loggerError(name string) zerolog.Logger {
	file, err := os.OpenFile(
		"playing.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0o600,
	)
	if err != nil {
		panic(err)
	}

	return zerolog.New(zerolog.ConsoleWriter{
		Out:        file,
		NoColor:    true,
		TimeFormat: time.RFC3339,
	}).
		With().
		Timestamp().
		Str("Service", name).
		Logger()
}
