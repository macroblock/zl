package log

import (
	"fmt"
	"io"
	"strings"
	"time"
)

// Default -
var defaultLog *TLog

// TLevel -
type TLevel int

// TFilter -
type TFilter uint

// General loglevel flags
const (
	TooLow TLevel = -1 + iota
	Panic
	Error
	Warning
	Recover
	Notice
	Info
	Debug
	TooHigh
)

var levelToStr = []string{"PNC", "ERR", "WRN", "RECOVER", "NTC", "INF", "DBG", "UNKNOWN"}

// Filter -
func (o TLevel) Filter() TFilter { return 1 << uint(o) }

// Below -
func (o TLevel) Below() TFilter { return o.Filter() - 1 }

// Above -
func (o TLevel) Above() TFilter { return TooHigh.Below() &^ o.OrLower() }

// OrLower -
func (o TLevel) OrLower() TFilter { return o.Filter() | (o.Filter() - 1) }

// OrHigher -
func (o TLevel) OrHigher() TFilter { return TooHigh.Below() &^ o.Below() }

// In -
func (o TLevel) In(f TFilter) bool { return f&o.Filter() != 0 }

// NotIn -
func (o TLevel) NotIn(f TFilter) bool { return f&o.Filter() == 0 }

// String -
func (o TLevel) String() string {
	if o <= TooLow || o >= TooHigh {
		o = TooHigh
	}
	return levelToStr[o]
}

// Include -
func (o TFilter) Include(f TFilter) TFilter { return o | f }

// Exclude -
func (o TFilter) Exclude(f TFilter) TFilter { return o &^ f }

// String -
func (o TFilter) String() string {
	sl := []string{}
	for i, x := 0, o&TooHigh.Below(); x != 0; i, x = i+1, x>>1 {
		if x&1 != 0 {
			sl = append(sl, levelToStr[i])
		}
	}
	if len(sl) == 0 {
		return levelToStr[TooHigh]
	}
	return strings.Join(sl, "|")
}

// TLog -
type TLog struct {
	writers  []TLogger
	hasError bool
}

// IFormat -
type IFormat interface {
	format(level TLevel, prefix string, err error, wasErr bool, text ...interface{}) string
}

// TLogger -
type TLogger struct {
	IFormat
	io.Writer
	filter   TFilter
	prefixes []string
}

func (o *TLogger) format(level TLevel, prefix string, err error, wasErr bool, text ...interface{}) string {
	msg := fmt.Sprintf("%v %v %v: %v\n", time.Now().Format("2006-01-02 15:04:05"), errPrefix(wasErr), level.String(), fmt.Sprint(text...))
	if err != nil {
		msg = fmt.Sprintf("%v        error: %v\n", msg, err.Error())
	}
	return msg
}

// New -
func New() *TLog {
	return &TLog{}
}

// Default -
func Default() *TLog {
	return defaultLog
}

func (o *TLog) String() string {
	sl := []string{}
	for _, l := range o.writers {
		sl = append(sl, l.filter.String()+": "+strings.Join(l.prefixes, ","))
	}
	return strings.Join(sl, "\n")
}

// AddLogger -
func (o *TLog) AddLogger(filter TFilter, prefixes []string, w io.Writer) {
	o.writers = append(o.writers, TLogger{Writer: w, filter: filter, prefixes: prefixes})
}

func errPrefix(hasError bool) string {
	if hasError {
		return "###"
	}
	return ""
}

// Log -
func (o *TLog) Log(level TLevel, prefix string, err error, text ...interface{}) {
	for _, writer := range o.writers {
		if level.NotIn(writer.filter) {
			continue
		}
		if level == Recover {
			o.hasError = false
		}
		if err != nil {
			o.hasError = true
		}
		msg := writer.format(level, prefix, err, o.hasError, text...)
		if _, err := writer.Write([]byte(msg)); err != nil {
			// TODO: smarter
			fmt.Println(err)
		}
	}
	if level == Panic {
		panic(fmt.Sprint(text...))
	}
}

// Recover -
func (o *TLog) Recover(text ...interface{}) {
}

// Info -
func (o *TLog) Info(text ...interface{}) {
	o.Log(Info, "", nil, text...)
}

// HasError -
func (o *TLog) HasError() bool {
	return o.hasError
}

func init() {
	defaultLog = New()
}
