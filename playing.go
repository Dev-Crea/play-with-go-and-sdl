package main

import (
	"fmt"
	"math/rand"
	"os"

	"sdl/playing/constants"
	event "sdl/playing/events"
	"sdl/playing/scenes/panels"
	"sdl/playing/space-traders/agents"
	"sdl/playing/space-traders/systems"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	RectWidth  = 20
	RectHeight = 20
	NumRects   = constants.WINDOW_WIDTH / constants.WINDOW_HEIGHT
)

var rects [NumRects]sdl.Rect

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
	systems.Init()
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
			window.Destroy()
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
			renderer.Destroy()
		})
	}()

	sdl.Do(func() {
		renderer.Clear()
	})

	for i := range rects {
		rects[i] = sdl.Rect{
			X: int32(rand.Int() % constants.WINDOW_WIDTH),
			Y: int32(i * constants.WINDOW_HEIGHT / len(rects)),
			W: RectWidth,
			H: RectHeight,
		}
	}

	running := true

	for running {
		sdl.Do(func() {
			running = event.HandleEvent()

			err := renderer.Clear()
			if err != nil {
				panic(err)
			}

			err = renderer.SetDrawColor(0, 0, 0, 0x20)
			if err != nil {
				panic(err)
			}

			err = renderer.FillRect(&sdl.Rect{X: 0, Y: 0, W: constants.WINDOW_WIDTH, H: constants.WINDOW_HEIGHT})
			if err != nil {
				panic(err)
			}
		})

		// updateSurface(renderer)
		panels.PanelGame(renderer)
		/*
			wg := sync.WaitGroup{}
			for i := range rects {
				wg.Add(1)
				go func(i int) {
					rects[i].X = (rects[i].X + 10) % constants.WINDOW_WIDTH
					sdl.Do(func() {
						renderer.SetDrawColor(0xff, 0xff, 0xff, 0xff)
						renderer.DrawRect(&rects[i])
					})
					wg.Done()
				}(i)
			}

			wg.Wait()
		*/

		sdl.Do(func() {
			renderer.Present()
			sdl.Delay(1000 / constants.FRAMERATE)
		})
	}
	/*
		UpdateSurface(surface)

		running := true
		for running {
			running = event.HandleEvent(surface)
			UpdateSurface(surface)
			err := window.UpdateSurface()
			if err != nil {
				panic(err)
			}
		}
	*/

	return 0
}

/*
func updateSurface(renderer sdl.Renderer) {
		panels.PanelGame(*surface)
		agents.GetCurrentOrbitals(*surface)
		player.Init(*surface)
}
*/
