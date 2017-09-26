package log

import (
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/macroblock/zl/core/loglevel"
)

// Default -
var defaultLog *TLog

// TLog -
type TLog struct {
	node *struct{ writers []TLogger }
	//writers  []TLogger
	hasError bool
	prefix   string
}

// IFormat -
type IFormat interface {
	format(level loglevel.TLevel, prefix string, err error, wasErr bool, text ...interface{}) string
}

// TLogger -
type TLogger struct {
	IFormat
	io.Writer
	filter   loglevel.TFilter
	prefixes []string
}

func (o *TLogger) format(level loglevel.TLevel, prefix string, err error, wasErr bool, text ...interface{}) string {
	msg := fmt.Sprintf("%v (%v) %v%v: %v\n", time.Now().Format("2006-01-02 15:04:05"), prefix, errPrefix(wasErr), level.String(), fmt.Sprint(text...))
	if err != nil {
		msg = fmt.Sprintf("%v    Cause: %v\n", msg, err.Error())
	}
	return msg
}

// New -
func New() *TLog {
	return &TLog{node: &struct{ writers []TLogger }{}}
}

// Default -
func Default() *TLog {
	return defaultLog
}

// Clone -
func (o *TLog) Clone() *TLog {
	log := TLog{}
	log = *o
	return &log
}

// SetPrefix -
func (o *TLog) SetPrefix(s string) TLog {
	o.prefix = s
	return *o
}

// String -
func (o *TLog) String() string {
	sl := []string{}
	for _, l := range o.node.writers {
		sl = append(sl, l.filter.String()+": "+strings.Join(l.prefixes, ","))
	}
	return strings.Join(sl, "\n")
}

// AddLogger -
func (o *TLog) AddLogger(filter loglevel.TFilter, prefixes []string, w io.Writer) {
	o.node.writers = append(o.node.writers, TLogger{Writer: w, filter: filter, prefixes: prefixes})
}

func errPrefix(hasError bool) string {
	if hasError {
		return "#"
	}
	return ""
}

// Log -
func (o *TLog) Log(level loglevel.TLevel, prefix string, err error, text ...interface{}) {
	for _, writer := range o.node.writers {
		if level.NotIn(writer.filter) {
			continue
		}
		if level == loglevel.Recover {
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
	if level == loglevel.Panic {
		panic(fmt.Sprint(text...))
	}
}

// Panic -
func (o *TLog) Panic(err error, text ...interface{}) {
	o.Log(loglevel.Panic, o.prefix, err, text...)
}

// Error -
func (o *TLog) Error(err error, text ...interface{}) {
	o.Log(loglevel.Error, o.prefix, err, text...)
}

// Warning -
func (o *TLog) Warning(err error, text ...interface{}) {
	o.Log(loglevel.Warning, o.prefix, err, text...)
}

// Recover -
func (o *TLog) Recover(text ...interface{}) {
	o.Log(loglevel.Recover, o.prefix, nil, text...)
}

// Notice -
func (o *TLog) Notice(text ...interface{}) {
	o.Log(loglevel.Notice, o.prefix, nil, text...)
}

// Info -
func (o *TLog) Info(text ...interface{}) {
	o.Log(loglevel.Info, o.prefix, nil, text...)
}

// Debug -
func (o *TLog) Debug(text ...interface{}) {
	o.Log(loglevel.Debug, o.prefix, nil, text...)
}

// HasError -
func (o *TLog) HasError() bool {
	return o.hasError
}

func init() {
	defaultLog = New()
	defaultLog.SetPrefix("Main")
}
