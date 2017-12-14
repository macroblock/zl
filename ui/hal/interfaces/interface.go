package interfaces

import (
	"github.com/macroblock/zl/types"
	"github.com/veandco/go-sdl2/ttf"
)

// IScreen -
type IScreen interface {
	Close()
	// AddChild(children ...interface{})
	AddChild(children ...IWidgetKernel)
	Remove(v IWidgetKernel)
	// FindWidget(x, y int, root IScreen) (IWidgetKernel, int, int)
	FindWidget(x, y int) (IWidgetKernel, int, int)
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
	Size() (int, int)
	OldSize() (int, int)
	// GetClipRect() *TRect
}

// IWidgetKernel -
type IWidgetKernel interface {
	Parent() IWidgetKernel
	Screen() IScreen
	SetScreen(scr IScreen)
	SetParent(parent IWidgetKernel)
}
