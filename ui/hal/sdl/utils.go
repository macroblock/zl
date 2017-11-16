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

// Rect -
type Rect sdl.Rect

// NewRect -
func NewRect(x, y, w, h int) *Rect {
	return &Rect{X: int32(x), Y: int32(y), W: int32(w), H: int32(h)}
}

// SetPos -
func (r *Rect) SetPos(x, y int) {
	r.X = int32(x)
	r.Y = int32(y)
}

// SetSize -
func (r *Rect) SetSize(w, h int) {
	r.W = int32(w)
	r.H = int32(h)
}

// SetBounds -
func (r *Rect) SetBounds(x, y, w, h int) {
	r.X = int32(x)
	r.Y = int32(y)
	r.W = int32(w)
	r.H = int32(h)
}

// Pos -
func (r *Rect) Pos() (x, y int) { return int(r.X), int(r.Y) }

// Size -
func (r *Rect) Size() (w, h int) { return int(r.W), int(r.H) }

// Bounds -
func (r *Rect) Bounds() (x, y, w, h int) { return int(r.X), int(r.Y), int(r.W), int(r.H) }

// SetPos32 -
func (r *Rect) SetPos32(x, y int32) {
	r.X = x
	r.Y = y
}

// SetSize32 -
func (r *Rect) SetSize32(w, h int32) {
	r.W = w
	r.H = h
}

// SetBounds32 -
func (r *Rect) SetBounds32(x, y, w, h int32) {
	r.X = x
	r.Y = y
	r.W = w
	r.H = h
}

// Pos32 -
func (r *Rect) Pos32() (x, y int32) { return r.X, r.Y }

// Size32 -
func (r *Rect) Size32() (w, h int32) { return r.W, r.H }

// Bounds32 -
func (r *Rect) Bounds32() (x, y, w, h int32) { return r.X, r.Y, r.W, r.H }
