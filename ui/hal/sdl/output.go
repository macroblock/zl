package hal

import (
	"github.com/macroblock/zl/types/vector"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// TOutput -
type TOutput struct {
	hal        *THal
	x, y, w, h int
	window     *sdl.Window
	renderer   *sdl.Renderer
	fillColor  sdl.Color
	drawColor  sdl.Color
	children   vector.TVector
	font       *ttf.Font
}

// IWidget -
type IWidget interface {
	Draw()
	Output() *TOutput
	SetParent(widget IWidget)
}

// Close -
func (o *TOutput) Close() {
	//_, err := hal.outputs.Remove(hal.outputs.IndexOf(o))
	//log.Warning(err, "TOutput.Close(): output not found")
	id, err := o.window.GetID()
	log.Error(err, "TOutput.Close(): Window.GetID")
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
func (o *TOutput) AddChild(child IWidget) {
	o.children.PushBack(child)
	child.SetParent(o)
}

// SetParent -
func (o *TOutput) SetParent(widget IWidget) {
	log.Panic(true, "SetParent() should not be called")
}

// Output -
func (o *TOutput) Output() *TOutput {
	return o
}

// Draw -
func (o *TOutput) Draw() {
	// o.renderer.SetDrawColor(100, 200, 0, 255)
	// o.renderer.FillRect(&sdl.Rect{X: 10, Y: 10, W: 40, H: 15})
	// o.renderer.Present()
	for i := 0; i < o.children.Len(); i++ {
		x, err := o.children.At(i)
		log.Error(err, "Draw() something wrong")
		child, ok := x.(IWidget)
		log.Error(!ok, "Draw() something wrong 2")
		child.Draw()
	}
}

// SetDrawColor -
func (o *TOutput) SetDrawColor(r, g, b, a int) {
	if o == nil {
		return
	}
	o.drawColor.R = uint8(r)
	o.drawColor.G = uint8(g)
	o.drawColor.B = uint8(b)
	o.drawColor.A = uint8(a)
}

// SetFillColor -
func (o *TOutput) SetFillColor(r, g, b, a int) {
	if o == nil {
		return
	}
	o.fillColor.R = uint8(r)
	o.fillColor.G = uint8(g)
	o.fillColor.B = uint8(b)
	o.fillColor.A = uint8(a)
}

func (o *TOutput) setDrawColor() {
	o.renderer.SetDrawColor(o.drawColor.R, o.drawColor.G, o.drawColor.B, o.drawColor.A)
}

func (o *TOutput) setFillColor() {
	o.renderer.SetDrawColor(o.fillColor.R, o.fillColor.G, o.fillColor.B, o.fillColor.A)
}

// DrawText -
func (o *TOutput) DrawText(s string) {
	var surfaceText *sdl.Surface
	var textureText *sdl.Texture
	err := error(nil)
	surfaceText, err = o.Font().RenderUTF8Blended(s, o.drawColor)
	log.Error(err != nil, "DrawText")
	defer surfaceText.Free()
	textureText, err = o.renderer.CreateTextureFromSurface(surfaceText)
	o.renderer.Copy(textureText, nil, nil)
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
	if o == nil {
		return
	}
	o.setFillColor()
	o.renderer.Clear()
}

// FillRect -
func (o *TOutput) FillRect(x1, y1, w, h int) {
	if o == nil {
		return
	}
	o.setFillColor()
	o.renderer.FillRect(&sdl.Rect{X: int32(x1), Y: int32(y1), W: int32(w), H: int32(h)})
}

// DrawLine -
func (o *TOutput) DrawLine(x1, y1, x2, y2 int) {
	if o == nil {
		return
	}
	o.setDrawColor()
	o.renderer.DrawLine(int32(x1), int32(y1), int32(x2), int32(y2))
}

// DrawRect -
func (o *TOutput) DrawRect(x1, y1, w, h int) {
	if o == nil {
		return
	}
	o.setDrawColor()
	o.renderer.DrawRect(&sdl.Rect{X: int32(x1), Y: int32(y1), W: int32(w), H: int32(h)})
}

// Flush -
func (o *TOutput) Flush() {
	if o == nil {
		return
	}
	o.renderer.Present()
}
