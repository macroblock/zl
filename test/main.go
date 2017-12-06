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

func randString(n int) string {
	str := ""
	for i := 0; i < n; i++ {
		str += "1"
	}
	return str
}

func main() {

	a := 1820
	log.Add(zlogger.Build().Styler(zlogger.AnsiStyler).Done())
	x, err := hal.New()
	log.Error(err, "New hal")
	out, err := x.NewScreen()
	tw := widget.NewTextWidget().
		SetColor(50, 0, 0, 255).
		SetText(randString(a)).
		SetPos(20, 10).
		SetSize(100, 100)
	iw := widget.NewWidget().
		SetName("Inner Widget").
		SetColor(0, 0, 50, 255).
		SetPos(40, 40).
		SetSize(150, 150).
		AddChild(tw)
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
			SetName("12fasdfasfd").
			SetColor(0, 0, 100, 255).
			SetSize(100, 100).
			SetPos(300, 200),
		// AddChild(
		// widget.NewTextWidget().
		// 	SetText("12345678").
		// 	SetColor(0, 200, 100, 255).
		// 	SetSize(100, 100).
		// 	SetPos(-80, 10).
		// 	AddChild(
		// 		widget.NewTextWidget().
		// 			SetText("12345678").
		// 			SetColor(200, 0, 100, 255).
		// 			SetSize(100, 100).
		// 			SetPos(-80, 0))),
		widget.NewWidget().
			SetName("Second").
			SetColor(0, 50, 0, 255).
			SetPos(40, 40).
			SetSize(200, 200).
			AddChild(iw))

	event := events.IEvent(nil)
	events.InitActionMap()

	events.NewKeyboardAction("refresh", "i", "", func(ev events.TKeyboardEvent) bool {
		log.Debug("refresh")
		x.Draw()
		return true
	})
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
		log.Debug(event)
		switch ev := event.(type) {
		case *events.TKeyboardEvent:
			//log.Info(ev.Rune())
			scr := ev.Screen()
			if scr == nil {
				break
			}
			switch ev.Rune() {
			case 'q':
				quit = true
			case 'a':
				tw.AddPos(-5, 0)
				scr.PostUpdate()
			case 'w':
				tw.AddPos(0, -5)
				scr.PostUpdate()
			case 'd':
				tw.AddPos(5, 0)
				scr.PostUpdate()
			case 's':
				tw.AddPos(0, 5)
				scr.PostUpdate()
			case 'h':
				iw.AddPos(-5, 0)
				scr.PostUpdate()
			case 'k':
				iw.AddPos(0, -5)
				scr.PostUpdate()
			case 'l':
				iw.AddPos(5, 0)
				scr.PostUpdate()
			case 'j':
				iw.AddPos(0, 5)
				scr.PostUpdate()
			}
		case *events.TWindowCloseEvent:
			quit = true
		case *events.TDropFileEvent:
			log.Info(ev.Content())
		}
	}
	x.Close()
}
