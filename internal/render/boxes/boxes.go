package boxes

import "github.com/rs/zerolog"

var LoggerSDL *zerolog.Logger

func Init(logger *zerolog.Logger) {
	LoggerSDL = logger
}
