package main

import (
	//"fmt"
	"github.com/nsf/termbox-go"
)

func main() {
	game := Game{term: GetTerminal(), m: CreateMap(), player: Player{hp: 100, loc: Location{x: 5, y: 5}, char: '@'}}
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
