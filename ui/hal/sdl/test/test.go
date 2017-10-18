package main

import (
	"fmt"
	"time"

	"github.com/macroblock/zl/core/zlog"
	"github.com/macroblock/zl/core/zlogger"
	"github.com/macroblock/zl/ui/event"
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

	quit := false
	for !quit {
		ev := x.Event()
		if ev == nil {
			time.Sleep(1)
		}
		switch t := ev.(type) {
		case *event.TKeyboard:
			fmt.Println(t.String())
			if t.Rune() == 'q' {
				quit = true
			}
		}
	}
	x.Close()
	x.Close()
}
