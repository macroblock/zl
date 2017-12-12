package hal

// TColor -
type TColor struct {
	R, G, B, A int
}

// NewEmptyColor -
func NewEmptyColor() *TColor {
	return &TColor{}
}

// NewColor -
func NewColor(r, g, b, a int) *TColor {
	return &TColor{R: r, G: g, B: b, A: a}
}

// SetRGBA -
func (o *TColor) SetRGBA(r, g, b, a int) {
	o.R = r
	o.G = g
	o.B = b
	o.A = a
}

// RGBA -
func (o *TColor) RGBA() (r, g, b, a int) {
	return o.R, o.G, o.B, o.A
}

// RGBA8 -
func (o *TColor) RGBA8() (r, g, b, a uint8) {
	return uint8(o.R), uint8(o.G), uint8(o.B), uint8(o.A)
}
