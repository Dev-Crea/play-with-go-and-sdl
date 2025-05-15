package playing

import (
	boxes "sdl/playing/internal/render"
	"sdl/playing/internal/space-traders/models"

	"github.com/veandco/go-sdl2/sdl"
)

func InitSpaceTradersData() {
	models.Init()
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

func InitSDL() {
	sdl.Main(func() {
		boxes.Run()
	})
}
