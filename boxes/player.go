package boxes

import (
	"sdl/playing/colors"
	"sdl/playing/constants"

	"github.com/veandco/go-sdl2/sdl"
)

var (
	playerPositionX int32 = constants.PANEL_GAME_CENTER_X
	playerPositionY int32 = constants.PANEL_GAME_CENTER_Y
)

func Player(renderer *sdl.Renderer) {
	err := renderer.SetDrawColor(colors.RGBARed())
	if err != nil {
		panic(err)
	}

	box := sdl.Rect{X: playerPositionX, Y: playerPositionY, W: 10, H: 10}
	err = renderer.FillRect(&box)
	if err != nil {
		panic(err)
	}
}

func moveX(ux int32, x int32) int32 {
	newPosition := playerPositionX + x - ux

	if newPosition == 0 || newPosition == constants.WINDOW_WIDTH-10 {
		return playerPositionX
	}

	return newPosition
}

func moveY(uy int32, y int32) int32 {
	newPosition := playerPositionY + y - uy

	if newPosition == 0 || newPosition == constants.WINDOW_HEIGHT-10 {
		return playerPositionY
	}

	return newPosition
}

func move(ux int32, x int32, uy int32, y int32) {
	playerPositionX = moveX(ux, x)
	playerPositionY = moveY(uy, y)
}

func MoveLeft() {
	move(constants.MOVE_PLAYER_SIZE, 0, 0, 0)
}

func MoveRight() {
	move(0, constants.MOVE_PLAYER_SIZE, 0, 0)
}

func MoveTop() {
	move(0, 0, constants.MOVE_PLAYER_SIZE, 0)
}

func MoveBottom() {
	move(0, 0, 0, constants.MOVE_PLAYER_SIZE)
}
