package main

type Entity interface {
	draw(term Terminal)
	next()
	act() ActResult
}

type Location struct {
	x, y int
}

type Player struct {
	hp    int
	loc   Location
	char  rune
	color Color
	input Input
}

func (p Player) draw(term Terminal) {
	term.draw_char_ex(p.loc.x, p.loc.y, p.char, p.color, DefaultColor)
}

type ActResult uint16

const (
	None ActResult = iota
	Quit
)

func (player *Player) act() ActResult {
	ev := player.input.getchar()
	if ev == 'q' {
		return Quit
	}
	if ev == 'k' {
		player.loc.y -= 1
	}
	if ev == 'j' {
		player.loc.y += 1
	}
	if ev == 'h' {
		player.loc.x -= 1
	}
	if ev == 'l' {
		player.loc.x += 1
	}
	if ev == 'y' {
		player.loc.y -= 1
		player.loc.x -= 1
	}
	if ev == 'u' {
		player.loc.y -= 1
		player.loc.x += 1
	}
	if ev == 'b' {
		player.loc.y += 1
		player.loc.x -= 1
	}
	if ev == 'n' {
		player.loc.y += 1
		player.loc.x += 1
	}
	return None
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
