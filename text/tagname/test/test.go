package main

import (
	"os"

	"github.com/macroblock/zl/core/loglevel"
	"github.com/macroblock/zl/core/zlog"
	"github.com/macroblock/zl/core/zlogger"
	"github.com/macroblock/zl/text/tagname"
)

var log = zlog.Instance("main")

var logflag = loglevel.Info.OrLower()

func main() {
	log.Add(
		zlogger.Build().
			Writer(os.Stdout).
			Styler(zlogger.AnsiStyler).
			LevelFilter(logflag).
			Format("(~m) ~l: ~x~e\n").
			Done())
	log.Info("programmstart")
	log.Error(nil, "programmstart")
	log.Warning(nil, "egega';,")
	tagname.Something()

	result := tagname.Parse("_hd_1994_ae2r6__q0w0_q1s3__q3w2_q2s2_q1w1_masdlkfjasd_q2w1")
	for _, tag := range result {
		log.Info(tag.Type, ": \"", tag.Value, "\"")
	}
}
