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
	outputs map[uint32]*TOutput
}

// New -
func New() (*THal, error) {
	if hal != nil {
		log.Warning(true, "New: HAL was already intialized")
		return nil, nil
	}

	hal = &THal{}
	hal.outputs = make(map[uint32]*TOutput)

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
	// for o.outputs.Len() > 0 {
	// 	v, err := o.outputs.Back()
	// 	log.Error(err, "THal.Close: something wrong")
	// 	output, ok := v.(*TOutput)
	// 	log.Error(!ok, "THal.Close: object is not TOutput")
	// 	output.Close()
	// }
	for key := range o.outputs {
		delete(o.outputs, key)
	}
	o.outputs = nil
	hal = nil
	ttf.Quit()
	sdl.Quit()
}

// NewOutput -
func (o *THal) NewOutput() (*TOutput, error) {
	win, r, err := sdl.CreateWindowAndRenderer(640, 480, sdl.WINDOW_SHOWN|sdl.WINDOW_RESIZABLE)
	log.Error(err, "NewOutput: sdl.CreateWindowAndRenderer")
	ret := &TOutput{hal: o, window: win, renderer: r}
	//o.outputs.PushBack(ret)
	id := win.GetID()
	log.Error(id == 0, "NewOutput: Window.GetID")
	o.outputs[id] = ret

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
			event := events.NewDropFileEvent(t.File)
			return event
		}
	case *sdl.CommonEvent:
		log.Warning(true, "default event: ", t.Type)
	} // end of switch
	return nil
}
