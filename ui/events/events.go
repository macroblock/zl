package events

import (
	"fmt"
	"time"
	"unicode"

	"github.com/macroblock/zl/core/zlog"
	"github.com/macroblock/zl/ui/hal/interfaces"
)

var log = zlog.Instance("events")

const _mE = "events "

// // interfaces.IScreen -
// type IScreen interface {
// 	Close()
// 	AddChild(children ...interface{})
// 	Draw()
// 	SetDrawColor(r, g, b, a int)
// 	SetFillColor(r, g, b, a int)
// 	DrawText(s string, x, y int)
// 	Font() *ttf.Font
// 	SetFont(font *ttf.Font)
// 	Clear()
// 	FillRect(x1, y1, w, h int)
// 	DrawLine(x1, y1, x2, y2 int)
// 	DrawRect(x1, y1, w, h int)
// 	Flush()
// 	PostUpdate()
// 	Size() (int, int)

// 	// GetClipRect() *TRect
// }

type (
	// IEvent -
	IEvent interface {
		Time() time.Time
		Type() string
		EventKey() string
		Screen() interfaces.IScreen
		String() string
	}

	// // IEvent -
	// IEvent interface {
	// 	hal.IEvent
	// }

	// TEvent -
	TEvent struct {
		IEvent
		time   time.Time
		screen interfaces.IScreen
	}
)

// NewEvent -
func NewEvent(scr interfaces.IScreen) *TEvent {
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
func (o *TEvent) Screen() interfaces.IScreen {
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
	x, y int
}

// NewKeyboardEvent -
func NewKeyboardEvent(scr interfaces.IScreen, id int, ch rune, scan int, x, y, mod int) *TKeyboardEvent {
	ret := &TKeyboardEvent{TWindowEvent: *NewWindowEvent(scr, id), x: x, y: y, ch: ch, scan: scan, mod: mod}
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
	format := "%v U+%x %v %v x:%v y:%v"
	if unicode.IsPrint(o.ch) {
		format = "%v %q %v %v x:%v y:%v"
	}
	return fmt.Sprintf(format, o.TWindowEvent.String(), o.ch, o.scan, o.mod, o.x, o.y)
}

// TDropFileEvent -
type TDropFileEvent struct {
	TWindowEvent
	content string
}

// NewDropFileEvent -
func NewDropFileEvent(scr interfaces.IScreen, id int, s string) *TDropFileEvent {
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
func NewWindowEvent(scr interfaces.IScreen, id int) *TWindowEvent {
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
func NewWindowCloseEvent(scr interfaces.IScreen, id int) *TWindowCloseEvent {
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
	dw, dh   int
}

// NewWindowResizedEvent -
func NewWindowResizedEvent(scr interfaces.IScreen, id, w, h, dw, dh int) *TWindowResizedEvent {
	ret := &TWindowResizedEvent{TWindowEvent: *NewWindowEvent(scr, id), w: w, h: h, dw: dw, dh: dh}
	ret.IEvent = ret

	return ret
}

// Size -
func (o *TWindowResizedEvent) Size() (int, int) {
	return o.w, o.h
}

// Delta _
func (o *TWindowResizedEvent) Delta() (int, int) {
	return o.dw, o.dh
	// scrW, scrH := scr.Size()
	// return o.w - scrW, o.h - scrH
}

// Type -
func (o *TWindowResizedEvent) Type() string {
	return "window resized"
}

// EventKey -
func (o *TWindowResizedEvent) EventKey() string {
	return fmt.Sprintf("window resized")
	// return fmt.Sprintf("window resized %v", o.windowID)
}

// String -
func (o *TWindowResizedEvent) String() string {
	return fmt.Sprintf("%v %vx%v", o.TWindowEvent.String(), o.w, o.h)
}

// TMouseMotionEvent -
type TMouseMotionEvent struct {
	TWindowEvent
	x, y       int
	xRel, yRel int
	state      int
}

// NewMouseMotionEvent -
func NewMouseMotionEvent(scr interfaces.IScreen, winID, x, y, xRel, yRel, state int) *TMouseMotionEvent {
	ret := &TMouseMotionEvent{TWindowEvent: *NewWindowEvent(scr, winID), x: x, y: y, xRel: xRel, yRel: yRel, state: state}
	ret.IEvent = ret
	return ret
}

// Type -
func (o *TMouseMotionEvent) Type() string {
	return "mouse motion"
}

// EventKey -
func (o *TMouseMotionEvent) EventKey() string {
	return "mouse motion"
}

// String -
func (o *TMouseMotionEvent) String() string {
	return fmt.Sprintf("%v X:%v Y:%v XRel:%v YRel:%v State:%v", o.TWindowEvent.String(), o.x, o.y, o.xRel, o.yRel, o.state)
}

// Pos -
func (o *TMouseMotionEvent) Pos() (int, int) { return o.x, o.y }

// TMouseButtonEvent -
type TMouseButtonEvent struct {
	TMouseMotionEvent
	button int
	key    rune
	press  int
	clicks int
}

// NewMouseButtonEvent -
func NewMouseButtonEvent(scr interfaces.IScreen, winID, x, y, state, button int) *TMouseButtonEvent {
	ret := &TMouseButtonEvent{TMouseMotionEvent: TMouseMotionEvent{TWindowEvent: *NewWindowEvent(scr, winID), x: x, y: y, state: state}, button: button}
	ret.IEvent = ret
	return ret
}

// Button -
func (o *TMouseButtonEvent) Button() int {
	return o.button
}

// Type -
func (o *TMouseButtonEvent) Type() string {
	return "mouse"
}

// String -
func (o *TMouseButtonEvent) String() string {
	// return fmt.Sprintf("%v X:%v Y:%v type:%v state:%v", o.x, o.y, o.TWindowEvent.String(), o.press, o.state)
	return fmt.Sprintf("%v X:%v Y:%v key: %v clicks:%v", o.TWindowEvent.String(), o.x, o.y, o.key, o.clicks)

}

// EventKey -
func (o *TMouseButtonEvent) EventKey() string {
	return fmt.Sprintf("%v", o.button)
}
