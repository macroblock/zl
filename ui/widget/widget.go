package widget

import (
	"github.com/macroblock/zl/core/zlog"
	"github.com/macroblock/zl/types/vector"
	hal "github.com/macroblock/zl/ui/hal/sdl"
)

var log = zlog.Instance("widget")

// TWidget -
type TWidget struct {
	rect     hal.TRect
	color    hal.TColor
	children vector.TVector
}

//TTextWidget -
type TTextWidget struct {
	TWidget
	s string
}

// NewWidget -
func NewWidget() *TWidget {
	ret := &TWidget{}
	ret.rect.SetSize(50, 50)
	return ret
}

// SetPos -
func (o *TWidget) SetPos(x, y int) *TWidget {
	o.rect.SetPos(x, y)
	return o
}

// SetSize -
func (o *TWidget) SetSize(w, h int) *TWidget {
	o.rect.SetSize(w, h)
	return o
}

// SetColor -
func (o *TWidget) SetColor(r, g, b, a int) *TWidget {
	o.color.SetRGBA(r, g, b, a)
	return o
}

// AddChild -
func (o *TWidget) AddChild(v ...hal.IDraw) *TWidget {
	for _, child := range v {
		o.children.PushBack(child)
	}
	return o
}

// Children -
func (o *TWidget) Children() []interface{} {
	return o.children.Data()
}

// Bounds -
func (o *TWidget) Bounds() *hal.TRect {
	return &hal.TRect{Rect: *o.rect.Sdl()}
}

// Draw -
func (o *TWidget) Draw() {
	scr := hal.Output()
	log.Error(scr == nil, "Output is nil")
	scr.SetFillColor(o.color.RGBA())
	scr.FillRect(0, 0, o.rect.W(), o.rect.H())
	scr.SetDrawColor(100, 100, 100, 255)
	scr.DrawRect(0, 0, o.rect.W(), o.rect.H())

}

// NewTextWidget -
func NewTextWidget() *TTextWidget {
	ret := &TTextWidget{TWidget: *NewWidget()}
	return ret
}

// Children -
func (o *TTextWidget) Children() []interface{} {
	return o.children.Data()
}

// SetText -
func (o *TTextWidget) SetText(s string) *TTextWidget {
	o.s = s
	return o
}

// SetPos -
func (o *TTextWidget) SetPos(x, y int) *TTextWidget {
	o.rect.SetPos(x, y)
	return o
}

// AddPos -
func (o *TTextWidget) AddPos(x, y int) *TTextWidget {
	o.rect.SetPos(o.rect.X()+x, o.rect.Y()+y)
	return o
}

// SetSize -
func (o *TTextWidget) SetSize(w, h int) *TTextWidget {
	o.rect.SetSize(w, h)
	return o
}

// SetColor -
func (o *TTextWidget) SetColor(r, g, b, a int) *TTextWidget {
	o.color.SetRGBA(r, g, b, a)
	return o
}

// Bounds -
func (o *TTextWidget) Bounds() *hal.TRect {
	return &hal.TRect{Rect: *o.rect.Sdl()}
}

// Draw -
func (o *TTextWidget) Draw() {
	scr := hal.Output()
	log.Error(scr == nil, "Output is nil")
	scr.SetFillColor(o.color.RGBA())
	scr.FillRect(0, 0, o.rect.W(), o.rect.H())
	scr.SetDrawColor(0, 200, 255, 255)
	scr.DrawText(o.s, 0, 0)
	scr.DrawLine(0, 0, o.rect.W(), o.rect.H())
	scr.SetDrawColor(100, 100, 100, 255)
	scr.DrawRect(0, 0, o.rect.W(), o.rect.H())
}
