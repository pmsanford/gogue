package main

import "math/rand"

type Goblin struct {
	char	rune
	hp	int
	loc	Location
}

func (goblin *Goblin) act() ActResult {
	dir := rand.Intn(4)
	switch dir {
	case 0:
		goblin.loc.y -= 1
		break
	case 1:
		goblin.loc.x += 1
		break
	case 2:
		goblin.loc.y += 1
		break
	case 3:
		goblin.loc.x -= 1
		break
	}
	return None
}

func (goblin *Goblin) draw(term Terminal) {
	term.draw_char_ex(goblin.loc.x, goblin.loc.y, goblin.char, Green, DefaultColor);
}

func (goblin *Goblin) next() {}