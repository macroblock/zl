package main

import (
	"fmt"
	"runtime"
	"time"

	"net/http"
	_ "net/http/pprof"

	"github.com/macroblock/zl/core/zlog"
	"github.com/macroblock/zl/core/zlogger"
	"github.com/macroblock/zl/ui/events"
	"github.com/macroblock/zl/ui/hal/sdl"
)

var log = zlog.Instance("main")

var mem runtime.MemStats

// func printStat() {
// fmt.Printf("HeapAlloc: %v Loockups: %v MallocsUsing: %v\n", mem.HeapAlloc, mem.Lookups, mem.Mallocs-mem.Frees)
// }

func main() {
	go http.ListenAndServe(":8080", nil)
	log.Add(
		zlogger.Build().
			Styler(zlogger.AnsiStyler).
			Done())

	x, _ := hal.New()
	x, _ = hal.New()

	quit := false
	// runtime.ReadMemStats(&mem)
	// printStat()
	for !quit {
		ev := x.Event()
		if ev == nil {
			time.Sleep(1)
			continue
		}
		runtime.ReadMemStats(&mem)
		switch t := ev.(type) {
		case *events.TKeyboardEvent:
			fmt.Println(t.String())
			if t.Rune() == 'q' {
				quit = true
			}
		case *events.TDropFileEvent:
			fmt.Println(t.String())
		}
		// printStat()
	}

	x.Close()
	x.Close()
}
