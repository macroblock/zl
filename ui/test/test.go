package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

func drawgl(w, h int) {
	println("21323")
}
func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	rect := sdl.Rect{0, 0, 200, 200}
	surface.FillRect(&rect, 0xffff0000)
	window.UpdateSurface()

	running := true

	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.MouseMotionEvent:
				// fmt.Printf("[%d ms] MouseMotion %v\n", t.Timestamp, t.Which)
			case *sdl.KeyDownEvent:
				// fmt.Printf("[%d ms] Key %v \n", t.Timestamp, t.Keysym)
				if t.Keysym.Sym == 27 {
					running = false
				}
			}
			drawgl(0, 0)
		}

		sdl.GL_SwapWindow(window)
	}
}
