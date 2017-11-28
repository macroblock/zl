package widget

import (
	"github.com/macroblock/zl/core/zlog"
	hal "github.com/macroblock/zl/ui/hal/sdl"
)

var log = zlog.Instance("widget")

// TWidget
type TWidget struct {
	parent                         hal.IWidget
	output                         *hal.TOutput
	x, y                           int
	w, h                           int
	colorR, colorG, colorB, colorA int
}

// NewWidget -
func NewWidget() (*TWidget, error) {
	ret := &TWidget{w: 50, h: 50}
	return ret, nil
}

// SetParent -
func (o *TWidget) SetParent(widget hal.IWidget) {
	o.parent = widget
	o.output = widget.Output()
}

func (o *TWidget) Output() *hal.TOutput {
	return o.output
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
	log.Error(o.output == nil, "Output is nil")
	o.output.SetFillColor(o.colorR, o.colorG, o.colorB, o.colorA)
	o.output.FillRect(o.x, o.y, o.w, o.h)
	o.output.SetDrawColor(100, 100, 100, 255)
	o.output.DrawRect(o.x, o.y, o.w, o.h)
}
