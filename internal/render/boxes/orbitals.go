package boxes

import (
	"sdl/playing/internal/render/assets"
	"sdl/playing/internal/space-traders/models"

	"github.com/veandco/go-sdl2/sdl"
)

func orbital(renderer *sdl.Renderer, position int32) {
	err := renderer.SetDrawColor(assets.RGBABrown())
	if err != nil {
		models.LoggerAPI.Error().Err(err).Msg("")
	}

	positionX := assets.PANEL_GAME_CENTER_X - position
	box := sdl.Rect{X: positionX, Y: assets.PANEL_GAME_CENTER_Y, W: 10, H: 10}
	err = renderer.FillRect(&box)
	if err != nil {
		panic(err)
	}
}

func Orbitals(renderer *sdl.Renderer) {
	for key := range models.MyPosition.Data.Orbitals {
		orbital(renderer, int32((key+1)*50))
	}
}
