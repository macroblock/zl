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
	throw     = log.Catcher()
	logFilter = loglevel.Warning.OrLower()
)

func doProcess(path string) {
	defer throw.Catch()
	log.Info("")
	log.Info("rename: " + path)
	dir, name := filepath.Split(path)
	ext := ""

	file, err := os.Open(path)
	throw.Error(err, "can not open file: ", path)

	stat, err := file.Stat()
	throw.Error(err, "can not get filestat: ", path)

	err = file.Close()
	throw.Error(err, "can not close file: ", path)

	if !stat.IsDir() {
		ext = filepath.Ext(path)
	}
	name = strings.TrimSuffix(name, ext)
	name, _ = text.Translit(name)
	err = os.Rename(path, dir+name+ext)
	throw.Error(err, "can not rename file")

	log.Notice("result: " + dir + name + ext)
}

func main() {
	log.Add(
		zlogger.Build().
			LevelFilter(logFilter).
			Styler(zlogger.AnsiStyler).
			Done(),
		zlogger.Build().
			LevelFilter(loglevel.Info.Only().Include(loglevel.Notice.Only())).
			Format("~x\n").
			Styler(zlogger.AnsiStyler).
			Done())

	defer func() {
		if log.State().Intersect(loglevel.Warning.OrLower()) != 0 {
			cmd := exec.Command("cmd", "/C", "pause")
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Run()
		}
	}()

	log.Debug("log initialized")
	if len(os.Args) <= 1 {
		log.Warning(true, "not enough parameters")
	}
	for _, path := range os.Args[1:] {
		doProcess(path)
	}
}
