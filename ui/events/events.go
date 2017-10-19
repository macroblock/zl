package events

import (
	"fmt"
	"time"
	"unicode"
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

// TKeyboardEvent -
type TKeyboardEvent struct {
	TEvent
	ch   rune
	scan int
	mod  int
}

// NewKeyboardEvent -
func NewKeyboardEvent(ch rune, mod int) *TKeyboardEvent {
	ret := &TKeyboardEvent{TEvent: *NewEvent(), ch: ch, mod: mod}
	ret.IEvent = ret
	return ret
}

// Type -
func (o *TKeyboardEvent) Type() string {
	return "keyboard"
}

// Rune -
func (o *TKeyboardEvent) Rune() rune {
	return o.ch
}

// ScanCode -
func (o *TKeyboardEvent) ScanCode() int {
	return o.scan
}

// EventKey -
func (o *TKeyboardEvent) EventKey() string {
	return fmt.Sprintf("%q", o.ch)
}

// String -
func (o *TKeyboardEvent) String() string {
	format := "%v: U+%x %v"
	if unicode.IsPrint(o.ch) {
		format = "%v: %q %v"
	}
	return fmt.Sprintf(format, o.TEvent.String(), o.ch, o.mod)
}
