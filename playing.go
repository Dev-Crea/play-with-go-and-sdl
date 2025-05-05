package main

import (
	"fmt"
	"os"
	"sync"

	"sdl/playing/boxes"
	"sdl/playing/colors"
	"sdl/playing/constants"
	event "sdl/playing/events"
	"sdl/playing/space-traders/agents"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	RectWidth  = 20
	RectHeight = 20
	NumRects   = constants.WINDOW_WIDTH / constants.WINDOW_HEIGHT
)

func main() {
	var exitcode int

	initSpaceTradersData()

	sdl.Main(func() {
		exitcode = run()
	})

	os.Exit(exitcode)
}

func initSpaceTradersData() {
	agents.Init()
	/*
		wg := sync.WaitGroup{}
		wg.Add(1)
		go func() {
			agents.Init()
			// systems.Init()
			wg.Done()
		}()
		wg.Wait()
	*/
}

func run() int {
	var window *sdl.Window
	var renderer *sdl.Renderer
	var err error

	sdl.Do(func() {
		window, err = sdl.CreateWindow(constants.WINDOW_TITLE,
			sdl.WINDOWPOS_UNDEFINED,
			sdl.WINDOWPOS_UNDEFINED,
			constants.WINDOW_WIDTH,
			constants.WINDOW_HEIGHT,
			sdl.WINDOW_SHOWN)
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		return 1
	}

	defer func() {
		sdl.Do(func() {
			err := window.Destroy()
			if err != nil {
				panic(err)
			}
		})
	}()

	sdl.Do(func() {
		renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		return 2
	}

	defer func() {
		sdl.Do(func() {
			err := renderer.Destroy()
			if err != nil {
				panic(err)
			}
		})
	}()

	sdl.Do(func() {
		err := renderer.Clear()
		if err != nil {
			panic(err)
		}
	})

	running := true

	for running {
		sdl.Do(func() {
			running = event.HandleEvent()

			err := renderer.Clear()
			if err != nil {
				panic(err)
			}

			err = renderer.SetDrawColor(colors.RGBABlack())
			if err != nil {
				panic(err)
			}

			err = renderer.FillRect(&sdl.Rect{X: 0, Y: 0, W: constants.WINDOW_WIDTH, H: constants.WINDOW_HEIGHT})
			if err != nil {
				panic(err)
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
			sdl.Delay(1000 / constants.FRAMERATE)
		})
	}

	return 0
}
