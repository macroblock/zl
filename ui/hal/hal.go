package hal

import (
	"github.com/macroblock/zl/core/zlog"
	"github.com/macroblock/zl/ui/events"
	"github.com/macroblock/zl/ui/hal/interfaces"
)

var log = zlog.Instance("hal")

// IHal -
type IHal interface {
	Close()
	Draw()
	NewScreen() (interfaces.IScreen, error)
	Screen(id int) interfaces.IScreen
	Event() events.IEvent
}

// // IEvent -
// type IEvent interface {
// 	Time() time.Time
// 	Type() string
// 	EventKey() string
// 	Screen() IScreen
// 	String() string
// }

// ToWidgetKernel -
func ToWidgetKernel(slice []interface{}) []interfaces.IWidgetKernel {
	ret := make([]interfaces.IWidgetKernel, len(slice))
	for i, v := range slice {
		ret[i] = v.(interfaces.IWidgetKernel)
	}
	return ret
}
