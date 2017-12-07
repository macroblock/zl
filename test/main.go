package main

import (
	"time"

	"github.com/macroblock/zl/core/zlog"
	"github.com/macroblock/zl/core/zlogger"
	"github.com/macroblock/zl/ui/events"
	"github.com/macroblock/zl/ui/hal/sdlhal"
	"github.com/macroblock/zl/ui/widget"
)

var log = zlog.Instance("main")

var quit = false

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
	x, err := sdlhal.New()
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
		widget.NewWidget().
			SetName("12fasdfasfd").
			SetColor(0, 0, 100, 255).
			SetSize(100, 100).
			SetPos(300, 200),
		widget.NewWidget().
			SetName("Second").
			SetColor(0, 50, 0, 255).
			SetPos(40, 40).
			SetSize(200, 200).
			AddChild(iw))

	events.NewKeyboardAction("quit", "q", "", func(ev events.TKeyboardEvent) bool {
		log.Debug("quit")
		quit = true
		return true
	})
	events.NewKeyboardAction("refresh", "i", "", func(ev events.TKeyboardEvent) bool {
		log.Debug("refresh")
		x.Draw()
		return true
	})
	events.ActionMap.Apply()
	event := events.IEvent(nil)
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
		// events.HandleEvent(event)
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
		} // switch
	} // for !quit
	x.Close()
}
