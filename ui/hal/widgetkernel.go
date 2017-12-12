package hal

import (
	"github.com/macroblock/zl/ui/hal/interfaces"
)

// IWidgetKernel -
type IWidgetKernel interface {
	Parent() IWidgetKernel
	Screen() interfaces.IScreen
}

// TWidgetKernel -
type TWidgetKernel struct {
	scr    interfaces.IScreen
	parent IWidgetKernel
}

// InitWidgetKernel -
func (o *TWidgetKernel) InitWidgetKernel() {
	o.scr = StubScreen()
}

// Parent -
func (o *TWidgetKernel) Parent() IWidgetKernel {
	return o.parent
}

// Screen -
func (o *TWidgetKernel) Screen() interfaces.IScreen {
	return o.scr
}

// SetScreen -
func (o *TWidgetKernel) SetScreen(scr interfaces.IScreen) {
	o.scr = scr
}

// SetParent -
func (o *TWidgetKernel) SetParent(parent IWidgetKernel) {
	o.parent = parent
}
