package system

import (
	"sdl/playing/colors"
	"sdl/playing/constants"

	"github.com/veandco/go-sdl2/sdl"
)

func System(surface sdl.Surface, position int32) {
	positionX := constants.PANEL_GAME_CENTER_X - position
	box := sdl.Rect{X: positionX, Y: constants.PANEL_GAME_CENTER_Y, W: 10, H: 10}

	color := colors.Brown()
	pixel := sdl.MapRGBA(surface.Format, color.R, color.G, color.B, color.A)

	err := surface.FillRect(&box, pixel)
	if err != nil {
		panic(err)
	}
}
