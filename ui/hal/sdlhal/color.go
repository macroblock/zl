package sdlhal

import "github.com/veandco/go-sdl2/sdl"

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
