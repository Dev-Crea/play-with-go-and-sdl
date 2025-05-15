package models

import (
	space_traders "sdl/playing/internal/space-traders"
	tools "sdl/playing/pkg/Tools"

	"github.com/rs/zerolog"
)

var LoggerAPI *zerolog.Logger

func Init() {
	LoggerAPI = tools.InitLoggerByService("API", zerolog.DebugLevel)

	space_traders.InitAPI(*LoggerAPI)

	InitAgent()
}
