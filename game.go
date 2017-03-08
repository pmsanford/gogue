package main

import (
	"time"
)

type Game struct {
	term          Terminal
	active_state  game_state
	display_state game_state
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
	input := GetInput()
	return Game{term: GetTerminal(), active_state: game_state{m: CreateMap(), player: Player{hp: 100, loc: Location{x: 5, y: 5}, char: '@', color: Red, input: input}, running: true}}
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
		if game.active_state.player.act() == Quit {
			break
		}
	}
}
