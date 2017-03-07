package main

func main() {
	game := Game{term: GetTerminal(), m: CreateMap(), player: Player{hp: 100, loc: Location{x: 5, y: 5}, char: '@'}}
	game.init()
	defer game.end()

	game.draw()

	input := GetInput()

	for {
		game.draw()
		ev := input.getchar()
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
