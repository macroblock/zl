package hal

import "github.com/veandco/go-sdl2/sdl"

type idestroyer interface {
	Destroy()
}

func destroy(o interface{}) {
	// if o == nil {
	// 	log.Warning(true, "destroy: can't get address of argument <nil>")
	// 	return
	// }
	switch v := o.(type) {
	case nil:
	case idestroyer:
		v.Destroy()
	default:
		log.Warning(true, "destroy: variable hasn't got Destroy() method")
	}
	o = nil
}

// TRect -
type TRect struct {
	Rect sdl.Rect
}

// NewEmptyRect -
func NewEmptyRect() *TRect {
	return &TRect{}
}

// NewRect -
func NewRect(x, y, w, h int) *TRect {
	return &TRect{sdl.Rect{X: int32(x), Y: int32(y), W: int32(w), H: int32(h)}}
}

// Sdl -
func (o *TRect) Sdl() *sdl.Rect {
	return &o.Rect
}

// SetPos -
func (o *TRect) SetPos(x, y int) *TRect {
	o.Rect.X = int32(x)
	o.Rect.Y = int32(y)
	return o
}

// SetSize -
func (o *TRect) SetSize(w, h int) {
	o.Rect.W = int32(w)
	o.Rect.H = int32(h)
}

// SetBounds -
func (o *TRect) SetBounds(x, y, w, h int) {
	o.Rect.X = int32(x)
	o.Rect.Y = int32(y)
	o.Rect.W = int32(w)
	o.Rect.H = int32(h)
}

// Move -
func (o *TRect) Move(dx, dy int) {
	o.Rect.X += int32(dx)
	o.Rect.Y += int32(dy)
}

// Pos -
func (o *TRect) Pos() []int {
	ret := []int{int(o.Rect.X), int(o.Rect.Y)}
	return ret
}

// X -
func (o *TRect) X() int {
	return int(o.Rect.X)
}

// Y -
func (o *TRect) Y() int {
	return int(o.Rect.Y)
}

// W -
func (o *TRect) W() int {
	return int(o.Rect.W)
}

// H -
func (o *TRect) H() int {
	return int(o.Rect.H)
}

// Size -
func (o *TRect) Size() (w, h int) { return int(o.Rect.W), int(o.Rect.H) }

// Bounds -
func (o *TRect) Bounds() (x, y, w, h int) {
	return int(o.Rect.X), int(o.Rect.Y), int(o.Rect.W), int(o.Rect.H)
}

// SetPos32 -
func (o *TRect) SetPos32(x, y int32) {
	o.Rect.X = x
	o.Rect.Y = y
}

// SetSize32 -
func (o *TRect) SetSize32(w, h int32) *TRect {
	o.Rect.W = w
	o.Rect.H = h
	return o
}

// SetBounds32 -
func (o *TRect) SetBounds32(x, y, w, h int32) {
	o.Rect.X = x
	o.Rect.Y = y
	o.Rect.W = w
	o.Rect.H = h
}

// Pos32 -
func (o *TRect) Pos32() (x, y int32) { return o.Rect.X, o.Rect.Y }

// Size32 -
func (o *TRect) Size32() (w, h int32) { return o.Rect.W, o.Rect.H }

// Bounds32 -
func (o *TRect) Bounds32() (x, y, w, h int32) { return o.Rect.X, o.Rect.Y, o.Rect.W, o.Rect.H }

// type TLine {x1,y1,x2,y2 int}  //линии??

// TColor -
type TColor sdl.Color

// NewEmptyColor -
func NewEmptyColor() *TColor {
	return &TColor{}
}

// NewColor -
func NewColor(r, g, b, a int) *TColor {
	return &TColor{R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(a)}
}

// Sdl -
func (o *TColor) Sdl() sdl.Color {
	return *(*sdl.Color)(o)
}

// SetRGBA -
func (o *TColor) SetRGBA(r, g, b, a int) {
	o.R = uint8(r)
	o.G = uint8(g)
	o.B = uint8(b)
	o.A = uint8(a)
}

// RGBA -
func (o *TColor) RGBA() (r, g, b, a int) {
	return int(o.R), int(o.G), int(o.B), int(o.A)
}

// RGBA8 -
func (o *TColor) RGBA8() (r, g, b, a uint8) {
	return o.R, o.G, o.B, o.A
}
