package render

import (
	"os"
	"sdl/playing/internal/render/assets"
	"sdl/playing/internal/render/boxes"
	"sync"

	event "sdl/playing/internal/render/events"
	tools "sdl/playing/pkg/Tools"

	"github.com/rs/zerolog"
	"github.com/veandco/go-sdl2/sdl"
)

func Run() int {
	var window *sdl.Window
	var renderer *sdl.Renderer
	var err error

	loggerSDL := tools.InitLoggerByService("SDL", zerolog.DebugLevel)

	boxes.Init(loggerSDL)

	sdl.Do(func() {
		window, err = sdl.CreateWindow(assets.WINDOW_TITLE,
			sdl.WINDOWPOS_UNDEFINED,
			sdl.WINDOWPOS_UNDEFINED,
			assets.WINDOW_WIDTH,
			assets.WINDOW_HEIGHT,
			sdl.WINDOW_SHOWN)
	})
	if err != nil {
		loggerSDL.Error().Stack().Err(err).Msg("Failed to create window")

		os.Exit(1)
	}

	defer func() {
		sdl.Do(func() {
			err := window.Destroy()
			if err != nil {
				loggerSDL.Error().Stack().Err(err).Msg("Failed destroy window")
			}
		})
	}()

	sdl.Do(func() {
		renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	})
	if err != nil {
		loggerSDL.Error().Stack().Err(err).Msg("Failed to create renderer")
	}

	defer func() {
		sdl.Do(func() {
			err := renderer.Destroy()
			if err != nil {
				loggerSDL.Error().Stack().Err(err).Msg("Failed destroy renderer")

				os.Exit(1)
			}
		})
	}()

	sdl.Do(func() {
		err := renderer.Clear()
		if err != nil {
			loggerSDL.Error().Stack().Err(err).Msg("Failed clear renderer")

			os.Exit(1)
		}
	})

	running := true

	for running {
		sdl.Do(func() {
			running = event.HandleEvent(*loggerSDL)

			err := renderer.Clear()
			if err != nil {
				loggerSDL.Error().Stack().Err(err).Msg("Failed clear renderer")

				os.Exit(1)
			}

			err = renderer.SetDrawColor(assets.RGBABlack())
			if err != nil {
				loggerSDL.Error().Stack().Err(err).Msg("Failed draw backgound renderer")

				os.Exit(1)
			}

			err = renderer.FillRect(&sdl.Rect{X: 0, Y: 0, W: assets.WINDOW_WIDTH, H: assets.WINDOW_HEIGHT})
			if err != nil {
				loggerSDL.Error().Stack().Err(err).Msg("Failed fill backgound renderer")

				os.Exit(1)

			}
		})

		wg := sync.WaitGroup{}
		wg.Add(1)
		go func() {
			sdl.Do(func() {
				boxes.PanelGame(renderer)
				boxes.Player(renderer)
				boxes.Orbitals(renderer)
			})
			wg.Done()
		}()
		wg.Wait()

		sdl.Do(func() {
			renderer.Present()
			sdl.Delay(1000 / assets.FRAMERATE)
		})
	}

	return 0
}
