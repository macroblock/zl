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
	tw := widget.NewTextWidget().
		SetColor(50, 0, 0, 255).
		SetText("test12345678456411").
		SetPos(-15, -5).
		SetSize(150, 150)
	out.AddChild(
		// widget.NewWidget().
		// 	SetColor(255, 0, 0, 255).
		// 	SetPos(50, 50),
		// widget.NewWidget().
		// 	SetColor(0, 255, 0, 255).
		// 	SetPos(100, 50),
		// widget.NewWidget().
		// 	SetColor(0, 0, 255, 255).
		// 	SetPos(150, 50),
		// widget.NewTextWidget().
		// 	SetText("text012345678").
		// 	SetColor(255, 0, 255, 255).
		// 	SetPos(200, 50),
		// widget.NewWidget().
		// 	SetColor(100, 0, 0, 255).
		// 	SetPos(50, 100).
		// 	SetSize(200, 300).
		// 	AddChild(
		// 		widget.NewTextWidget().
		// 			SetText("one0123456789").
		// 			SetColor(0, 100, 0, 255).
		// 			SetPos(-20, 5),
		// 		widget.NewTextWidget().
		// 			SetText("two685410651065").
		// 			SetColor(0, 0, 100, 255).
		// 			SetPos(55, 5),
		// 		widget.NewWidget().
		// 			SetSize(100, 100).
		// 			SetPos(-20, 150).
		// 			AddChild(
		// 				widget.NewTextWidget().
		// 					SetText("three012345678").
		// 					SetColor(0, 0, 100, 255).
		// 					SetPos(5, 5),
		// 				widget.NewTextWidget().
		// 					SetText("four68406540").
		// 					SetColor(0, 0, 100, 255).
		// 					SetPos(55, 55)),
		// 		widget.NewTextWidget().
		// 			SetText("five").
		// 			SetPos(50, 50).
		// 			SetSize(70, 50)),
		widget.NewWidget().
			SetColor(0, 0, 100, 255).
			SetSize(100, 100).
			SetPos(300, 50).
			AddChild(widget.NewTextWidget().SetText("12345678").SetColor(0, 200, 100, 255).SetSize(100, 100).SetPos(270, 50).
				AddChild(widget.NewTextWidget().SetText("12345678").SetColor(200, 0, 100, 255).SetSize(100, 100).SetPos(250, 50))),
		widget.NewWidget().
			SetColor(0, 0, 255, 255).
			SetPos(20, 20).
			SetSize(100, 100).
			AddChild(tw))

	event := events.IEvent(nil)
	quit := false
	hal.Output().SetViewport(nil)
	log.Debug("viewport ", hal.Output().GetViewport())
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
			if ev.Rune() == 'a' {
				tw.AddPos(-1, 0)
			}
			if ev.Rune() == 'w' {
				tw.AddPos(0, -1)
			}
			if ev.Rune() == 'd' {
				tw.AddPos(1, 0)
			}
			if ev.Rune() == 's' {
				tw.AddPos(0, 1)
			}
		case *events.TDropFileEvent:
			log.Info(ev.Content())
		}

	}

	x.Close()

}
