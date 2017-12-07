package types

import "github.com/macroblock/zl/utils"

// TRect -
type TRect struct {
	X, Y int
	W, H int
}

// NewEmptyRect -
func NewEmptyRect() *TRect {
	return &TRect{}
}

// NewRect -
func NewRect(x, y, w, h int) *TRect {
	return &TRect{X: x, Y: y, W: w, H: h}
}

// Empty - reports whether a rectangle has no area
func (o *TRect) Empty() bool {
	return o == nil || o.W <= 0 || o.H <= 0
}

// Equals - reports whether two rectangles are equal
func (o *TRect) Equals(b *TRect) bool {
	if (o != nil) && (b != nil) &&
		(o.X == b.X) && (o.Y == b.Y) &&
		(o.W == b.W) && (o.H == b.H) {
		return true
	}
	return false
}

// Intersect calculates the intersection of two rectangles.
func (o *TRect) Intersect(b *TRect) bool {
	if o.Empty() || b.Empty() {
		o.W = 0
		o.H = 0
		return false
	}
	o.W = utils.Min(o.X+o.W, b.X+b.W)
	o.H = utils.Min(o.Y+o.H, b.Y+b.H)
	o.X = utils.Max(o.X, b.X)
	o.Y = utils.Max(o.Y, b.Y)
	o.W -= o.X
	o.H -= o.Y
	if o.W <= 0 || o.H <= 0 {
		o.W = 0
		o.H = 0
		return false
	}
	return true
}

// Copy -
func (o *TRect) Copy() *TRect {
	r := *o
	return &r
}

// SetPos -
func (o *TRect) SetPos(x, y int) *TRect {
	o.X = x
	o.Y = y
	return o
}

// SetSize -
func (o *TRect) SetSize(w, h int) {
	o.W = w
	o.H = h
}

// SetBounds -
func (o *TRect) SetBounds(x, y, w, h int) {
	o.X = x
	o.Y = y
	o.W = w
	o.H = h
}

// Move -
func (o *TRect) Move(dx, dy int) {
	o.X += dx
	o.Y += dy
}

// Bounds -
func (o *TRect) Bounds() (x, y, w, h int) {
	return o.X, o.Y, o.W, o.H
}

// Grow -
func (o *TRect) Grow(n int) {
	o.X -= n
	o.Y -= n
	o.W += n * 2
	o.H += n * 2
}

// Shrink -
func (o *TRect) Shrink(n int) {
	o.X += n
	o.Y += n
	o.W -= n * 2
	o.H -= n * 2
}
