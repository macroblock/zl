package action

import "github.com/macroblock/zl/ui/event"

// TActionMap -
type TActionMap struct {
	byName map[string]IAction
}

type (
	// TAction -
	TAction struct {
		IAction
		name        string
		eventKey    string
		description string
		handler     TEventHandler
	}

	// TEventHandler -
	TEventHandler func(ev event.IEvent) bool

	// IAction -
	IAction interface {
		Name() string
		EventKey() string
		Description() string
		Do(ev event.IEvent) bool
	}
)

// Name -
func (o *TAction) Name() string {
	return o.name
}

// EventKey -
func (o *TAction) EventKey() string {
	return o.eventKey
}

// Description -
func (o *TAction) Description() string {
	return o.description
}

// Do -
func (o *TAction) Do(ev event.IEvent) bool {
	if o.handler == nil {
		return false
	}
	return o.handler(ev)
}
