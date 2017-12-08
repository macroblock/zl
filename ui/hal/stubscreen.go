package hal

import (
	"github.com/macroblock/zl/types"
	"github.com/veandco/go-sdl2/ttf"
)

var stubScreen IScreen = &tStubOutput{}

type tStubOutput struct {
}

var _ IScreen = (*tStubOutput)(nil)

// StubScreen -
func StubScreen() IScreen {
	return stubScreen
}

func errMsg() { log.Error(true, "Output is not initialized") }

// Close -
func (o *tStubOutput) Close() {}

// AddChild -
func (o *tStubOutput) AddChild(children ...interface{}) { errMsg() }

// Draw -
func (o *tStubOutput) Draw() { errMsg() }

// SetDrawColor -
func (o *tStubOutput) SetDrawColor(r, g, b, a int) { errMsg() }

// SetFillColor -
func (o *tStubOutput) SetFillColor(r, g, b, a int) { errMsg() }

// DrawText -
func (o *tStubOutput) DrawText(s string, x, y int) { errMsg() }

// Font -
func (o *tStubOutput) Font() *ttf.Font { errMsg(); return nil }

// SetFont -
func (o *tStubOutput) SetFont(font *ttf.Font) { errMsg() }

// Clear -
func (o *tStubOutput) Clear() { errMsg() }

// FillRect -
func (o *tStubOutput) FillRect(x1, y1, w, h int) { errMsg() }

// DrawLine -
func (o *tStubOutput) DrawLine(x1, y1, x2, y2 int) { errMsg() }

// DrawRect -
func (o *tStubOutput) DrawRect(x1, y1, w, h int) { errMsg() }

// PostUpdate -
func (o *tStubOutput) PostUpdate() { errMsg() }

// Size -
func (o *tStubOutput) Size() (int, int) { errMsg(); return 0, 0 }

// Flush -
func (o *tStubOutput) Flush() { errMsg() }

// SetClipRect -
func (o *tStubOutput) SetClipRect(rect *types.TRect) error { errMsg(); return nil }
