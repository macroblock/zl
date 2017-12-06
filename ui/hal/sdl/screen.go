package hal

import (
	"github.com/macroblock/zl/types/vector"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// IScreen -
type IScreen interface {
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
	PostUpdate()
	// GetClipRect() *TRect
}

// TScreen -
type TScreen struct {
	hal          *THal
	zeroX, zeroY int
	x, y         int
	window       *sdl.Window
	renderer     *sdl.Renderer
	fillColor    TColor
	drawColor    TColor
	children     vector.TVector
	font         *ttf.Font
	needUpdate   bool
}

var _ IScreen = (*TScreen)(nil)

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
		Bounds() *TRect
	}

	// IClientRect -
	IClientRect interface {
		ClientRect() *TRect
	}
)

// Close -
func (o *TScreen) Close() {
	//_, err := hal.outputs.Remove(hal.outputs.IndexOf(o))
	//log.Warning(err, "TOutput.Close(): output not found")
	id, err := o.window.GetID()
	log.Debug("close window id:", id)
	log.Error(err, "TOutput.Close(): Window.GetID")
	if o == Screen() {
		makeCurrent(stubScreen)
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

func (o *TScreen) drawBounds(vp, cr *TRect) {
	SetClipRect(o, nil)

	o.SetDrawColor(255, 0, 0, 255)
	r := vp.Copy()
	r.Grow(2)
	o.DrawRect(r.Bounds())

	r = cr.Copy()
	r.Grow(1)
	o.SetDrawColor(0, 255, 0, 255)
	o.DrawRect(r.Bounds())
}

func (o *TScreen) drawChildren(children []interface{}, clipRect *TRect) {
	scrW, scrH := o.Size()
	oldX, oldY := o.GetZeroPoint()
	for _, i := range children {
		bounds := NewRect(-oldX, -oldY, scrW, scrH)
		if child, ok := i.(IBounds); ok {
			bounds = child.Bounds()
		}
		bx, by := bounds.Pos()
		cr := clipRect.Copy()
		if ok := cr.Intersect(bounds); !ok {
			cr.SetSize(0, 0)
		}
		// o.drawBounds(bounds, cr) //debug
		SetClipRect(o, cr)
		o.MoveZeroPoint(bx, by)
		if child, ok := i.(IDraw); ok {
			child.Draw()
		}
		cr.Move(-bx, -by)
		if child, ok := i.(IClientRect); ok {
			cb := child.ClientRect()
			cbx, cby := cb.Pos()
			if ok := cr.Intersect(cb); !ok {
				cr.SetSize(0, 0)
			}
			cr.Move(-cbx, -cby)
			o.MoveZeroPoint(cbx, cby)
		}
		if child, ok := i.(IChildren); ok {
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
	o.drawChildren(o.children.Data(), NewRect(0, 0, w, h))
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
	o.renderer.SetDrawColor(o.drawColor.R, o.drawColor.G, o.drawColor.B, o.drawColor.A)
}

func (o *TScreen) setFillColor() {
	o.renderer.SetDrawColor(o.fillColor.R, o.fillColor.G, o.fillColor.B, o.fillColor.A)
}

// DrawText -
func (o *TScreen) DrawText(s string, x, y int) {
	var surfaceText *sdl.Surface
	var textureText *sdl.Texture
	x += o.zeroX
	y += o.zeroY
	err := error(nil)
	surfaceText, err = o.Font().RenderUTF8Blended(s, o.drawColor.Sdl())

	log.Error(err, surfaceText, " DrawText")
	if surfaceText == nil {
		return
	}
	log.Debug(surfaceText.W)
	defer surfaceText.Free()
	textureText, err = o.renderer.CreateTextureFromSurface(surfaceText)
	rect := NewEmptyRect().
		SetPos(x, y).
		SetSize32(surfaceText.W, surfaceText.H)
	o.renderer.Copy(textureText, nil, rect.Sdl())
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
	rect := NewRect(x1, y1, w, h)
	o.renderer.FillRect(rect.Sdl())
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
	rect := NewRect(x1, y1, w, h)
	o.renderer.DrawRect(rect.Sdl())
	o.PostUpdate()
}

// SetClipRect -
func SetClipRect(scr *TScreen, rect *TRect) error {
	if rect == nil {
		return scr.renderer.SetClipRect(nil)
	}
	r := *rect
	r.Move(scr.zeroX, scr.zeroY)
	return scr.renderer.SetClipRect(r.Sdl())
}

// Size -
func (o *TScreen) Size() (int, int) {
	w, h, err := o.renderer.GetOutputSize()
	log.Error(err, "Size(): GetOutputSize")
	return int(w), int(h)
}

// Flush -
func (o *TScreen) Flush() {
	o.renderer.Present()
}
