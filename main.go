package main

import (
	"os"

	"github.com/macroblock/zl/core/log"
)

func main() {
	l := log.Default()
	l.AddLogger(log.Info.Filter(), nil, os.Stdout)
	l.Info("test")
	l.Info(log.Notice.Below())
}
