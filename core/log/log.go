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
type TLevel uint

// General loglevel flags
const (
	//All   TLogLevel = uint(-1)
	Quiet TLevel = 0
	Panic        = 1 << iota
	Error
	Warning
	Recover
	Notice
	Info
	Debug
	outOfRange
)

var levelToStr = []string{"UNKNOWN", "PNC", "ERR", "WRN", "RECOVER", "NTC", "INF", "DBG"}

func (o TLevel) maxLevel() TLevel {
	return ^(o - 1) & o & (outOfRange - 1)
}

func (o TLevel) minLevel() TLevel {
	i := 1
	for x := o & (outOfRange - 1); x != 0; x >>= 1 {
		i <<= 1
	}
	i >>= 1
	return TLevel(i)
}

// func (O TLevel) Leq

// String -
func (o TLevel) String() string {
	s := []string{}
	i := 1
	for x := o & (outOfRange - 1); x != 0; x >>= 1 {
		if x&1 != 0 {
			s = append(s, levelToStr[i])
		}
	}
	return strings.Join(s, "|")
}

// TLog -
type TLog struct {
	writers  []TLogger
	hasError bool
}

// TLogger -
type TLogger struct {
	io.Writer
	level    TLevel
	prefixes []string
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
		sl = append(sl, l.level.String()+": "+strings.Join(l.prefixes, ","))
	}
	return strings.Join(sl, "\n")
}

// AddLogger -
func (o *TLog) AddLogger(level TLevel, prefixes []string, w io.Writer) {
	o.writers = append(o.writers, TLogger{Writer: w, level: level, prefixes: prefixes})
}

func errPrefix(hasError bool) string {
	if hasError {
		return "###"
	}
	return ""
}

// Log -
func (o *TLog) Log(level TLevel, prefix string, err error, text ...interface{}) {
	level = level.maxLevel()
	for _, writer := range o.writers {
		if writer.level&level == 0 {
			continue
		}
		if level&Recover != 0 {
			o.hasError = false
		}
		msg := fmt.Sprintf("%v %v %v: %v\n", time.Now().Format("2006-01-02 15:04:05"), errPrefix(o.hasError), level.String(), fmt.Sprint(text...))
		if err != nil {
			o.hasError = true
			msg = fmt.Sprintf("%v        error: %v\n", msg, err.Error())
		}
		if _, err := writer.Write([]byte(msg)); err != nil {
			// TODO: smarter
			fmt.Println(err)
		}
	}
	if level&Panic != 0 {
		panic(fmt.Sprint(text...))
	}
}

// Recover -
func (o *TLog) Recover(text ...interface{}) {

}

// HasError -
func (o *TLog) HasError() bool {
	return o.hasError
}

// // TLogger -
// type TLogger struct {
// 	writer   io.Writer
// 	logLevel TLevel
// }

// func (o *TLogger) Close() {
// }

// func (o *TLogger) Log(level TLevel, err error, text ...interface{}) {
// }

// ILog -
type ILog interface {
	Log(level TLevel, err error, text ...interface{})
}

func init() {
	defaultLog = New()
}
