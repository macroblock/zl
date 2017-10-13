package hal

import (
	"github.com/macroblock/zl/core/zlog"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

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
	log.Error(error(err), "New: sdl.Init")

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
		log.Warning(true, "Close: HAL was not initialized yet")
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

// IEvent -
type IEvent interface {
	String() string
	Key() string
}

// Event -
func (o *THal) Event() IEvent {
	return nil
}
