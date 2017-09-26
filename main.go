package main

import (
	"os"

	"github.com/macroblock/zl/core/log"
	"github.com/macroblock/zl/core/loglevel"
)

func main() {
	l := log.Default()
	l.AddLogger(loglevel.Info.Only(), nil, os.Stdout)
	l.Info("test")
	l.Info(loglevel.Notice.Below())
	l.Info(loglevel.Notice.OrLower().Exclude(loglevel.Error.OrLower()))
	b := l.Clone()
	b.Info("other log")
	b.SetPrefix("other")
	b.Info("other prefix")
	l.Info("main log")
	b.AddLogger(loglevel.Info.Only(), nil, os.Stdout)
	b.Info("other dup message")
	l.Info("main dup msg")
}
