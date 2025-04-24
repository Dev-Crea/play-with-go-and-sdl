package panels

import (
	"sdl/playing/colors"
	"sdl/playing/constants"
	"sdl/playing/space-traders/agents"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func PanelGame(surface sdl.Surface) {
	// Draw rect background panel game
	boxBackground := sdl.Rect{X: 0, Y: constants.WINDOW_HEIGHT - constants.PANEL_HEIGHT, W: constants.WINDOW_WIDTH, H: constants.PANEL_HEIGHT}
	color := colors.Green()

	pixel := sdl.MapRGBA(surface.Format, color.R, color.G, color.B, color.A)

	surface.FillRect(&boxBackground, pixel)

	// Display info about agent in panel game
	if ttf.Init() != nil {
		return
	}
	defer ttf.Quit()

	font, err := ttf.OpenFont(constants.FONT_PATH, constants.FONT_PANEL_SIZE)
	if err != nil {
		panic(err)
	}
	defer font.Close()

	// fmt.Println("Symbol :: %s", agents.GetAgentData().GetSymbol())

	myAgent := agents.MyAgent

	writeInPanelLeft(*font, myAgent.GetSymbol(), &surface, 510)
	writeInPanelLeft(*font, myAgent.GetHeadquarter(), &surface, 535)
	writeInPanelLeft(*font, myAgent.GetCredits(), &surface, 560)

	writeInPanelRight(*font, myAgent.GetFaction(), &surface, 510)
	writeInPanelRight(*font, myAgent.GetFleet(), &surface, 560)
}

func writeInPanelLeft(font ttf.Font, word string, surface *sdl.Surface, positionY int32) {
	text, err := font.RenderUTF8Blended(word, colors.Black())
	if err != nil {
		panic(err)
	}

	defer text.Free()

	boxText := sdl.Rect{X: 10, Y: positionY, W: 0, H: 0}

	text.Blit(nil, surface, &boxText)
}

func writeInPanelRight(font ttf.Font, word string, surface *sdl.Surface, positionY int32) {
	text, err := font.RenderUTF8Blended(word, colors.Black())
	if err != nil {
		panic(err)
	}

	defer text.Free()

	boxText := sdl.Rect{X: 500, Y: positionY, W: 0, H: 0}

	text.Blit(nil, surface, &boxText)
}
