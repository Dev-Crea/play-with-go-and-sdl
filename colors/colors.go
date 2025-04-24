package colors

import (
	"github.com/veandco/go-sdl2/sdl"
)

func Purple() sdl.Color {
	return sdl.Color{R: 255, G: 0, B: 255, A: 255}
}

func Green() sdl.Color {
	return sdl.Color{R: 0, G: 158, B: 48, A: 255}
}

func Red() sdl.Color {
	return sdl.Color{R: 255, G: 0, B: 0, A: 255}
}

func Brown() sdl.Color {
	return sdl.Color{R: 153, G: 76, B: 0, A: 255}
}

func Black() sdl.Color {
	return sdl.Color{R: 0, G: 0, B: 0, A: 255}
}

func White() sdl.Color {
	return sdl.Color{R: 255, G: 255, B: 255, A: 255}
}
