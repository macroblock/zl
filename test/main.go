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
	out2, err := x.NewOutput()

	w := widget.NewWidget()
	w.SetColor(255, 0, 0, 255)
	w.SetPos(50, 50)
	out.AddChild(w)

	w = widget.NewWidget()
	w.SetColor(0, 255, 0, 255)
	w.SetPos(100, 50)
	out.AddChild(w)

	w = widget.NewWidget()
	w.SetColor(0, 0, 255, 255)
	w.SetPos(150, 50)
	out.AddChild(w)

	t := widget.NewTextWidget()
	t.SetText("text")
	t.SetColor(255, 0, 255, 255)
	t.SetPos(200, 50)
	out.AddChild(t)
	event := events.IEvent(nil)
	quit := false
	for !quit {
		x.Draw()
		hal.Output().Flush()
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
	out2.Close()

	x.Close()
}
