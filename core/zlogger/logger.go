package zlogger

import (
	"fmt"
	"io"
	"time"

	"github.com/macroblock/zl/core/loglevel"
)

// IFormat -
type IFormat interface {
	format(level loglevel.TLevel, name string, wasErr bool, err error, text ...interface{}) string
}

// TLogger -
type TLogger struct {
	IFormat
	io.Writer
	filter   loglevel.TFilter
	prefixes []string
}

var (
	_ IFormat   = (*TLogger)(nil)
	_ io.Writer = (*TLogger)(nil)
)

func errPrefix(hasError bool) string {
	if hasError {
		return "#"
	}
	return ""
}

// New -
func New(filter loglevel.TFilter, prefixes []string, w io.Writer) *TLogger {
	return &TLogger{filter: filter, prefixes: prefixes, Writer: w}
}

// Filter -
func (o *TLogger) Filter() loglevel.TFilter {
	return o.filter
}

// Format -
func (o *TLogger) Format(level loglevel.TLevel, name string, wasErr bool, err error, text ...interface{}) string {
	msg := fmt.Sprintf("%v (%v) %v%v: %v\n", time.Now().Format("2006-01-02 15:04:05"), name, errPrefix(wasErr), level.String(), fmt.Sprint(text...))
	if err != nil {
		msg = fmt.Sprintf("%v    Cause: %v\n", msg, err.Error())
	}
	return msg
}
