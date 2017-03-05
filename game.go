package main

type Game struct {
	term Terminal
	m Map
}

func (g Game) init() {
	g.term.init()
}

func (g Game) end() {
	g.term.close()
}

func (g Game) draw() {
	g.m.draw(g.term)
	g.term.flush()
}