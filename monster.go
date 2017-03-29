package main

import (
	"math/rand"
)

type Monster struct {
	char rune
	hp   int
	loc  Location
}

type MonsterType uint16

const (
	Goblin MonsterType = iota
	Troll
)

func (goblin *Monster) act(lvl LevelInfo) ActResult {
	newloc := goblin.loc
	for {
		newloc = goblin.loc
		dir := rand.Intn(4)
		switch dir {
		case 0:
			newloc.y -= 1
			break
		case 1:
			newloc.x += 1
			break
		case 2:
			newloc.y += 1
			break
		case 3:
			newloc.x -= 1
			break
		}
		if lvl.check(newloc) == Empty{
			break
		}
	}
	goblin.loc = newloc
	return Acted
}

func (goblin *Monster) draw(term Terminal) {
	term.draw_char_ex(goblin.loc.x, goblin.loc.y, goblin.char, Green, DefaultColor)
}

func (goblin *Monster) next() {}

type MonsterGenerator struct {
	monster_table map[int]MonsterType
}

func (generator MonsterGenerator) get_monster(typ MonsterType) Monster {
	switch typ {
	case Goblin:
		return Monster{'g', 100, Location{5, 5}}
	case Troll:
		return Monster{'T', 150, Location{5, 5}}
	}
	return Monster{}
}

func (generator MonsterGenerator) new_monster() Monster {
	max := 0
	for k := range generator.monster_table {
		max = k
	}
	mon_val := rand.Intn(max)
	for k, v := range generator.monster_table {
		if mon_val < k {
			return generator.get_monster(v)
		}
	}
	return generator.get_monster(Goblin)
}
