package event

import "time"

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
	scan int
	mod  int
	ch   rune
}
