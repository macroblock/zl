package zlog

import (
	"fmt"
	"io"
	"strings"

	"github.com/macroblock/zl/core/loglevel"
	"github.com/macroblock/zl/core/zlogger"
)

// Default -
var defaultLog = New("main")

type tNode struct {
	hasError bool
	writers  []zlogger.TLogger
}

// TLog -
type TLog struct {
	node *tNode
	name string
}

// New -
func New(name string) *TLog {
	return &TLog{name: name, node: &tNode{}}
}

// Get -
func Get() *TLog {
	return defaultLog
}

// Clone -
func Clone(name string) *TLog {
	return defaultLog.Clone(name)
}

// Clone -
func (o *TLog) Clone(name string) *TLog {
	log := *o
	log.name = name
	return &log
}

// HasError -
func (o *TLog) HasError() bool {
	return o.node.hasError
}

// String -
func (o *TLog) String() string {
	sl := []string{}
	for _, l := range o.node.writers {
		sl = append(sl, l.Filter().String()) //+": "+strings.Join(l.prefixes, ","))
	}
	return strings.Join(sl, "\n")
}

// AddLogger -
func (o *TLog) AddLogger(filter loglevel.TFilter, prefixes []string, w io.Writer) {
	o.node.writers = append(o.node.writers, *zlogger.New(filter, prefixes, w))
}

// Log -
func (o *TLog) Log(level loglevel.TLevel, err error, text ...interface{}) {
	for _, writer := range o.node.writers {
		if level.NotIn(writer.Filter()) {
			continue
		}
		if level == loglevel.Recover {
			o.node.hasError = false
		}
		if err != nil {
			o.node.hasError = true
		}
		msg := writer.Format(level, o.name, o.node.hasError, err, text...)
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
	o.Log(loglevel.Panic, err, text...)
}

// Error -
func (o *TLog) Error(err error, text ...interface{}) {
	o.Log(loglevel.Error, err, text...)
}

// Warning -
func (o *TLog) Warning(err error, text ...interface{}) {
	o.Log(loglevel.Warning, err, text...)
}

// Recover -
func (o *TLog) Recover(text ...interface{}) {
	o.Log(loglevel.Recover, nil, text...)
}

// Notice -
func (o *TLog) Notice(text ...interface{}) {
	o.Log(loglevel.Notice, nil, text...)
}

// Info -
func (o *TLog) Info(text ...interface{}) {
	o.Log(loglevel.Info, nil, text...)
}

// Debug -
func (o *TLog) Debug(text ...interface{}) {
	o.Log(loglevel.Debug, nil, text...)
}
