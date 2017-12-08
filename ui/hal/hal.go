package hal

import (
	"github.com/macroblock/zl/core/zlog"
	"github.com/macroblock/zl/ui/events"
)

var log = zlog.Instance("hal")

// IHal -
type IHal interface {
	Close()
	Draw()
	NewScreen() (IScreen, error)
	Screen(id int) IScreen
	Event() events.IEvent
}
