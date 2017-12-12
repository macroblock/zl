package widget

import (
	"github.com/macroblock/zl/core/zlog"
	"github.com/macroblock/zl/types"
	"github.com/macroblock/zl/ui/hal"
	"github.com/macroblock/zl/ui/hal/sdlhal"
)

var log = zlog.Instance("widget")

// TWidget -
type TWidget struct {
	name     string
	rect     types.TRect
	color    hal.TColor
	children types.TVector
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

// SetName -
func (o *TWidget) SetName(s string) *TWidget {
	o.name = s
	return o
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
func (o *TWidget) AddChild(v ...sdlhal.IDraw) *TWidget {
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
func (o *TWidget) Bounds() *types.TRect {
	r := o.rect
	return &r
}

// AddPos -
func (o *TWidget) AddPos(x, y int) *TWidget {
	o.rect.X += x
	o.rect.Y += y
	return o
}

// ClientRect -
func (o *TWidget) ClientRect() *types.TRect {
	r := types.NewRect(0, 0, o.rect.W, o.rect.H)
	r.Y += 20
	r.H -= 20
	r.Shrink(5)
	return r
}

// Draw -
func (o *TWidget) Draw() {
	scr := sdlhal.Screen()
	log.Error(scr == nil, "Output is nil")
	scr.SetFillColor(50, 60, 70, 255)
	r := types.NewRect(0, 0, o.rect.W, o.rect.H)
	scr.FillRect(r.Bounds())

	scr.SetDrawColor(0, 0, 0, 255)
	scr.DrawText(o.name, 7, 4)
	scr.SetDrawColor(192, 192, 192, 255)
	scr.DrawText(o.name, 5, 2)
	scr.SetDrawColor(100, 100, 100, 255)
	scr.DrawRect(r.Bounds())

	scr.SetFillColor(o.color.RGBA())
	r.Y += 20
	r.H -= 20
	r.Shrink(4)
	scr.FillRect(r.Bounds())

	scr.SetDrawColor(100, 100, 100, 255)
	scr.DrawRect(r.Bounds())

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
	o.rect.SetPos(o.rect.X+x, o.rect.Y+y)
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
func (o *TTextWidget) Bounds() *types.TRect {
	r := o.rect
	return &r
}

// Draw -
func (o *TTextWidget) Draw() {
	scr := sdlhal.Screen()
	log.Error(scr == nil, "Output is nil")
	scr.SetFillColor(o.color.RGBA())
	scr.FillRect(0, 0, o.rect.W, o.rect.H)
	scr.SetDrawColor(0, 200, 255, 255)
	scr.DrawText(o.s, 0, 0)
	scr.DrawLine(0, 0, o.rect.W, o.rect.H)
	scr.SetDrawColor(100, 100, 100, 255)
	scr.DrawRect(0, 0, o.rect.W, o.rect.H)
}
