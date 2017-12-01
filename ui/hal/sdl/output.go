package hal

import (
	"github.com/macroblock/zl/types/vector"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// IOutput -
type IOutput interface {
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
}

// TOutput -
type TOutput struct {
	hal        *THal
	x, y, w, h int
	window     *sdl.Window
	renderer   *sdl.Renderer
	fillColor  TColor
	drawColor  TColor
	children   vector.TVector
	font       *ttf.Font
}

var _ IOutput = (*TOutput)(nil)

// IDraw -
type IDraw interface {
	Draw()
}

// IChildren -
type IChildren interface {
	Children() []interface{}
}

// IBounds -
type IBounds interface {
	Bounds() *TRect
}

// Close -
func (o *TOutput) Close() {
	//_, err := hal.outputs.Remove(hal.outputs.IndexOf(o))
	//log.Warning(err, "TOutput.Close(): output not found")
	id, err := o.window.GetID()
	log.Debug("close window id:", id)
	log.Error(err, "TOutput.Close(): Window.GetID")
	if o == Output() {
		makeCurrent(stubOutput)
	}
	delete(o.hal.outputs, id)

	if o.renderer != nil {
		o.renderer.Destroy()
		o.renderer = nil
	}
	if o.window != nil {
		o.window.Destroy()
		o.window = nil
	}
}

// AddChild -
func (o *TOutput) AddChild(children ...interface{}) {
	for _, child := range children {
		o.children.PushBack(child)
	}
}

func (o *TOutput) drawChildren(children []interface{}, dx, dy, dw, dh int) {
	for _, i := range children {
		if child, ok := i.(IDraw); ok {
			child.Draw()
		}
		rect := NewEmptyRect()
		if child, ok := i.(IBounds); ok {
			rect = child.Bounds()
			rect.Move(dx, dy)
		}
		oldRect := o.GetViewport()
		log.Error(o.SetViewport(rect), "drawChildren: SetViewport")
		rect = o.GetViewport()

		if child, ok := i.(IChildren); ok {
			log.Error(o.SetClipRect(oldRect), "drawChildren: SetClipRect")
			// o.SetClipRect(NewRect(0, 0, 30, 30))
			o.drawChildren(child.Children(), rect.X(), rect.Y(), rect.W(), rect.H())
			r := o.GetClipRect()
			o.SetClipRect(nil)
			o.SetDrawColor(255, 0, 0, 255)
			o.DrawRect(o.GetViewport().Bounds())
			o.SetDrawColor(0, 255, 0, 255)
			o.DrawRect(r.Bounds())
			o.SetClipRect(nil)
		}
		log.Error(o.SetViewport(oldRect), "drawChildren: SetViewport")
		o.DrawRect(o.GetViewport().Bounds())
	}
}

// Draw -
func (o *TOutput) Draw() {
	o.drawChildren(o.children.Data(), 0, 0, o.w, o.h)
}

// SetDrawColor -
func (o *TOutput) SetDrawColor(r, g, b, a int) {
	o.drawColor.SetRGBA(r, g, b, a)
}

// SetFillColor -
func (o *TOutput) SetFillColor(r, g, b, a int) {
	o.fillColor.SetRGBA(r, g, b, a)
}

func (o *TOutput) setDrawColor() {
	o.renderer.SetDrawColor(o.drawColor.R, o.drawColor.G, o.drawColor.B, o.drawColor.A)
}

func (o *TOutput) setFillColor() {
	o.renderer.SetDrawColor(o.fillColor.R, o.fillColor.G, o.fillColor.B, o.fillColor.A)
}

// DrawText -
func (o *TOutput) DrawText(s string, x, y int) {
	var surfaceText *sdl.Surface
	var textureText *sdl.Texture
	err := error(nil)
	surfaceText, err = o.Font().RenderUTF8Blended(s, o.drawColor.Sdl())
	log.Error(err != nil, "DrawText")
	defer surfaceText.Free()
	textureText, err = o.renderer.CreateTextureFromSurface(surfaceText)
	rect := NewEmptyRect().
		SetPos(x, y).
		SetSize32(surfaceText.W, surfaceText.H)
	o.renderer.Copy(textureText, nil, rect.Sdl())
}

// Font -
func (o *TOutput) Font() *ttf.Font {
	return o.font
}

// SetFont -
func (o *TOutput) SetFont(font *ttf.Font) {
	o.font = font
}

// Clear -
func (o *TOutput) Clear() {
	o.setFillColor()
	o.renderer.Clear()
}

// FillRect -
func (o *TOutput) FillRect(x1, y1, w, h int) {
	o.setFillColor()
	// o.renderer.FillRect(&sdl.Rect{X: int32(x1), Y: int32(y1), W: int32(w), H: int32(h)})
	rect := NewRect(x1, y1, w, h)
	o.renderer.FillRect(rect.Sdl())
}

// DrawLine -
func (o *TOutput) DrawLine(x1, y1, x2, y2 int) {
	o.setDrawColor()
	o.renderer.DrawLine(int32(x1), int32(y1), int32(x2), int32(y2))
}

// DrawRect -
func (o *TOutput) DrawRect(x1, y1, w, h int) {
	o.setDrawColor()
	// o.renderer.DrawRect(&sdl.Rect{X: int32(x1), Y: int32(y1), W: int32(w), H: int32(h)})
	rect := NewRect(x1, y1, w, h)
	o.renderer.DrawRect(rect.Sdl())

}

// SetViewport -
func (o *TOutput) SetViewport(rect *TRect) error {
	if rect == nil {
		return o.renderer.SetViewport(nil)
	}
	return o.renderer.SetViewport(rect.Sdl())
}

// GetViewport -
func (o *TOutput) GetViewport() *TRect {
	return &TRect{Rect: o.renderer.GetViewport()}
}

// SetClipRect -
func (o *TOutput) SetClipRect(rect *TRect) error {
	if rect == nil {
		return o.renderer.SetClipRect(nil)
	}
	return o.renderer.SetClipRect(rect.Sdl())
}

// GetClipRect -
func (o *TOutput) GetClipRect() *TRect {
	return &TRect{Rect: o.renderer.GetClipRect()}
}

// Flush -
func (o *TOutput) Flush() {
	o.renderer.Present()
}
