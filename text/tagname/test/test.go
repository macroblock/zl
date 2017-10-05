package main

import (
	"fmt"
	"os"

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
			// LevelFilter(loglevel.Error.OrLower()).
			Format("(~m) ~l: ~x~e\n").
			Done())
	log.Info("programmstart")
	log.Error(nil, "programmstart")
	log.Warning(nil, "egega';,")
	tagname.Something()

	result := tagname.Parse("_ae2r6__q0w0_q1s3__q3w2_q2s2_q1w1_")
	for _, tag := range result {
		log.Info(fmt.Sprintf("%v: %v", tag.Type, tag.Value))
	}
}
