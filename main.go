package main

import (
	"fmt"
	"os"

	"github.com/macroblock/zl/core/loglevel"
	"github.com/macroblock/zl/core/zlog"
	"github.com/macroblock/zl/core/zlogger"
)

var log = zlog.Instance("main")

func main() {
	log.Add(
		zlogger.Build().
			SetName("test").
			SetLevelFilter(loglevel.All).
			SetWriter(os.Stdout).
			SetStyler(zlogger.AnsiStyler).
			Done(),
		zlogger.Build().
			SetName("short").
			SetLevelFilter(loglevel.Notice.Only().Include(loglevel.Info.Only())).
			SetModuleFilter([]string{"other"}).
			SetWriter(os.Stdout).
			SetFormat("---- ~m ---- ~x\n").
			Done(),
	)

	log.Debug("debug")
	log.Info("info")
	log.Notice("notice")
	log.Warning(fmt.Errorf("test Warning error"), "warning")
	log.Error(fmt.Errorf("test Error error"), "error")
	log.Reset(loglevel.Reset.OrLower(), "recover error state")
	log.Info("without error")

	log.Info("test")
	log.Info(loglevel.Notice.Below())
	log.Info(loglevel.Notice.OrLower().Exclude(loglevel.Error.OrLower()))
	log2 := log.Instance("other")
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
