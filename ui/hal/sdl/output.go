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
	o.renderer.SetDrawColor(100, 200, 0, 255)
	o.renderer.FillRect(&sdl.Rect{X: 10, Y: 10, W: 40, H: 15})
	o.renderer.Present()
}
