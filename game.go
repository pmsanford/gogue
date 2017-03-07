package main

type Game struct {
	term   Terminal
	m      Map
	player Player
	input  Input
}

func (g Game) init() {
	g.term.init()
}

func (g Game) end() {
	g.term.close()
}

func (g Game) draw() {
	g.m.draw(g.term)
	g.player.draw(g.term)
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
			game.player.loc.y -= 1
		}
		if ev == 'j' {
			game.player.loc.y += 1
		}
		if ev == 'h' {
			game.player.loc.x -= 1
		}
		if ev == 'l' {
			game.player.loc.x += 1
		}
	}
}
