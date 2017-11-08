package hal

import (
	"github.com/macroblock/zl/core/zlog"
	"github.com/macroblock/zl/ui/events"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

import "C"

var log = zlog.Instance("zl/sdl")

var hal = (*THal)(nil)

// THal -
type THal struct {
	window   *sdl.Window
	renderer *sdl.Renderer
}

// New -
func New() (*THal, error) {
	if hal != nil {
		log.Warning(true, "New: HAL was already intialized")
		return nil, nil
	}

	hal = &THal{}

	err := sdl.Init(sdl.INIT_EVERYTHING)
	log.Error(err, "New: sdl.Init")

	err = ttf.Init()
	log.Error(err, "New: ttf.Init")

	InitFonts()

	hal.window, hal.renderer, err = sdl.CreateWindowAndRenderer(640, 480, sdl.WINDOW_SHOWN|sdl.WINDOW_RESIZABLE)
	log.Error(err, "New: sdl.CreateWindowAndRenderer")

	return hal, nil
}

// Close -
func (o *THal) Close() {
	if hal == nil {
		log.Warning(true, "Close: HAL isn't initialized yet")
		return
	}
	//for _, root := range glob.sysWindows {
	//	root.Close()
	//}
	//glob.rootWindows = nil

	if hal.renderer != nil {
		hal.renderer.Destroy()
		hal.renderer = nil
	}
	if hal.window != nil {
		hal.window.Destroy()
		hal.window = nil
	}

	hal = nil
	ttf.Quit()
	sdl.Quit()
}

// Event -
func (o *THal) Event() events.IEvent {
	e := sdl.PollEvent()
	if e == nil {
		return nil
	}
	switch t := e.(type) {
	case *sdl.KeyDownEvent:
		event := events.NewKeyboardEvent(rune(t.Keysym.Sym), int(t.Keysym.Mod))
		return event
	case *sdl.DropEvent:
		if t.Type == sdl.DROPFILE {
			event := events.NewDropFileEvent(t.File) //C.GoString((*C.char)(t.File)))
			return event
		}
	}
	return nil
}
