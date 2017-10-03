package main

import (
	"os"

	"github.com/macroblock/zl/core/loglevel"

	"github.com/macroblock/zl/core/zlog"
	"github.com/macroblock/zl/core/zlogger"
	"github.com/macroblock/zl/text/tagname"
)

var log = zlog.Instance("main")

func main() {
	log.Add(
		zlogger.Build().
			Writer(os.Stdout).
			Styler(zlogger.AnsiStyler).
			LevelFilter(loglevel.Error.OrLower()).
			Format("(~m) ~l: ~x~e\n").
			Done())
	log.Info("programmstart")
	log.Error(nil, "programmstart")
	log.Warning(nil, "egega';,")
	tagname.Something()
}
