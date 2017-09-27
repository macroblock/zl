package main

import (
	"fmt"
	"os"

	"github.com/macroblock/zl/core/loglevel"
	"github.com/macroblock/zl/core/zlog"
	"github.com/macroblock/zl/core/zlogger"
)

var log = zlog.Get()

func main() {
	log.Add(
		zlogger.Build().
			SetName("test").
			SetFilter(loglevel.All).
			SetWriter(os.Stdout).
			SetStyler(zlogger.AnsiStyler).
			Done(),
		zlogger.Build().
			SetName("short").
			SetFilter(loglevel.Notice.Only().Include(loglevel.Info.Only())).
			SetWriter(os.Stdout).
			SetFormat("~x\n").
			Done())

	log.Debug("debug")
	log.Info("info")
	log.Notice("notice")
	log.Warning(fmt.Errorf("test warning Error"), "warning")
	log.Error(fmt.Errorf("test error Error"), "error")
	log.Recover("recover error state")
	log.Info("without error")

	log.Info("test")
	log.Info(loglevel.Notice.Below())
	log.Info(loglevel.Notice.OrLower().Exclude(loglevel.Error.OrLower()))
	log2 := log.Clone("other")
	log2.Info("other log")
	log2.Info("other prefix")
	log.Info("main log")
	log2.Add(
		zlogger.Build().
			SetName("test2").
			Done())

	log2.Info("other dup message")
	log.Info("main dup msg")

}
