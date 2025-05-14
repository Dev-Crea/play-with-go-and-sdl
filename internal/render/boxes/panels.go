package boxes

import (
	"sdl/playing/internal/render/assets"
	"sdl/playing/internal/space-traders/models"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func PanelGame(renderer *sdl.Renderer) {
	// Draw rect background panel game (orange)
	err := renderer.SetDrawColor(assets.RGBAOrange())
	if err != nil {
		LoggerSDL.Error().Stack().Err(err).Msg("")
	}

	boxBackground := sdl.Rect{X: 0, Y: assets.WINDOW_HEIGHT - assets.PANEL_HEIGHT, W: assets.WINDOW_WIDTH, H: assets.PANEL_HEIGHT}

	err = renderer.FillRect(&boxBackground)
	if err != nil {
		LoggerSDL.Error().Stack().Err(err).Msg("")
	}

	// Writte info about this game
	myAgent := models.MyAgent

	window, err := renderer.GetWindow()
	if err != nil {
		LoggerSDL.Error().Stack().Err(err).Msg("")
	}

	surface, err := window.GetSurface()
	if err != nil {
		LoggerSDL.Error().Stack().Err(err).Msg("")
	}
	defer surface.Free()

	err = ttf.Init()
	if err != nil {
		LoggerSDL.Error().Stack().Err(err).Msg("")
	}
	defer ttf.Quit()

	font, err := ttf.OpenFont(assets.FONT_PATH, assets.FONT_PANEL_SIZE)
	if err != nil {
		LoggerSDL.Error().Stack().Err(err).Msg("")
	}
	defer font.Close()

	writeInPanel(*font, myAgent.GetSymbol(), surface, renderer, 10, 510)
	writeInPanel(*font, myAgent.GetHeadquarter(), surface, renderer, 10, 539)
	writeInPanel(*font, myAgent.GetCredits(), surface, renderer, 10, 568)
	writeInPanel(*font, myAgent.GetFaction(), surface, renderer, 500, 510)
	writeInPanel(*font, myAgent.GetFleet(), surface, renderer, 500, 560)
}

func writeInPanel(font ttf.Font, word string, surface *sdl.Surface, renderer *sdl.Renderer, x, y int32) {
	text, err := font.RenderUTF8Solid(word, assets.Green())
	if err != nil {
		LoggerSDL.Error().Stack().Err(err).Msg("")
	}

	defer text.Free()

	boxTextDest := sdl.Rect{X: x, Y: y, W: 150, H: assets.FONT_PANEL_SIZE}

	err = text.Blit(nil, surface, &boxTextDest)
	if err != nil {
		LoggerSDL.Error().Stack().Err(err).Msg("")
	}

	texture, err := renderer.CreateTextureFromSurface(surface)
	if err != nil {
		LoggerSDL.Error().Stack().Err(err).Msg("")
	}
	defer func() {
		err := texture.Destroy()
		if err != nil {
			LoggerSDL.Error().Stack().Err(err).Msg("")
		}
	}()

	err = renderer.Copy(texture, &boxTextDest, &boxTextDest)
	if err != nil {
		LoggerSDL.Error().Stack().Err(err).Msg("")
	}
}
