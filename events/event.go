package event

import (
	"fmt"
	"sync"

	"sdl/playing/boxes"

	"github.com/veandco/go-sdl2/sdl"
)

var runningMutext sync.Mutex

func HandleEvent() bool {
	running := true
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			runningMutext.Lock()
			running = false
			runningMutext.Unlock()
			/*
				case *sdl.MouseMotionEvent:
					fmt.Printf("[%d ms] MouseMotion\ttype:%d\tid:%d\tx:%d\ty:%d\txrel:%d\tyrel:%d\n",
						t.Timestamp, t.Type, t.Which, t.X, t.Y, t.XRel, t.YRel)
				case *sdl.MouseButtonEvent:
					fmt.Printf("[%d ms] MouseButton\ttype:%d\tid:%d\tx:%d\ty:%d\tbutton:%d\tstate:%d\n",
						t.Timestamp, t.Type, t.Which, t.X, t.Y, t.Button, t.State)
				case *sdl.MouseWheelEvent:
					fmt.Printf("[%d ms] MouseWheel\ttype:%d\tid:%d\tx:%d\ty:%d\n",
						t.Timestamp, t.Type, t.Which, t.X, t.Y)
			*/
		case *sdl.KeyboardEvent:
			fmt.Printf("[%d ms] Keyboard\ttype:%d\tsym:%c\tkeycode:%v\tmodifiers:%d\tstate:%d\trepeat:%d\n",
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
			fmt.Printf("[%d ms] UserEvent\tcode:%d\n", t.Timestamp, t.Code)
		}
	}

	return running
}
