package events

import (
	"fmt"
	"time"
	"unicode"

	"github.com/macroblock/zl/core/zlog"
	"github.com/veandco/go-sdl2/ttf"
)

var log = zlog.Instance("events")

const _mE = "events "

// IScreen -
type IScreen interface {
	Close()
	AddChild(children ...interface{})
	Draw()
	SetDrawColor(r, g, b, a int)
	SetFillColor(r, g, b, a int)
	DrawText(s string, x, y int)
	Font() *ttf.Font
	SetFont(font *ttf.Font)
	Clear()
	FillRect(x1, y1, w, h int)
	DrawLine(x1, y1, x2, y2 int)
	DrawRect(x1, y1, w, h int)
	Flush()
	PostUpdate()
	// GetClipRect() *TRect
}

type (
	// IEvent -
	IEvent interface {
		Time() time.Time
		Type() string
		EventKey() string
		Screen() IScreen
		String() string
	}

	// TEvent -
	TEvent struct {
		IEvent
		time   time.Time
		screen IScreen
	}
)

// NewEvent -
func NewEvent(scr IScreen) *TEvent {
	ret := &TEvent{time: time.Now()}
	ret.IEvent = ret
	ret.screen = scr
	return ret
}

// Time -
func (o *TEvent) Time() time.Time {
	return o.time
}

// Type -
func (o *TEvent) Type() string {
	return "unknown"
}

// EventKey -
func (o *TEvent) EventKey() string {
	return ""
}

// Screen -
func (o *TEvent) Screen() IScreen {
	return o.screen
}

// String -
func (o *TEvent) String() string {
	return o.time.Format("15:04:05.000") + " " + o.IEvent.Type() + " [" + o.IEvent.EventKey() + "] "
}

// // TActionKeyboardMap -
// type TActionKeyboardMap struct {
// 	name map[string]IAction
// 	mode string
// }
// var ActionKMap TActionKeyboardMap

// func (o *TActionKeyboardMap) initActionKeyboardMap(){
// 	ActionKMap=TActionKeyboardMap{}
// 	ActionKMap.name=map[string]IAction
// }

// func (o *TActionKeyboardMap) Add(event IAction){
// 	o.name[event.]
// }

// TKeyboardEvent -
type TKeyboardEvent struct {
	TWindowEvent
	ch   rune
	scan int
	mod  int
}

// NewKeyboardEvent -
func NewKeyboardEvent(scr IScreen, id int, ch rune, mod int) *TKeyboardEvent {
	ret := &TKeyboardEvent{TWindowEvent: *NewWindowEvent(scr, id), ch: ch, mod: mod}
	ret.IEvent = ret
	return ret
}

// Type -
func (o *TKeyboardEvent) Type() string {
	return "keyboard"
}

// Rune -
func (o *TKeyboardEvent) Rune() rune {
	return o.ch
}

// ScanCode -
func (o *TKeyboardEvent) ScanCode() int {
	return o.scan
}

// EventKey -
func (o *TKeyboardEvent) EventKey() string {
	return fmt.Sprintf("%s", string(o.ch))
}

// String -
func (o *TKeyboardEvent) String() string {
	format := "%v U+%x %v"
	if unicode.IsPrint(o.ch) {
		format = "%v %q %v"
	}
	return fmt.Sprintf(format, o.TWindowEvent.String(), o.ch, o.mod)
}

// TDropFileEvent -
type TDropFileEvent struct {
	TWindowEvent
	content string
}

// NewDropFileEvent -
func NewDropFileEvent(scr IScreen, id int, s string) *TDropFileEvent {
	ret := &TDropFileEvent{TWindowEvent: *NewWindowEvent(scr, id), content: s}
	ret.IEvent = ret
	return ret
}

// Type -
func (o *TDropFileEvent) Type() string {
	return "drop file"
}

// Content -
func (o *TDropFileEvent) Content() string {
	return o.content
}

// String -
func (o *TDropFileEvent) String() string {
	return fmt.Sprintf("%v %v", o.TWindowEvent.String(), o.content)
}

// TWindowEvent -
type TWindowEvent struct {
	TEvent
	windowID int
}

// NewWindowEvent -
func NewWindowEvent(scr IScreen, id int) *TWindowEvent {
	ret := &TWindowEvent{TEvent: *NewEvent(scr), windowID: id}
	ret.IEvent = ret
	return ret
}

// WindowID -
func (o *TWindowEvent) WindowID() int {
	return o.windowID
}

// Type -
func (o *TWindowEvent) Type() string {
	return "window unknown"
}

// String -
func (o *TWindowEvent) String() string {
	return fmt.Sprintf("%v WinID:%v", o.TEvent.String(), o.windowID)
}

// EventKey -
func (o *TWindowEvent) EventKey() string {
	return fmt.Sprintf("window unknown %v", o.windowID)
}

// TWindowCloseEvent -
type TWindowCloseEvent struct {
	TWindowEvent
}

// NewWindowCloseEvent -
func NewWindowCloseEvent(scr IScreen, id int) *TWindowCloseEvent {
	ret := &TWindowCloseEvent{TWindowEvent: *NewWindowEvent(scr, id)}
	ret.IEvent = ret
	return ret
}

// Type -
func (o *TWindowCloseEvent) Type() string {
	return "window close"
}

// String -
func (o *TWindowCloseEvent) String() string {
	return fmt.Sprintf("%v", o.TWindowEvent.String())
}

// EventKey -
func (o *TWindowCloseEvent) EventKey() string {
	return fmt.Sprintf("window close %v", o.windowID)
}

// TWindowResizedEvent -
type TWindowResizedEvent struct {
	TWindowEvent
	windowID int
	w, h     int
}

// NewWindowResizedEvent -
func NewWindowResizedEvent(scr IScreen, id, w, h int) *TWindowResizedEvent {
	ret := &TWindowResizedEvent{TWindowEvent: *NewWindowEvent(scr, id), w: w, h: h}
	ret.IEvent = ret
	return ret
}

// Size -
func (o *TWindowResizedEvent) Size() (int, int) {
	return o.w, o.h
}

// Type -
func (o *TWindowResizedEvent) Type() string {
	return "window resized"
}

// String -
func (o *TWindowResizedEvent) String() string {
	return fmt.Sprintf("%v %vx%v", o.TWindowEvent.String(), o.w, o.h)
}

// EventKey -
func (o *TWindowResizedEvent) EventKey() string {
	return fmt.Sprintf("window resized")

	// return fmt.Sprintf("window resized %v", o.windowID)
}
