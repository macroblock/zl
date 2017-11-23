package hal

import (
	"github.com/veandco/go-sdl2/sdl"
)

// TOutput -
type TOutput struct {
	hal        *THal
	x, y, w, h int
	window     *sdl.Window
	renderer   *sdl.Renderer
	fillColor  sdl.Color
	drawColor  sdl.Color
}

// Close -
func (o *TOutput) Close() {
	//_, err := hal.outputs.Remove(hal.outputs.IndexOf(o))
	//log.Warning(err, "TOutput.Close(): output not found")
	id := o.window.GetID()
	log.Error(id == 0, "TOutput.Close(): Window.GetID")
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

// Draw -
func (o *TOutput) Draw() {
	// o.renderer.SetDrawColor(100, 200, 0, 255)
	// o.renderer.FillRect(&sdl.Rect{X: 10, Y: 10, W: 40, H: 15})
	// o.renderer.Present()
}

// SetDrawColor -
func (o *TOutput) SetDrawColor(r, g, b, a int) {
	o.drawColor.R = uint8(r)
	o.drawColor.G = uint8(g)
	o.drawColor.B = uint8(b)
	o.drawColor.A = uint8(a)
}

// SetFillColor -
func (o *TOutput) SetFillColor(r, g, b, a int) {
	o.fillColor.R = uint8(r)
	o.fillColor.G = uint8(g)
	o.fillColor.B = uint8(b)
	o.fillColor.A = uint8(a)
}

func (o *TOutput) setDrawColor() {
	err := o.renderer.SetDrawColor(o.drawColor.R, o.drawColor.G, o.drawColor.B, o.drawColor.A)
	log.Warning(err, "SetDrawColor")
}

func (o *TOutput) setFillColor() {
	o.renderer.SetDrawColor(o.fillColor.R, o.fillColor.G, o.fillColor.B, o.fillColor.A)
}

// Clear -
func (o *TOutput) Clear() {
	o.setFillColor()
	o.renderer.Clear()
}

// FillRect -
func (o *TOutput) FillRect(x1, y1, w, h int) {
	o.setFillColor()
	o.renderer.FillRect(&sdl.Rect{X: int32(x1), Y: int32(y1), W: int32(w), H: int32(h)})
}

// DrawLine -
func (o *TOutput) DrawLine(x1, y1, x2, y2 int) {
	o.setDrawColor()
	o.renderer.DrawLine(int32(x1), int32(y1), int32(x2), int32(y2))
}

// DrawRect -
func (o *TOutput) DrawRect(x1, y1, w, h int) {
	o.setDrawColor()
	o.renderer.DrawRect(&sdl.Rect{X: int32(x1), Y: int32(y1), W: int32(w), H: int32(h)})
}

func (o *TOutput) Flush() {
	o.renderer.Present()
}
