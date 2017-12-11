package events

// ToKeyboard -
func ToKeyboard(ev IEvent) *TKeyboardEvent {
	return ev.(*TKeyboardEvent)
}

// ToWindow -
func ToWindow(ev IEvent) *TWindowEvent {
	return ev.(*TWindowEvent)
}

// ToWindowClose -
func ToWindowClose(ev IEvent) *TWindowCloseEvent {
	return ev.(*TWindowCloseEvent)
}

// ToWindowResized -
func ToWindowResized(ev IEvent) *TWindowResizedEvent {
	return ev.(*TWindowResizedEvent)
}

// ToMouse -
func ToMouse(ev IEvent) *TMouseButtonEvent {
	return ev.(*TMouseButtonEvent)
}
