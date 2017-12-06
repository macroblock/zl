package hal

import (
	"github.com/macroblock/zl/core/zlog"
	"github.com/macroblock/zl/ui/events"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// import "C"

var log = zlog.Instance("zl/sdl")

var hal *THal

var currentScreen = stubScreen

// Screen -
func Screen() IScreen       { return currentScreen }
func makeCurrent(o IScreen) { currentScreen = o }

// THal -
type THal struct {
	screen map[uint32]*TScreen
}

// New -
func New() (*THal, error) {
	if hal != nil {
		log.Warning(true, "New: HAL was already intialized")
		return nil, nil
	}

	hal = &THal{}
	hal.screen = make(map[uint32]*TScreen)

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
	for key := range o.screen {
		delete(o.screen, key)
	}
	o.screen = nil
	hal = nil
	ttf.Quit()
	sdl.Quit()
}

// NewScreen -
func (o *THal) NewScreen() (IScreen, error) {
	win, r, err := sdl.CreateWindowAndRenderer(640, 480, sdl.WINDOW_SHOWN|sdl.WINDOW_RESIZABLE)
	log.Error(err, "NewOutput: sdl.CreateWindowAndRenderer")
	scr := &TScreen{hal: o, window: win, renderer: r, font: defaultFont}
	//o.outputs.PushBack(scr)
	id, err := win.GetID()
	log.Error(err, "NewOutput: Window.GetID")
	o.screen[id] = scr
	makeCurrent(scr)
	scr.PostUpdate()
	log.Debug("Create window id: ", id)
	return scr, err
}

// Screen -
func (o *THal) Screen(id int) IScreen {
	scr := IScreen(nil)
	ok := false
	scr, ok = o.screen[uint32(id)]
	log.Warning(!ok, "screen(): screen not found - id: ", id)
	if !ok || scr == nil {
		scr = stubScreen
	}
	return scr
}

// Event -
func (o *THal) Event() events.IEvent {
	e := sdl.PollEvent()
	if e == nil {
		return nil
	}
	switch t := e.(type) {
	case *sdl.KeyboardEvent:
		event := events.NewKeyboardEvent(o.Screen(int(t.WindowID)), int(t.WindowID), rune(t.Keysym.Sym), int(t.Keysym.Mod))
		return event
	case *sdl.DropEvent:
		if t.Type == sdl.DROPFILE {
			event := events.NewDropFileEvent(o.Screen(int(t.WindowID)), int(t.WindowID), t.File)
			return event
		}
	case *sdl.WindowEvent:
		switch t.Event {
		case sdl.WINDOWEVENT_CLOSE:
			event := events.NewWindowCloseEvent(o.Screen(int(t.WindowID)), int(t.WindowID))
			return event
		case sdl.WINDOWEVENT_RESIZED:
			event := events.NewWindowResizedEvent(o.Screen(int(t.WindowID)), int(t.WindowID), int(t.Data1), int(t.Data2))
			o.Screen(int(t.WindowID)).PostUpdate()
			return event
		}
	case *sdl.CommonEvent:
		log.Warning(true, "default event: ", t.Type)
	} // end of switch
	return nil
}

// Draw -
func (o *THal) Draw() {
	for id, output := range o.screen {
		log.Warning(output == nil, "output id: ", id, " is nil")
		if output == nil || !output.NeedUpdate() {
			continue
		}
		makeCurrent(output)
		output.Draw()
		output.Flush()
		output.ResetUpdate()

	}
	makeCurrent(stubScreen)
}
