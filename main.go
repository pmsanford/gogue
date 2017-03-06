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
		game.draw()
		ev := termbox.PollEvent()
		if ev.Ch == 'q' {
			break
		}
		if ev.Ch == 'k' {
			game.player.loc.y -= 1
		}
		if ev.Ch == 'j' {
			game.player.loc.y += 1
		}
		if ev.Ch == 'h' {
			game.player.loc.x -= 1
		}
		if ev.Ch == 'l' {
			game.player.loc.x += 1
		}
	}
}
