package event

import (
	"sdl/playing/internal/render/boxes"
	"sync"

	"github.com/rs/zerolog"
	"github.com/veandco/go-sdl2/sdl"
)

var runningMutext sync.Mutex

func HandleEvent(logger zerolog.Logger) bool {
	running := true
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			runningMutext.Lock()
			running = false
			runningMutext.Unlock()
			/*
				case *sdl.MouseMotionEvent:
					logger.Debug.Msgf("[%d ms] MouseMotion type:%d id:%d x:%d y:%d xrel:%d yrel:%d\n",
						t.Timestamp, t.Type, t.Which, t.X, t.Y, t.XRel, t.YRel)
				case *sdl.MouseButtonEvent:
					logger.Debug.Msgf("[%d ms] MouseButton type:%d id:%d x:%d y:%d button:%d state:%d\n",
						t.Timestamp, t.Type, t.Which, t.X, t.Y, t.Button, t.State)
				case *sdl.MouseWheelEvent:
					logger.Debug.Msgf("[%d ms] MouseWheel type:%d id:%d x:%d y:%d\n",
						t.Timestamp, t.Type, t.Which, t.X, t.Y)
			*/
		case *sdl.KeyboardEvent:
			logger.Debug().Msgf("[%d ms] Keyboard type: \"%d\" sym: \"%c\" keycode: \"%v\" modifiers: \"%d\" state: \"%d\" repeat: \"%d\"",
				t.Timestamp, t.Type, t.Keysym.Sym, t.Keysym.Scancode, t.Keysym.Mod, t.State, t.Repeat)
			if t.Type == sdl.KEYDOWN {
				switch t.Keysym.Sym {
				case sdl.K_LEFT:
					boxes.MoveLeft()
				case sdl.K_q:
					boxes.MoveLeft()
				case sdl.K_DOWN:
					boxes.MoveBottom()
				case sdl.K_s:
					boxes.MoveBottom()
				case sdl.K_RIGHT:
					boxes.MoveRight()
				case sdl.K_d:
					boxes.MoveRight()
				case sdl.K_UP:
					boxes.MoveTop()
				case sdl.K_z:
					boxes.MoveTop()
				case sdl.K_ESCAPE:
					running = false
				}
			}
		case *sdl.UserEvent:
			logger.Debug().Msgf("[%d ms] UserEvent code:%d\n", t.Timestamp, t.Code)
		}
	}

	return running
}
