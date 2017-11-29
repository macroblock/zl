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
type TRect sdl.Rect

// NewEmptyRect -
func NewEmptyRect() *TRect {
	return &TRect{}
}

// NewRect -
func NewRect(x, y, w, h int) *TRect {
	return &TRect{X: int32(x), Y: int32(y), W: int32(w), H: int32(h)}
}

// Sdl -
func (o *TRect) Sdl() *sdl.Rect {
	return (*sdl.Rect)(o)
}

// SetPos -
func (o *TRect) SetPos(x, y int) *TRect {
	o.X = int32(x)
	o.Y = int32(y)
	return o
}

// SetSize -
func (o *TRect) SetSize(w, h int) {
	o.W = int32(w)
	o.H = int32(h)
}

// SetBounds -
func (o *TRect) SetBounds(x, y, w, h int) {
	o.X = int32(x)
	o.Y = int32(y)
	o.W = int32(w)
	o.H = int32(h)
}

// Pos -
func (o *TRect) Pos() (x, y int) { return int(o.X), int(o.Y) }

// Size -
func (o *TRect) Size() (w, h int) { return int(o.W), int(o.H) }

// Bounds -
func (o *TRect) Bounds() (x, y, w, h int) { return int(o.X), int(o.Y), int(o.W), int(o.H) }

// SetPos32 -
func (o *TRect) SetPos32(x, y int32) {
	o.X = x
	o.Y = y
}

// SetSize32 -
func (o *TRect) SetSize32(w, h int32) *TRect {
	o.W = w
	o.H = h
	return o
}

// SetBounds32 -
func (o *TRect) SetBounds32(x, y, w, h int32) {
	o.X = x
	o.Y = y
	o.W = w
	o.H = h
}

// Pos32 -
func (o *TRect) Pos32() (x, y int32) { return o.X, o.Y }

// Size32 -
func (o *TRect) Size32() (w, h int32) { return o.W, o.H }

// Bounds32 -
func (o *TRect) Bounds32() (x, y, w, h int32) { return o.X, o.Y, o.W, o.H }

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
