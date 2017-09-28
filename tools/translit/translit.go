package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/macroblock/zl/core/loglevel"
	"github.com/macroblock/zl/core/zlog"
	"github.com/macroblock/zl/core/zlogger"
	"github.com/macroblock/zl/text"
)

var (
	log       = zlog.Instance("main")
	logFilter = loglevel.Warning.OrLower()
)

func main() {
	log.Add(
		zlogger.Build().
			SetFilter(logFilter).
			SetStyler(zlogger.AnsiStyler).
			Done(),
		zlogger.Build().
			SetFilter(loglevel.Info.Only().Include(loglevel.Notice.Only())).
			SetFormat("~x\n").
			SetStyler(zlogger.AnsiStyler).
			Done())

	defer func() {
		if log.HasError() {
			cmd := exec.Command("cmd", "/C", "pause")
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Run()
		}
	}()

	log.Debug("log initialized")
	if len(os.Args) <= 1 {
		log.Warning(nil, "not enough parameters")
	}
	for _, path := range os.Args[1:] {
		log.Info("")
		log.Info("rename: " + path)
		dir, name := filepath.Split(path)
		ext := filepath.Ext(path)
		name = strings.TrimSuffix(name, ext)
		name, _ = text.Translit(name)
		err := os.Rename(path, dir+name+ext)
		if err != nil {
			log.Error(err, "cannot rename file")
		} else {
			log.Notice("result: " + dir + name + ext)
		}
	}
}
