package main

type Entity interface {
	draw(term Terminal)
	next()
	act(lvl LevelInfo) ActResult
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

func (player Player) draw(term Terminal) {
	term.draw_char_ex(player.loc.x, player.loc.y, player.char, player.color, DefaultColor)
}

type ActResult uint16

const (
	None ActResult = iota
	Acted
	Quit
)

func (player *Player) act(lvl LevelInfo) ActResult {
	ev := player.input.getchar()
	if ev == 'q' {
		return Quit
	}
	newloc := player.loc
	if ev == 'k' {
		newloc.y -= 1
	}
	if ev == 'j' {
		newloc.y += 1
	}
	if ev == 'h' {
		newloc.x -= 1
	}
	if ev == 'l' {
		newloc.x += 1
	}
	if ev == 'y' {
		newloc.y -= 1
		newloc.x -= 1
	}
	if ev == 'u' {
		newloc.y -= 1
		newloc.x += 1
	}
	if ev == 'b' {
		newloc.y += 1
		newloc.x -= 1
	}
	if ev == 'n' {
		newloc.y += 1
		newloc.x += 1
	}
	if lvl.check(newloc) == Empty {
		player.loc = newloc
		return Acted
	}
	return None
}

func (player *Player) next() {
	if player.color == Red {
		player.color = White
	} else if player.color == White {
		player.color = Green
	} else if player.color == Green {
		player.color = Red
	}
}
