package sdlhal

import "github.com/veandco/go-sdl2/sdl"

// TCursor -
type TCursor struct {
	cursor *sdl.Cursor
}

// ShowCursor -
func ShowCursor(id int) {
	sdl.ShowCursor(id)
}

// NewCursor -
func NewCursor(id sdl.SystemCursor) *TCursor {
	return &TCursor{cursor: sdl.CreateSystemCursor(id)}
}

// SetCursor -
func (o *TCursor) SetCursor(id *sdl.Cursor) {
	o.cursor = id
}

// Free -
func (o *TCursor) Free() {
	sdl.FreeCursor(o.cursor)
}
