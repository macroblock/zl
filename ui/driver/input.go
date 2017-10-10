package driver

import "time"

// IEvent -
type IEvent interface {
	Time() time.Time
}

// IInput -
type IInput interface {
	Event() IEvent
}

// IGraph -
type IGraph interface {
}
