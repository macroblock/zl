package zlogger

import (
	"fmt"
	"io"
	"os"

	"github.com/macroblock/zl/core/loglevel"
)

var loggerNumber int

// tBuild -
type tBuild struct {
	logger *TLogger
}

func autoName() string {
	ret := fmt.Sprintf("logger%v", loggerNumber)
	loggerNumber++
	return ret
}

// Default -
func Default() *TLogger {
	return &TLogger{Writer: os.Stdout, IStyler: DefaultStyler, name: autoName(), filter: loglevel.All, format: defaultFormat}
}

// Build -
func Build() tBuild {
	return tBuild{logger: Default()}
}

// Done -
func (o tBuild) Done() ILogger {
	return o.logger
}

// SetName -
func (o tBuild) SetName(name string) tBuild {
	o.logger.name = name
	return o
}

// SetFilter -
func (o tBuild) SetFilter(filter loglevel.TFilter) tBuild {
	o.logger.filter = filter
	return o
}

// SetWriter -
func (o tBuild) SetWriter(writer io.Writer) tBuild {
	o.logger.Writer = writer
	return o
}

// SetStyler -
func (o tBuild) SetStyler(styler IStyler) tBuild {
	o.logger.IStyler = styler
	return o
}

// SetFormat -
func (o tBuild) SetFormat(format string) tBuild {
	o.logger.format = format
	return o
}
