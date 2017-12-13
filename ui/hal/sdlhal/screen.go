package sdlhal

import (
	"github.com/macroblock/zl/types"
	"github.com/macroblock/zl/ui/hal"
	"github.com/macroblock/zl/ui/hal/interfaces"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// TScreen -
type TScreen struct {
	hal          *THal
	zeroX, zeroY int
	x, y         int
	oldW, oldH   int
	window       *sdl.Window
	renderer     *sdl.Renderer
	fillColor    hal.TColor
	drawColor    hal.TColor
	children     types.TVector
	font         *ttf.Font
	needUpdate   bool
}

var _ interfaces.IScreen = (*TScreen)(nil)

type (
	// IDraw -
	IDraw interface {
		Draw()
	}

	// IChildren -
	IChildren interface {
		Children() []interface{}
	}

	// IBounds -
	IBounds interface {
		Bounds() *types.TRect
	}

	// IClientRect -
	IClientRect interface {
		ClientRect() *types.TRect
	}

	// IKernel interface{

	// }
)

// Close -
func (o *TScreen) Close() {
	//_, err := hal.outputs.Remove(hal.outputs.IndexOf(o))
	//log.Warning(err, "TOutput.Close(): output not found")
	id, err := o.window.GetID()
	log.Debug("close window id:", id)
	log.Error(err, "TOutput.Close(): Window.GetID")
	if o == Screen() {
		makeCurrent(hal.StubScreen())
	}
	delete(o.hal.screen, id)

	if o.renderer != nil {
		o.renderer.Destroy()
		o.renderer = nil
	}
	if o.window != nil {
		o.window.Destroy()
		o.window = nil
	}
}

// PostUpdate -
func (o *TScreen) PostUpdate() {
	o.needUpdate = true
}

// ResetUpdate -
func (o *TScreen) ResetUpdate() {
	o.needUpdate = false
}

// NeedUpdate -
func (o *TScreen) NeedUpdate() bool {
	return o.needUpdate
}

// AddChild -
func (o *TScreen) AddChild(children ...interface{}) {
	for _, child := range children {
		o.children.PushBack(child)
	}
}

func (o *TScreen) drawBounds(vp, cr *types.TRect) {
	o.SetClipRect(nil)

	o.SetDrawColor(255, 0, 0, 255)
	r := vp.Copy()
	r.Grow(2)
	o.DrawRect(r.Bounds())

	r = cr.Copy()
	r.Grow(1)
	o.SetDrawColor(0, 255, 0, 255)
	o.DrawRect(r.Bounds())
}

// func (o *TScreen) drawChildren(children []interface{}, clipRect *types.TRect) {
// 	scrW, scrH := o.Size()
// 	oldX, oldY := o.GetZeroPoint()
// 	for _, i := range children {
// 		bounds := types.NewRect(-oldX, -oldY, scrW, scrH)
// 		if child, ok := i.(IBounds); ok {
// 			bounds = child.Bounds()
// 		}
// 		cr := clipRect.Copy()
// 		hasIntersect := cr.Intersect(bounds)
// 		o.drawBounds(bounds, cr) //debug
// 		o.SetClipRect(cr)
// 		o.MoveZeroPoint(bounds.X, bounds.Y)
// 		if child, ok := i.(IDraw); ok {
// 			child.Draw()
// 		}
// 		cr.Move(-bounds.X, -bounds.Y)
// 		if child, ok := i.(IClientRect); ok && hasIntersect {
// 			cb := child.ClientRect()
// 			hasIntersect = cr.Intersect(cb)
// 			cr.Move(-cb.X, -cb.Y)
// 			o.MoveZeroPoint(cb.X, cb.Y)
// 		}
// 		if child, ok := i.(IChildren); ok && hasIntersect {
// 			o.drawChildren(child.Children(), cr)
// 		}
// 		o.SetZeroPoint(oldX, oldY)
// 	}
// }

// // Draw -
// func (o *TScreen) Draw() {
// 	o.SetFillColor(0, 0, 0, 0)
// 	o.Clear()
// 	o.SetZeroPoint(0, 0)
// 	w, h := o.Size()
// 	o.drawChildren(o.children.Data(), types.NewRect(0, 0, w, h))
// 	o.SetZeroPoint(0, 0)
// 	log.Debug("_____________________________________________")
// }

// func (o *TScreen) AddChild(parent hal.IWidgetKernel, child interface{}) {

// 	o.children.PushBack(child)
// }

func (o *TScreen) detach(v IDraw) {
	o.children.Remove(o.children.IndexOf(v))
}

func (o *TScreen) drawChildren(children []interface{}, clipRect *types.TRect) {
	scrW, scrH := o.Size()
	oldX, oldY := o.GetZeroPoint()
	for _, i := range children {
		bounds := types.NewRect(-oldX, -oldY, scrW, scrH)
		if child, ok := i.(IBounds); ok {
			bounds = child.Bounds()
		}
		cr := clipRect.Copy()
		hasIntersect := cr.Intersect(bounds)
		o.drawBounds(bounds, cr) //debug
		o.SetClipRect(cr)
		o.MoveZeroPoint(bounds.X, bounds.Y)
		if child, ok := i.(IDraw); ok {
			child.Draw()
		}
		cr.Move(-bounds.X, -bounds.Y)
		if child, ok := i.(IClientRect); ok && hasIntersect {
			cb := child.ClientRect()
			hasIntersect = cr.Intersect(cb)
			cr.Move(-cb.X, -cb.Y)
			o.MoveZeroPoint(cb.X, cb.Y)
		}
		if child, ok := i.(IChildren); ok && hasIntersect {
			o.drawChildren(child.Children(), cr)
		}
		o.SetZeroPoint(oldX, oldY)
	}
}

// Draw -
func (o *TScreen) Draw() {
	o.SetFillColor(0, 0, 0, 0)
	o.Clear()
	o.SetZeroPoint(0, 0)
	w, h := o.Size()
	o.drawChildren(o.children.Data(), types.NewRect(0, 0, w, h))
	o.SetZeroPoint(0, 0)
	log.Debug("_____________________________________________")
}

// SetZeroPoint -
func (o *TScreen) SetZeroPoint(x, y int) {
	o.zeroX = x
	o.zeroY = y
}

// MoveZeroPoint -
func (o *TScreen) MoveZeroPoint(x, y int) {
	o.zeroX += x
	o.zeroY += y
}

// GetZeroPoint -
func (o *TScreen) GetZeroPoint() (x, y int) {
	return o.zeroX, o.zeroY
}

// SetDrawColor -
func (o *TScreen) SetDrawColor(r, g, b, a int) {
	o.drawColor.SetRGBA(r, g, b, a)
}

// SetFillColor -
func (o *TScreen) SetFillColor(r, g, b, a int) {
	o.fillColor.SetRGBA(r, g, b, a)
}

func (o *TScreen) setDrawColor() {
	o.renderer.SetDrawColor(o.drawColor.RGBA8())
}

func (o *TScreen) setFillColor() {
	o.renderer.SetDrawColor(o.fillColor.RGBA8())
}

// DrawText -
func (o *TScreen) DrawText(s string, x, y int) {
	var surfaceText *sdl.Surface
	var textureText *sdl.Texture
	x += o.zeroX
	y += o.zeroY
	err := error(nil)
	color := sdl.Color{}
	color.R, color.G, color.B, color.A = o.drawColor.RGBA8()
	surfaceText, err = o.Font().RenderUTF8Blended(s, color)

	// surfaceText, err = o.Font().RenderUTF8Blended(s, o.drawColor.Sdl())
	log.Error(err, surfaceText, " DrawText")
	if surfaceText == nil {
		return
	}
	defer surfaceText.Free()
	textureText, err = o.renderer.CreateTextureFromSurface(surfaceText)
	rect := sdl.Rect{X: int32(x), Y: int32(y), W: surfaceText.W, H: surfaceText.H}
	o.renderer.Copy(textureText, nil, &rect)
	o.PostUpdate()
}

// Font -
func (o *TScreen) Font() *ttf.Font {
	return o.font
}

// SetFont -
func (o *TScreen) SetFont(font *ttf.Font) {
	o.font = font
}

// Clear -
func (o *TScreen) Clear() {
	o.setFillColor()
	o.renderer.Clear()
}

// FillRect -
func (o *TScreen) FillRect(x1, y1, w, h int) {
	x1 += o.zeroX
	y1 += o.zeroY
	o.setFillColor()
	// o.renderer.FillRect(&sdl.Rect{X: int32(x1), Y: int32(y1), W: int32(w), H: int32(h)})
	rect := sdl.Rect{X: int32(x1), Y: int32(y1), W: int32(w), H: int32(h)}
	o.renderer.FillRect(&rect)
	o.PostUpdate()
}

// DrawLine -
func (o *TScreen) DrawLine(x1, y1, x2, y2 int) {
	x1 += o.zeroX
	y1 += o.zeroY
	x2 += o.zeroX
	y2 += o.zeroY
	o.setDrawColor()
	o.renderer.DrawLine(int32(x1), int32(y1), int32(x2), int32(y2))
	o.PostUpdate()
}

// DrawRect -
func (o *TScreen) DrawRect(x1, y1, w, h int) {
	x1 += o.zeroX
	y1 += o.zeroY
	o.setDrawColor()
	// o.renderer.DrawRect(&sdl.Rect{X: int32(x1), Y: int32(y1), W: int32(w), H: int32(h)})
	rect := sdl.Rect{X: int32(x1), Y: int32(y1), W: int32(w), H: int32(h)}
	o.renderer.DrawRect(&rect)
	o.PostUpdate()
}

// SetClipRect -
func (o *TScreen) SetClipRect(rect *types.TRect) error {
	if rect == nil {
		return o.renderer.SetClipRect(nil)
	}
	r := sdl.Rect{X: int32(rect.X + o.zeroX), Y: int32(rect.Y + o.zeroY), W: int32(rect.W), H: int32(rect.H)}
	return o.renderer.SetClipRect(&r)
}

// Size -
func (o *TScreen) Size() (int, int) {
	w, h, err := o.renderer.GetOutputSize()
	log.Error(err, "Size(): GetOutputSize")
	return int(w), int(h)
}

// OldSize -
func (o *TScreen) OldSize() (int, int) {
	return o.oldW, o.oldH
}

// Flush -
func (o *TScreen) Flush() {
	o.renderer.Present()
}
