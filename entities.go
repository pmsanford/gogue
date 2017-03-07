package main

type Entity interface {
	draw(term Terminal)
	next()
	act()
}

type Location struct {
	x, y int
}

type Player struct {
	hp    int
	loc   Location
	char  rune
	color Color
}

func (p Player) draw(term Terminal) {
	term.draw_char_ex(p.loc.x, p.loc.y, p.char, p.color, DefaultColor)
}

func (p *Player) next() {
	if p.color == Red {
		p.color = White
	} else if p.color == White {
		p.color = Green
	} else if p.color == Green {
		p.color = Red
	}
}
