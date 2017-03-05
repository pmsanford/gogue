package main

import (
	//"fmt"
	"github.com/nsf/termbox-go"
)

func main() {
	g := Game {term: GetTerminal(), m: CreateMap()}
	g.init()
	defer g.end()

	g.draw()

	for {
		ev := termbox.PollEvent();
		if ev.Key == termbox.KeyCtrlQ {
			break
		}
	}
}