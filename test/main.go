package main

import (
	"time"

	"github.com/macroblock/zl/core/zlog"
	"github.com/macroblock/zl/core/zlogger"
	"github.com/macroblock/zl/ui/events"
	"github.com/macroblock/zl/ui/hal/sdl"
	"github.com/macroblock/zl/ui/widget"
)

var log = zlog.Instance("main")

func main() {
	log.Add(zlogger.Build().Styler(zlogger.AnsiStyler).Done())
	x, err := hal.New()
	log.Error(err, "New hal")
	out, err := x.NewOutput()
	w, err := widget.NewWidget()
	w.SetColor(140, 0, 0, 255)
	w.SetPos(50, 50)
	out.AddChild(w)

	w, err = widget.NewWidget()
	w.SetColor(0, 255, 0, 255)
	w.SetPos(100, 100)
	out.AddChild(w)
	event := events.IEvent(nil)
	quit := false
	for !quit {
		out.SetFillColor(0, 0, 0, 0)
		out.Clear()

		out.Draw()
		out.DrawText("text")
		out.Flush()
		for {
			event = x.Event()
			if event != nil {
				break
			}
			time.Sleep(1)
		}
		switch ev := event.(type) {
		case *events.TKeyboardEvent:
			if ev.Rune() == 'q' {
				quit = true
			}
		case *events.TDropFileEvent:
			log.Info(ev.Content())
		}

	}
	out.Close()
	x.Close()
}
