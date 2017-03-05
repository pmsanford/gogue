package main

type Entity interface {
	draw(term Terminal)
	act()
}

type Location struct {
	x, y int
}

type Player struct {
	hp   int
	loc  Location
	char rune
}

func (p Player) draw(term Terminal) {
	term.draw_char(p.loc.x, p.loc.y, p.char)
}
