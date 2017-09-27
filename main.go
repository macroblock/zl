package main

import (
	"os"

	"github.com/macroblock/zl/core/loglevel"
	"github.com/macroblock/zl/core/zlog"
)

var log = zlog.Get()

func main() {
	log.AddLogger(loglevel.Info.Only(), nil, os.Stdout)
	log.Info("test")
	log.Info(loglevel.Notice.Below())
	log.Info(loglevel.Notice.OrLower().Exclude(loglevel.Error.OrLower()))
	log2 := log.Clone("other")
	log2.Info("other log")
	log2.Info("other prefix")
	log.Info("main log")
	log2.AddLogger(loglevel.Info.Only(), nil, os.Stdout)
	log2.Info("other dup message")
	log.Info("main dup msg")
}
