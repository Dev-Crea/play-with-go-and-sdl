package assets

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

func Orange() sdl.Color {
	return sdl.Color{R: Red().R, G: Green().G, B: Green().B, A: 255}
}

func RGBAOrange() (r, g, b, a uint8) {
	return Red().R, Green().G, Green().B, 255
}

func RGBARed() (r, g, b, a uint8) {
	return Red().R, Red().G, Red().B, 255
}

func RGBABlack() (r, g, b, a uint8) {
	return 0, 0, 0, 255
}

func RGBABrown() (r, g, b, a uint8) {
	return Brown().R, Brown().G, Brown().B, 255
}
