package events

const _mA = "actions "

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

// type (
// 	// TWindowAction -
// 	TWindowAction struct {
// 		TAction
// 		handler TWindowHandler
// 	}
// 	// TWindowHandler -
// 	TWindowHandler func(ev TWindowEvent) bool
// )

// type (
// 	// TWindowResizedAction -
// 	TWindowResizedAction struct {
// 		TAction
// 		handler TWindowResizedHandler
// 	}
// 	// TWindowResizedHandler -
// 	TWindowResizedHandler func(ev TWindowResizedEvent) bool
// )

// type (
// 	// TWindowCloseAction -
// 	TWindowCloseAction struct {
// 		TAction
// 		handler TWindowCloseHandler
// 	}
// 	// TWindowCloseHandler -
// 	TWindowCloseHandler func(ev TWindowCloseEvent) bool
// )

// // TKeyboardAction -
// type (
// 	TKeyboardAction struct {
// 		TAction
// 		handler TKeyboardHandler
// 	}
// 	// TKeyboardHandler -
// 	TKeyboardHandler func(ev TKeyboardEvent) bool
// )

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

// // NewKeyboardAction -
// func NewKeyboardAction(name, eventKey, descr string, handler TKeyboardHandler) IAction {
// 	act := &TKeyboardAction{}
// 	act.name = name
// 	act.eventKey = eventKey
// 	act.description = descr
// 	act.handler = handler
// 	ActionMap.Add(act)
// 	return act
// }

// // Do -
// func (o *TKeyboardAction) Do(ev IEvent) bool {
// 	if o.handler == nil {
// 		return false
// 	}
// 	keybEvent, ok := ev.(*TKeyboardEvent)
// 	log.Warning(!ok, _mA, "Do: incompatible event <", ev, "> in <", o, "> method")
// 	return o.handler(*keybEvent)
// }

// // NewWindowAction -
// func NewWindowAction(name, eventKey, descr string, handler TWindowHandler) IAction {
// 	act := &TWindowAction{}
// 	act.name = name
// 	act.eventKey = eventKey
// 	act.description = descr
// 	act.handler = handler
// 	ActionMap.Add(act)
// 	return act
// }

// // NewWindowResizedAction -
// func NewWindowResizedAction(name, eventKey, descr string, handler TWindowResizedHandler) IAction {
// 	act := &TWindowResizedAction{}
// 	act.name = name
// 	act.eventKey = eventKey
// 	act.description = descr
// 	act.handler = handler
// 	ActionMap.Add(act)
// 	return act
// }

// // Do -
// func (o *TWindowResizedAction) Do(ev IEvent) bool {
// 	if o.handler == nil {
// 		return false
// 	}
// 	winResize, ok := ev.(*TWindowResizedEvent)
// 	log.Debug(winResize)
// 	log.Warning(!ok, _mA, "Do: incompatible WindowResizedEvent <", ev, "> in <", o, "> method")
// 	return o.handler(*winResize)
// }

// // NewWindowCloseAction -
// func NewWindowCloseAction(name, eventKey, descr string, handler TWindowCloseHandler) IAction {
// 	act := &TWindowCloseAction{}
// 	act.name = name
// 	act.eventKey = eventKey
// 	act.description = descr
// 	act.handler = handler
// 	ActionMap.Add(act)
// 	return act
// }

// // Do -
// func (o *TWindowCloseAction) Do(ev IEvent) bool {
// 	if o.handler == nil {
// 		return false
// 	}
// 	winResize, ok := ev.(*TWindowCloseEvent)
// 	log.Debug(winResize)
// 	log.Warning(!ok, _mA, "Do: incompatible WindowCloseEvent <", ev, "> in <", o, "> method")
// 	return o.handler(*winResize)
// }
