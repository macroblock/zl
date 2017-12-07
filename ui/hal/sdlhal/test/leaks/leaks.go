package main

import (
	"errors"
	_ "net/http/pprof"
)

type test struct{}

type taaa int32
type tbbb int32

const AAAA taaa = 10
const AABB tbbb = 15

func (o *test) foo() (err error) {
	defer enterleave(&err)()

	println(o)
	return
}

func enterleave(err *error) func() {
	//println("enter: ", err.Error())
	return func() {
		*err = errors.New("changed")
		//println("leave: ", err.Error())
	}
}

func main() {
	// str := errors.New("test")
	// defer enterleave(&str)()
	// go http.ListenAndServe(":8080", nil)
	// sdl.Init(sdl.INIT_VIDEO)
	// win, _ := sdl.CreateWindow("leak test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 640, 480, 0)
	// quit := false
	// for !quit {
	// 	ev := sdl.PollEvent()
	// 	switch t := ev.(type) {
	// 	case *sdl.QuitEvent:
	// 		quit = true
	// 	case *sdl.DropEvent:
	// 		println(t.File)
	// 	}
	// }
	// win.Destroy()
	// sdl.Quit()
	x := &test{}
	err := x.foo()
	println("error: ", err.Error())
	x = nil
	err = x.foo()
	println("error: ", err.Error())

	ttt := AAAA
	_ = ttt
}
