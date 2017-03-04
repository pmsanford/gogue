package main

import (
	//"fmt"
	"github.com/nsf/termbox-go"
)

func print_tb(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	print_tb(0, 0, termbox.ColorWhite, termbox.ColorBlack, "Hello World")

	termbox.Flush()

	for {
		ev := termbox.PollEvent();
		if ev.Key == termbox.KeyCtrlQ {
			break
		}
	}
}