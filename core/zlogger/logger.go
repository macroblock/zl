package zlogger

import (
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/macroblock/zl/core/loglevel"
)

const defaultFormat = "~t (~n)~w~l: ~x~e\n"

// TStyler -
type TStyler func(format string, level loglevel.TLevel, name string, wasErr bool, err error, text ...interface{}) (formatStr, timeStr, levelStr, nameStr, wasErrStr, errStr, textStr string)

// Supported stylers
var (
	DefaultStyler TStyler = defaultStyler
	AnsiStyler    TStyler = ansiStyler
)

// TLogger -
type TLogger struct {
	name     string
	writer   io.Writer
	styler   TStyler
	filter   loglevel.TFilter
	format   string
	prefixes []string
}

func errPrefix(hasError bool) string {
	if hasError {
		return "#"
	}
	return ""
}

// Filter -
func (o *TLogger) Filter() loglevel.TFilter {
	return o.filter
}

// Format -
func (o *TLogger) Format() string {
	return o.format
}

// Writer -
func (o *TLogger) Writer() io.Writer {
	return o.writer
}

// Styler -
func (o *TLogger) Styler() TStyler {
	return o.styler
}

// FormatLog -
// ~t - current time
// ~l - loglevel
// ~n - module name
// ~w - wasErr string
// ~e - error text
// ~x - text
// ~t (~n)~w~l: ~x~e
func FormatLog(format string, time, level, name, wasErr, err, text string) string {
	ret := format
	ret = strings.Replace(ret, "~e", err, 1)
	ret = strings.Replace(ret, "~x", text, 1)
	ret = strings.Replace(ret, "~l", level, 1)
	ret = strings.Replace(ret, "~w", wasErr, 1)
	ret = strings.Replace(ret, "~n", name, 1)
	ret = strings.Replace(ret, "~t", time, 1)
	return ret
}

func defaultStyler(format string, level loglevel.TLevel, name string, wasErr bool, err error, text ...interface{}) (formatStr, timeStr, levelStr, nameStr, wasErrStr, errStr, textStr string) {
	formatStr = format
	timeStr = time.Now().Format("2006-01-02 15:04:05")
	levelStr = level.String()
	nameStr = name
	textStr = fmt.Sprint(text...)
	if err != nil {
		errStr = fmt.Sprintf("\n    +Cause: %v", err.Error())
	}
	wasErrStr = " "
	if wasErr {
		wasErrStr = "!"
	}
	return
}

func ansiStyler(format string, level loglevel.TLevel, name string, wasErr bool, err error, text ...interface{}) (formatStr, timeStr, levelStr, nameStr, wasErrStr, errStr, textStr string) {
	reset := "\x1b[0m"
	color := ""
	switch level {
	case loglevel.Debug:
		color = "\x1b[1;30m" // bright black
	case loglevel.Info:
		color = "\x1b[0m" // reset //white (lightgrey)
	case loglevel.Notice:
		color = "\x1b[1;32m" // bright green
	case loglevel.Recover:
		color = "\x1b[1;36m" // bright cyan
	case loglevel.Warning:
		color = "\x1b[1;33m" // bright yellow
	case loglevel.Error:
		color = "\x1b[1;31m" // bright red
	case loglevel.Panic:
		color = "\x1b[1;31m" // bright red
	}
	formatStr = color + format + reset
	timeStr = time.Now().Format("2006-01-02 15:04:05")
	levelStr = level.String()
	nameStr = name
	textStr = fmt.Sprint(text...)
	if err != nil {
		errStr = fmt.Sprintf("\n    +Cause: %v", err.Error())
	}
	wasErrStr = " "
	if wasErr {
		wasErrStr = "!"
	}
	return
}
