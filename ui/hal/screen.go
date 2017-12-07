package hal

import (
	"github.com/macroblock/zl/types"
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
	SetClipRect(rect *types.TRect) error
	// GetClipRect() *TRect
}