package tools

import (
	"os"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

func InitLoggerByService(name string, level zerolog.Level) *zerolog.Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = func(err error) interface{} {
		return pkgerrors.MarshalStack(errors.WithStack(err))
	}
	zerolog.SetGlobalLevel(level)

	logger := zerolog.New(os.Stdout).
		With().
		Timestamp().
		Str("Service", name).
		Logger()

	logger.Debug().Msg("Start application in DEBUG mode !")

	return &logger
}
