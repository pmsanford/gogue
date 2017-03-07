package main

import (
	"github.com/nsf/termbox-go"
)

type Input interface {
	getchar() rune
}

type tboxinput struct {
}

func (t tboxinput) getchar() rune {
	return termbox.PollEvent().Ch
}

func GetInput() Input {
	return tboxinput{}
}
