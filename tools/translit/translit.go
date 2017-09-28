package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/macroblock/zl/core/loglevel"
	"github.com/macroblock/zl/core/zlog"
	"github.com/macroblock/zl/core/zlogger"
	"github.com/macroblock/zl/text"
)

var (
	log       = zlog.Instance("main")
	logFilter = loglevel.Warning.OrLower() //loglevel.All
)

func main() {
	log.Add(
		zlogger.Build().
			SetName("sdtout").
			SetFilter(logFilter).
			SetWriter(os.Stdout).
			Done())

	log.Debug("log initialized")
	if len(os.Args) <= 1 {
		log.Warning(nil, "not enough parameters")
	}
	for _, filename := range os.Args[1:] {
		dir, name := filepath.Split(filename)
		ext := filepath.Ext(filename)
		name = strings.TrimSuffix(name, ext)
		name, _ = text.Translit(name)
		err := os.Rename(filename, dir+name+ext)
		if err != nil {
			log.Error(err, "cannot rename file")
		}
	}
}
