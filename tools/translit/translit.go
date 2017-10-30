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
		log.Info("")
		log.Info("rename: " + path)
		dir, name := filepath.Split(path)
		ext := ""

		file, err := os.Open(path)
		if err != nil {
			log.Error(err, "connot open file: ", path)
			continue
		}
		stat, err := file.Stat()
		if err != nil {
			log.Error(err, "connot get filestat: ", path)
			continue
		}
		if !stat.IsDir() {
			ext = filepath.Ext(path)
		}
		name = strings.TrimSuffix(name, ext)
		name, _ = text.Translit(name)
		err = os.Rename(path, dir+name+ext)
		if err != nil {
			log.Error(err, "cannot rename file")
		} else {
			log.Notice("result: " + dir + name + ext)
		}
	}
}
