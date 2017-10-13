package main

import (
	"github.com/macroblock/zl/core/zlog"
	"github.com/macroblock/zl/core/zlogger"
	"github.com/macroblock/zl/ui/hal/sdl"
)

var log = zlog.Instance("main")

func main() {
	log.Add(
		zlogger.Build().
			Styler(zlogger.AnsiStyler).
			Done())

	x, _ := hal.New()
	x, _ = hal.New()
	x.Close()
	x.Close()
}
