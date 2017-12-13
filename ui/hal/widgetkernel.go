package hal

import (
	"github.com/macroblock/zl/ui/hal/interfaces"
)

// TWidgetKernel -
type TWidgetKernel struct {
	scr    interfaces.IScreen
	parent interfaces.IWidgetKernel
}

// InitWidgetKernel -
func (o *TWidgetKernel) InitWidgetKernel() {
	// o.scr = StubScreen()
}

// Parent -
func (o *TWidgetKernel) Parent() interfaces.IWidgetKernel {
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
func (o *TWidgetKernel) SetParent(parent interfaces.IWidgetKernel) {
	o.parent = parent
}
