package zlog

import (
	"fmt"
	"strings"
	"time"

	"github.com/macroblock/zl/core/loglevel"
	"github.com/macroblock/zl/core/zlogger"
)

// Default -
var defaultLog *TLog

type tNode struct {
	hasError bool
	loggers  []*zlogger.TLogger
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

// Instance -
func Instance(name string) *TLog {
	if defaultLog == nil {
		defaultLog = New(name)
		return defaultLog
	}
	ret := *defaultLog
	ret.name = name
	return &ret
}

// Instance -
func (o *TLog) Instance(name string) *TLog {
	if o == nil {
		o = New(name)
		return o
	}
	ret := *o
	ret.name = name
	return &ret
}

// HasError -
func (o *TLog) HasError() bool {
	return o.node.hasError
}

// String -
func (o *TLog) String() string {
	sl := []string{}
	for _, l := range o.node.loggers {
		sl = append(sl, l.Filter().String()) //+": "+strings.Join(l.prefixes, ","))
	}
	return strings.Join(sl, "\n")
}

// Add -
func (o *TLog) Add(logger ...*zlogger.TLogger) {
	// TODO: check on nil
	o.node.loggers = append(o.node.loggers, logger...)
}

// Log -
func (o *TLog) Log(level loglevel.TLevel, err error, text ...interface{}) {
	formatParams := zlogger.NewFormatParams(time.Now(), level, fmt.Sprint(text...), err, o.node.hasError, o.name)
	_ = formatParams
	for _, writer := range o.node.loggers {
		if level.NotIn(writer.Filter()) {
			continue
		}
		if level == loglevel.Recover {
			o.node.hasError = false
		}
		if err != nil {
			o.node.hasError = true
		}
		msg := zlogger.FormatLog(writer.Styler()(writer.Format(), level, o.name, o.node.hasError, err, text...))
		if _, err := writer.Writer().Write([]byte(msg)); err != nil {
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
