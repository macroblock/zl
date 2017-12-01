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
	out.AddChild(
		widget.NewWidget().
			SetColor(255, 0, 0, 255).
			SetPos(50, 50),
		widget.NewWidget().
			SetColor(0, 255, 0, 255).
			SetPos(100, 50),
		widget.NewWidget().
			SetColor(0, 0, 255, 255).
			SetPos(150, 50),
		widget.NewTextWidget().
			SetText("text012345678").
			SetColor(255, 0, 255, 255).
			SetPos(200, 50),
		widget.NewWidget().
			SetColor(100, 0, 0, 255).
			SetPos(50, 100).
			SetSize(200, 300).
			AddChild(
				widget.NewTextWidget().
					SetText("one0123456789").
					SetColor(0, 100, 0, 255).
					SetPos(-20, 5),
				widget.NewTextWidget().
					SetText("two685410651065").
					SetColor(0, 0, 100, 255).
					SetPos(55, 5),
				widget.NewWidget().
					SetSize(100, 100).
					SetPos(-20, 150).
					AddChild(
						widget.NewTextWidget().
							SetText("three012345678").
							SetColor(0, 0, 100, 255).
							SetPos(5, 5),
						widget.NewTextWidget().
							SetText("four68406540").
							SetColor(0, 0, 100, 255).
							SetPos(55, 55).
							AddChild(
								widget.NewTextWidget().
									SetColor(50, 0, 50, 255).
									SetSize(50, 50).
									SetPos(5, 5))),
				widget.NewTextWidget().
					SetText("five").
					SetPos(50, 50).
					SetSize(70, 50)),
		widget.NewTextWidget().
			SetText("12345678").
			SetColor(0, 0, 100, 255).
			SetPos(300, 50),
		widget.NewWidget().
			SetColor(0, 0, 255, 255).
			SetPos(300, 200))
	event := events.IEvent(nil)
	quit := false
	for !quit {
		x.Draw()
		// hal.Output().Flush()
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

	x.Close()

}
