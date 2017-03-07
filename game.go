package main

import (
	"time"
)

type Game struct {
	term          Terminal
	active_state  game_state
	display_state game_state
	input         Input
}

type game_state struct {
	m       Map
	player  Player
	running bool
}

func (state game_state) copy_to(other game_state) {
	other.m = state.m
	other.player = state.player
	other.running = state.running
}

func new_game() Game {
	return Game{term: GetTerminal(), active_state: game_state{m: CreateMap(), player: Player{hp: 100, loc: Location{x: 5, y: 5}, char: '@', color: Red}, running: true}, input: GetInput()}
}

func (g *Game) init() {
	g.term.init()
	g.blit()
	go g.draw_loop()
}

func (g Game) end() {
	g.term.close()
}

func (g *Game) draw_loop() {
	for g.display_state.running {
		g.draw()
		g.display_state.m.next()
		g.display_state.player.next()
		time.Sleep(100 * time.Millisecond)
	}
}

func (g *Game) draw() {
	g.display_state.m.draw(g.term)
	g.display_state.player.draw(g.term)
	g.term.flush()
}

func (g *Game) blit() {
	g.display_state = g.active_state
}

func (game *Game) loop() {
	for {
		game.blit()
		ev := game.input.getchar()
		if ev == 'q' {
			break
		}
		if ev == 'k' {
			game.active_state.player.loc.y -= 1
		}
		if ev == 'j' {
			game.active_state.player.loc.y += 1
		}
		if ev == 'h' {
			game.active_state.player.loc.x -= 1
		}
		if ev == 'l' {
			game.active_state.player.loc.x += 1
		}
		if ev == 'y' {
			game.active_state.player.loc.y -= 1
			game.active_state.player.loc.x -= 1
		}
		if ev == 'u' {
			game.active_state.player.loc.y -= 1
			game.active_state.player.loc.x += 1
		}
		if ev == 'b' {
			game.active_state.player.loc.y += 1
			game.active_state.player.loc.x -= 1
		}
		if ev == 'n' {
			game.active_state.player.loc.y += 1
			game.active_state.player.loc.x += 1
		}
	}
}
