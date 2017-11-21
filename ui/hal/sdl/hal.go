package hal

import (
	"github.com/macroblock/zl/core/zlog"
	"github.com/macroblock/zl/types/vector"
	"github.com/macroblock/zl/ui/events"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

import "C"

var log = zlog.Instance("zl/sdl")

var hal = (*THal)(nil)

// THal -
type THal struct {
	outputs vector.TVector
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

	return hal, nil
}

// Close -
func (o *THal) Close() {
	if hal == nil {
		log.Warning(true, "THal.Close: HAL isn't initialized yet")
		return
	}
	for o.outputs.Len() > 0 {
		v, err := o.outputs.Back()
		log.Error(err, "THal.Close: something wrong")
		output, ok := v.(*TOutput)
		log.Error(!ok, "THal.Close: object is not TOutput")
		output.Close()
	}
	hal = nil
	ttf.Quit()
	sdl.Quit()
}

// NewOutput -
func (o *THal) NewOutput() (*TOutput, error) {
	win, r, err := sdl.CreateWindowAndRenderer(640, 480, sdl.WINDOW_SHOWN|sdl.WINDOW_RESIZABLE)
	log.Error(err, "New: sdl.CreateWindowAndRenderer")
	ret := &TOutput{hal: o, window: win, renderer: r}
	o.outputs.PushBack(ret)
	return ret, err
}

// Event -
func (o *THal) Event() events.IEvent {
	e := sdl.PollEvent()
	if e == nil {
		return nil
	}
	switch t := e.(type) {
	case *sdl.KeyboardEvent:
		event := events.NewKeyboardEvent(rune(t.Keysym.Sym), int(t.Keysym.Mod))
		return event
	case *sdl.DropEvent:
		if t.Type == sdl.DROPFILE {
			event := events.NewDropFileEvent(t.File) //C.GoString((*C.char)(t.File)))
			return event
		}
	} // end of switch
	return nil
}
