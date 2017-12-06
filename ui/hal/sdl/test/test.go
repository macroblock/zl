package main

import (
	"fmt"
	"reflect"
	"runtime"
	"time"

	"github.com/veandco/go-sdl2/sdl"

	"net/http"
	_ "net/http/pprof"

	"github.com/macroblock/zl/core/zlog"
	"github.com/macroblock/zl/core/zlogger"
	"github.com/macroblock/zl/ui/events"
	"github.com/macroblock/zl/ui/hal/sdl"
)

var log = zlog.Instance("main")

var mem runtime.MemStats

func main() {
	go http.ListenAndServe(":8080", nil)
	log.Add(
		zlogger.Build().
			Styler(zlogger.AnsiStyler).
			Done())

	x, _ := hal.New()
	v := sdl.Version{}
	sdl.GetVersion(&v)
	log.Info("version: ", v)
	output, _ := x.NewScreen()
	// x.NewOutput()
	// x.NewOutput()
	// x.NewOutput()

	quit := false
	ev := events.IEvent(nil)
	for !quit {
		output.Draw()
		output.SetFillColor(100, 50, 25, 0)
		output.SetDrawColor(100, 100, 100, 0)
		output.Clear()
		output.SetFillColor(50, 25, 50, 0)
		output.FillRect(20, 20, 200, 50)
		output.DrawRect(20, 20, 200, 50)
		output.Flush()

		for ev == nil {
			time.Sleep(1)
			ev = x.Event()
		}

		switch t := ev.(type) {
		case *events.TKeyboardEvent:
			fmt.Println(t.String())
			if t.Rune() == 'q' {
				quit = true
			}
		case *events.TDropFileEvent:
			fmt.Println(t.String())
		}
	}
	//output.Close()
	output.Close()

	x.Close()
	//str := 24
	ptr := output
	fmt.Printf("%T: %v\n", ptr, ptr)
	setNil(&ptr)
	fmt.Printf("%T: %v\n", ptr, ptr)
	// println(math.Log10(255))
	// println(math.Log10(100))
	// println(int(math.Log10(math.MaxInt64)))
	// maxInt := -1
	// maxInt = int(uint(maxInt) >> 1)
	// println(int(math.Log10(float64(maxInt))))
}

func setNil(i interface{}) {
	t := reflect.TypeOf(i)
	if t.Kind() != reflect.Ptr {
		println("is not pointer")
		return
	}
	t = t.Elem()
	if t.Kind() != reflect.Ptr {
		println("is not pointer to pointer")
		return
	}
	t = t.Elem()
	// v := reflect.ValueOf(i).Elem()
	// v.Set(reflect.Zero(v.Type()))
}
