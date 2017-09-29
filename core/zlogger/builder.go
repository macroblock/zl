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
	return &TLogger{writer: os.Stdout, styler: DefaultStyler, name: autoName(), levelFilter: loglevel.All, format: defaultFormat}
}

// Build -
func Build() tBuild {
	return tBuild{logger: Default()}
}

// Done -
func (o tBuild) Done() *TLogger {
	return o.logger
}

// SetName -
func (o tBuild) SetName(name string) tBuild {
	o.logger.name = name
	return o
}

// SetLevelFilter -
func (o tBuild) SetLevelFilter(filter loglevel.TFilter) tBuild {
	o.logger.levelFilter = filter
	return o
}

// SetModuleFilter -
func (o tBuild) SetModuleFilter(filter []string) tBuild {
	o.logger.moduleFilter = filter
	//sort.Strings(o.logger.moduleFilter)
	return o
}

// SetWriter -
func (o tBuild) SetWriter(writer io.Writer) tBuild {
	o.logger.writer = writer
	return o
}

// SetStyler -
func (o tBuild) SetStyler(styler TStyler) tBuild {
	o.logger.styler = styler
	return o
}

// SetFormat -
func (o tBuild) SetFormat(format string) tBuild {
	o.logger.format = format
	return o
}
