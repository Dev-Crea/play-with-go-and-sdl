package boxes

import (
	"sdl/playing/assets"
	"sdl/playing/space-traders/agents"

	"github.com/veandco/go-sdl2/sdl"
)

func orbital(renderer *sdl.Renderer, position int32) {
	err := renderer.SetDrawColor(assets.RGBABrown())
	if err != nil {
		panic(err)
	}

	positionX := assets.PANEL_GAME_CENTER_X - position
	box := sdl.Rect{X: positionX, Y: assets.PANEL_GAME_CENTER_Y, W: 10, H: 10}
	err = renderer.FillRect(&box)
	if err != nil {
		panic(err)
	}
}

func Orbitals(renderer *sdl.Renderer) {
	for key := range agents.MyPosition.Data.Orbitals {
		orbital(renderer, int32((key+1)*50))
	}
}
