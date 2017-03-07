package main

type Game struct {
	term          Terminal
	active_state  game_state
	display_state game_state
	input         Input
}

type game_state struct {
	m      Map
	player Player
}

func new_game() Game {
	return Game{term: GetTerminal(), active_state: game_state{m: CreateMap(), player: Player{hp: 100, loc: Location{x: 5, y: 5}, char: '@', color: Red}}, input: GetInput()}
}

func (g Game) init() {
	g.display_state = g.active_state
	g.term.init()
}

func (g Game) end() {
	g.term.close()
}

func (g Game) draw() {
	g.display_state = g.active_state
	g.display_state.m.draw(g.term)
	g.display_state.player.draw(g.term)
	g.term.flush()
}

func (game Game) loop() {
	for {
		game.draw()
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
