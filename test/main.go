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

var (
	quit                   = false
	currentWidget, w1, w11 *widget.TWidget
	err                    error
	event                  events.IEvent
)

func randString(n int) string {
	str := ""
	for i := 0; i < n; i++ {
		str += "1"
	}
	return str
}

func initialize() {
	events.NewAction("quit", "q", "", func(ev events.IEvent) bool {
		log.Debug("quit")
		quit = true
		return true
	})
	events.NewAction("left", "a", "", func(ev events.IEvent) bool {
		currentWidget.AddPos(-5, 0)
		currentWidget.Screen().PostUpdate()
		return true
	})
	events.NewAction("right", "d", "", func(ev events.IEvent) bool {
		currentWidget.AddPos(5, 0)
		currentWidget.Screen().PostUpdate()
		return true
	})
	events.NewAction("up", "w", "", func(ev events.IEvent) bool {
		currentWidget.AddPos(0, -5)
		currentWidget.Screen().PostUpdate()
		return true
	})
	events.NewAction("down", "s", "", func(ev events.IEvent) bool {
		currentWidget.AddPos(0, 5)
		currentWidget.Screen().PostUpdate()
		return true
	})
	events.NewAction("1", "1", "", func(ev events.IEvent) bool {
		currentWidget = w1
		return true
	})
	events.NewAction("2", "2", "", func(ev events.IEvent) bool {
		currentWidget = w11
		return true
	})
	events.NewAction("window close", "window close", "", func(ev events.IEvent) bool {
		quit = true
		return true
	})
	events.NewAction("mouse motion", "mouse motion", "", func(ev events.IEvent) bool {
		event := events.ToMouseMotion(ev)
		x := event.X()
		y := event.Y()
		// xScr, yScr := ev.Screen().Size()
		mouseOver, _, _ := ev.Screen().FindWidget(x, y)
		log.Debug("widget: ", mouseOver)
		return true
	})
	events.NewAction("remove", "b", "", func(ev events.IEvent) bool {
		scr := w1.Screen()
		if scr != nil {
			scr.Remove(w1)
			scr.PostUpdate()
		}
		return true
	})
	events.NewAction("add", "n", "", func(ev events.IEvent) bool {
		scr := ev.Screen()
		if scr != nil {
			scr.AddChild(w1)
			scr.PostUpdate()
		}
		return true
	})
	events.NewAction("window resized", "window resized", "", func(ev events.IEvent) bool {
		event := events.ToWindowResized(ev)
		// w, h := w1.Bounds
		dw, dh := event.Delta()
		// dw, dh := event.Delta(event.Screen())
		log.Debug(dw, " ", dh)
		// w1.SetSize(w1.Bounds().W+dw, w1.Bounds().H+dh)
		currentWidget.Screen().PostUpdate()
		return true
	})
	events.NewAction("mouse", "m1", "", func(ev events.IEvent) bool {
		// event := events.ToMouse(ev)
		// if event.Button==1 && event.Type=={

		// }
		return true
	})

	events.ActionMap.Apply()
	log.Debug(events.ActionMap)
}

func main() {

	a := 1820
	log.Add(zlogger.Build().Styler(zlogger.AnsiStyler).Done())
	hal, err := sdlhal.New()
	log.Error(err, "New hal")
	scr, err := hal.NewScreen()
	currentScreen := hal.Screen(1)
	log.Debug(currentScreen)
	w111 := widget.NewWidget().
		SetColor(50, 0, 0, 255).
		SetName(randString(a)).
		SetPos(20, 10).
		SetSize(100, 100)
	w11 = widget.NewWidget().
		SetName("Inner Widget").
		SetColor(0, 0, 50, 255).
		SetPos(40, 40).
		SetSize(150, 150)
	w11.AddChild(w111)

	w1 = widget.NewWidget().
		SetName("Second").
		SetColor(0, 50, 0, 255).
		SetPos(40, 40).
		SetSize(200, 200)
	w1.AddChild(w11)

	w2 := widget.NewWidget().
		SetName("12fasdfasfd").
		SetColor(0, 0, 100, 255).
		SetSize(100, 100).
		SetPos(300, 200)
	scr.AddChild(w1, w2)
	currentWidget = w1
	initialize()
	scr.Remove(w1)
	scr.Remove(w1)
	scr.AddChild(w1)
	scr.AddChild(w1)
	// scr.Remove(w1)
	// scr.Remove(w1)
	event := events.IEvent(nil)
	for !quit {
		hal.Draw()
		// hal.Output().Flush()
		for {
			event = hal.Event()
			if event != nil {
				break
			}
			time.Sleep(1)
		}
		//log.Debug(event)
		events.HandleEvent(event)
	} // for !quit
	hal.Close()
}
