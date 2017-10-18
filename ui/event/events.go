package event

import (
	"fmt"
	"time"
)

type (
	// IEvent -
	IEvent interface {
		Time() time.Time
		Type() string
		EventKey() string
		String() string
	}

	// TEvent -
	TEvent struct {
		IEvent
		time time.Time
	}
)

// NewEvent -
func NewEvent() *TEvent {
	ret := &TEvent{time: time.Now()}
	ret.IEvent = ret
	return ret
}

// Time -
func (o *TEvent) Time() time.Time {
	return o.time
}

// Type -
func (o *TEvent) Type() string {
	return "unknown"
}

// EventKey -
func (o *TEvent) EventKey() string {
	return ""
}

// String -
func (o *TEvent) String() string {
	return o.time.Format("15:04:05.000") + " " + o.IEvent.Type()
}

// TKeyboard -
type TKeyboard struct {
	TEvent
	ch   rune
	scan int
	mod  int
}

// NewKeyboard -
func NewKeyboard(ch rune, mod int) *TKeyboard {
	ret := &TKeyboard{TEvent: *NewEvent(), ch: ch, mod: mod}
	ret.IEvent = ret
	return ret
}

// Type -
func (o *TKeyboard) Type() string {
	return "keyboard"
}

// Rune -
func (o *TKeyboard) Rune() rune {
	return o.ch
}

// ScanCode -
func (o *TKeyboard) ScanCode() int {
	return o.scan
}

// EventKey -
func (o *TKeyboard) EventKey() string {
	return fmt.Sprintf("%q", o.ch)
}

// String -
func (o *TKeyboard) String() string {
	return fmt.Sprintf("%v: %q %v", o.TEvent.String(), o.ch, o.mod)
}
