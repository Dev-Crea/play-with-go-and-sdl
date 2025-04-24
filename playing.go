package main

import (
	"sdl/playing/constants"
	event "sdl/playing/events"
	"sdl/playing/scenes/panels"
	player "sdl/playing/scenes/players"
	"sdl/playing/space-traders/agents"
	"sdl/playing/space-traders/systems"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	agents.Init()
	systems.Init()

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow(constants.WINDOW_TITLE,
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		constants.WINDOW_WIDTH,
		constants.WINDOW_HEIGHT,
		sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}

	defer window.Destroy() //nolint

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

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
}

func UpdateSurface(surface *sdl.Surface) {
	err := surface.FillRect(nil, 0)
	if err != nil {
		panic(err)
	}

	panels.PanelGame(*surface)
	agents.GetCurrentOrbitals(*surface)
	player.Init(*surface)
}
