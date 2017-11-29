package widget

import (
	"github.com/macroblock/zl/core/zlog"
	hal "github.com/macroblock/zl/ui/hal/sdl"
)

var log = zlog.Instance("widget")

// TWidget -
type TWidget struct {
	x, y                           int
	w, h                           int
	colorR, colorG, colorB, colorA int
}

//TTextWidget -
type TTextWidget struct {
	TWidget
	s string
}

// NewWidget -
func NewWidget() *TWidget {
	ret := &TWidget{w: 50, h: 50}
	return ret
}

// SetPos -
func (o *TWidget) SetPos(x, y int) {
	o.x = x
	o.y = y
}

// SetColor -
func (o *TWidget) SetColor(r, g, b, a int) {
	o.colorR = r
	o.colorG = g
	o.colorB = b
	o.colorA = a
}

// Draw -
func (o *TWidget) Draw() {
	scr := hal.Output()
	log.Error(scr == nil, "Output is nil")
	scr.SetFillColor(o.colorR, o.colorG, o.colorB, o.colorA)
	scr.FillRect(o.x, o.y, o.w, o.h)
	scr.SetDrawColor(100, 100, 100, 255)
	scr.DrawRect(o.x, o.y, o.w, o.h)
}

// NewTextWidget -
func NewTextWidget() *TTextWidget {
	ret := &TTextWidget{TWidget: *NewWidget()}

	return ret
}

func (o *TTextWidget) SetText(s string) {
	o.s = s
}

// Draw -
func (o *TTextWidget) Draw() {
	scr := hal.Output()
	log.Error(scr == nil, "Output is nil")
	scr.SetFillColor(o.colorR, o.colorG, o.colorB, o.colorA)
	scr.FillRect(o.x, o.y, o.w, o.h)
	scr.DrawText(o.s, o.x, o.y)
	scr.SetDrawColor(100, 100, 100, 255)
	scr.DrawRect(o.x, o.y, o.w, o.h)
}
