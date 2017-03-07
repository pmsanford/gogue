package main

func main() {
	game := new_game()
	game.init()
	defer game.end()

	game.draw()
	game.loop()
}
