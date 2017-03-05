package main

import (
	//"fmt"
	"github.com/nsf/termbox-go"
)

func main() {
	game := Game{term: GetTerminal(), m: CreateMap()}
	game.init()
	defer game.end()

	game.draw()

	for {
		ev := termbox.PollEvent()
		if ev.Key == termbox.KeyCtrlQ {
			break
		}
	}
}
