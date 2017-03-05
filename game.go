package main

type Game struct {
	term   Terminal
	m      Map
	player Player
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
