package main

import (
	"net/http"
	_ "net/http/pprof"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	go http.ListenAndServe(":8080", nil)
	sdl.Init(sdl.INIT_VIDEO)
	win, _ := sdl.CreateWindow("leak test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 640, 480, 0)
	quit := false
	for !quit {
		ev := sdl.PollEvent()
		switch t := ev.(type) {
		case *sdl.QuitEvent:
			quit = true
		case *sdl.DropEvent:
			println(t.File)
		}
	}
	win.Destroy()
	sdl.Quit()
}
