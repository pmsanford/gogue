package main

func main() {
	game := Game{term: GetTerminal(), m: CreateMap(), player: Player{hp: 100, loc: Location{x: 5, y: 5}, char: '@'}, input: GetInput()}
	game.init()
	defer game.end()

	game.draw()
	game.loop()
}
