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
var currentWidget, w1, w2 *widget.TWidget

func randString(n int) string {
	str := ""
	for i := 0; i < n; i++ {
		str += "1"
	}
	return str
}

func initialize() {
	events.NewKeyboardAction("quit", "q", "", func(ev events.TKeyboardEvent) bool {
		log.Debug("quit")
		quit = true
		return true
	})
	events.NewKeyboardAction("left", "a", "", func(ev events.TKeyboardEvent) bool {
		currentWidget.AddPos(-5, 0)
		ev.Screen().PostUpdate()
		return true
	})
	events.NewKeyboardAction("right", "d", "", func(ev events.TKeyboardEvent) bool {
		currentWidget.AddPos(5, 0)
		ev.Screen().PostUpdate()
		return true
	})
	events.NewKeyboardAction("up", "w", "", func(ev events.TKeyboardEvent) bool {
		currentWidget.AddPos(0, -5)
		ev.Screen().PostUpdate()
		return true
	})
	events.NewKeyboardAction("down", "s", "", func(ev events.TKeyboardEvent) bool {
		currentWidget.AddPos(0, 5)
		ev.Screen().PostUpdate()
		return true
	})
	events.NewKeyboardAction("1", "1", "", func(ev events.TKeyboardEvent) bool {
		currentWidget = w1
		return true
	})
	events.NewKeyboardAction("2", "2", "", func(ev events.TKeyboardEvent) bool {
		currentWidget = w2
		return true
	})
	events.NewKeyboardAction("3", "3", "", func(ev events.TKeyboardEvent) bool {

		return true
	})
	events.ActionMap.Apply()
}

func main() {

	a := 1820
	log.Add(zlogger.Build().Styler(zlogger.AnsiStyler).Done())
	x, err := sdlhal.New()
	log.Error(err, "New hal")
	out, err := x.NewScreen()
	w1 = widget.NewWidget().
		SetColor(50, 0, 0, 255).
		SetName(randString(a)).
		SetPos(20, 10).
		SetSize(100, 100)
	w2 = widget.NewWidget().
		SetName("Inner Widget").
		SetColor(0, 0, 50, 255).
		SetPos(40, 40).
		SetSize(150, 150).
		AddChild(w1)
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
			AddChild(w2))
	currentWidget = w1
	initialize()
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
		events.HandleEvent(event)
		// switch ev := event.(type) {
		// case *events.TKeyboardEvent:
		// 	//log.Info(ev.Rune())
		// 	scr := ev.Screen()
		// 	if scr == nil {
		// 		break
		// 	}
		// 	switch ev.Rune() {
		// 	case 'q':
		// 		quit = true
		// 	case 'a':
		// tw.AddPos(-5, 0)
		// 		scr.PostUpdate()
		// 	case 'w':
		// 		tw.AddPos(0, -5)
		// 		scr.PostUpdate()
		// 	case 'd':
		// 		tw.AddPos(5, 0)
		// 		scr.PostUpdate()
		// 	case 's':
		// 		tw.AddPos(0, 5)
		// 		scr.PostUpdate()
		// 	case 'h':
		// 		iw.AddPos(-5, 0)
		// 		scr.PostUpdate()
		// 	case 'k':
		// 		iw.AddPos(0, -5)
		// 		scr.PostUpdate()
		// 	case 'l':
		// 		iw.AddPos(5, 0)
		// 		scr.PostUpdate()
		// 	case 'j':
		// 		iw.AddPos(0, 5)
		// 		scr.PostUpdate()
		// 	}
		// case *events.TWindowCloseEvent:
		// 	quit = true
		// case *events.TDropFileEvent:
		// 	log.Info(ev.Content())
		// } // switch
	} // for !quit
	x.Close()
}
