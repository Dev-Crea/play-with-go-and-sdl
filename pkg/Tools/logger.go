package tools

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

func InitLoggerByService(name string, level zerolog.Level) *zerolog.Logger {
	var logger zerolog.Logger

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = func(err error) interface{} {
		return pkgerrors.MarshalStack(errors.WithStack(err))
	}
	zerolog.SetGlobalLevel(level)

	switch level {
	case zerolog.DebugLevel:
		logger = loggerDebug(name)
	case zerolog.InfoLevel:
		logger = loggerInfo(name)
	case zerolog.ErrorLevel:
		logger = loggerError(name)
	default:
		logger = loggerDebug(name)
	}

	logger.Info().Msgf("Start application in %s mode with level %s !", name, level)

	return &logger
}

func loggerOutput() zerolog.ConsoleWriter {
	return zerolog.ConsoleWriter{
		Out:           os.Stdout,
		NoColor:       false,
		TimeFormat:    time.RFC3339,
		PartsOrder:    []string{"level", "Service", "time", "message"},
		FieldsExclude: []string{"Service"},
		FormatLevel:   func(i interface{}) string { return strings.ToUpper(fmt.Sprintf("%-6s", i)) },
	}
}

func loggerDebug(name string) zerolog.Logger {
	return zerolog.New(loggerOutput()).
		With().
		Timestamp().
		Str("Service", name).
		Logger()
}

func loggerInfo(name string) zerolog.Logger {
	return zerolog.New(loggerOutput()).
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
