package events

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
	TEventHandler func(ev IEvent) bool

	// IAction -
	IAction interface {
		Name() string
		EventKey() string
		Description() string
		Do(ev IEvent) bool
	}
)

// TKeyboardAction -
type (
	TKeyboardAction struct {
		TAction
		handler TKeyboardHandler
	}
	// TKeyboardHandler -
	TKeyboardHandler func(ev TKeyboardEvent) bool
)

// NewAction -
func NewAction(name, eventKey, descr string, handler TEventHandler) IAction {
	act := &TAction{}
	act.name = name
	act.eventKey = eventKey
	act.description = descr
	act.handler = handler
	ActionMap.Add(act)
	return act
}

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
func (o *TAction) Do(ev IEvent) bool {
	if o.handler == nil {
		return false
	}
	return o.handler(ev)
}

// NewKeyboardAction -
func NewKeyboardAction(name, eventKey, descr string, handler TKeyboardHandler) IAction {
	act := &TKeyboardAction{}
	act.name = name
	act.eventKey = eventKey
	act.description = descr
	act.handler = handler
	ActionMap.Add(act)
	return act
}

// Do -
func (o *TKeyboardAction) Do(ev IEvent) bool {
	if o.handler == nil {
		return false
	}
	keybEvent, ok := ev.(*TKeyboardEvent)
	log.Warning(ok, "incompatible event <,", ev, "> in <,", o, "> method")
	return o.handler(*keybEvent)
}
