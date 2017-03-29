package main

import (
	"time"
)

type Game struct {
	term          Terminal
	active_state  game_state
	display_state game_state
	generator     MonsterGenerator
}

type LocationType uint16

const (
	Empty LocationType = iota
	Blocked
)

type game_state struct {
	m        imap
	player   Player
	running  bool
	monsters []Entity
}

func (state game_state) copy_to(other game_state) {
	other.m = state.m
	other.player = state.player
	other.running = state.running
}

func new_game() Game {
	input := GetInput()
	ng := Game{
		term: GetTerminal(),
		active_state: game_state{
			m: CreateMap(),
			player: Player{
				hp:    100,
				loc:   Location{x: 5, y: 5},
				char:  '@',
				color: Red,
				input: input},
			running:  true,
			monsters: make([]Entity, 0)},
		generator: MonsterGenerator{
			monster_table: map[int]MonsterType{100: Goblin}}}
	return ng
}

func (g *Game) init() {
	g.term.init()
	mon := g.generator.new_monster()
	g.active_state.monsters = append(g.active_state.monsters, &mon)
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
	for _, m := range g.display_state.monsters {
		m.draw(g.term)
	}
	g.term.flush()
}

func (g *Game) blit() {
	g.display_state = g.active_state
}

func (game *Game) loop() {
Quit:
	for {
		game.blit()
		for {
			result := game.active_state.player.act(&game.active_state.m)
			if result == Quit {
				break Quit
			}
			if result == Acted {
				break
			}
		}
		for _, m := range game.active_state.monsters {
			m.act(&game.active_state.m)
		}
	}
}
