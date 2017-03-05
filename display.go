package main

import (
	//"fmt"
	"github.com/nsf/termbox-go"
)

type Drawable interface {
	draw(term Terminal)
}

type Terminal interface {
	init()
	draw_char(x, y int, c rune)
	draw_str(x, y int, s string)
	flush()
	close()
}

type tboxterm struct {
	fg, bg termbox.Attribute
}

func (t tboxterm) draw_char(x, y int, c rune) {
	termbox.SetCell(x, y, c, t.fg, t.bg)
}

func (t tboxterm) draw_str(x, y int, s string) {
	for _, c := range s {
		t.draw_char(x, y, c)
		x++
	}	
}

func (t tboxterm) init() {
	termbox.Init()
	termbox.Clear(termbox.ColorWhite, termbox.ColorBlack);
}

func (t tboxterm) flush() {
	termbox.Flush()
}

func (t tboxterm) close() {
	termbox.Close()
}

func GetTerminal() Terminal {
	return tboxterm {fg: termbox.ColorWhite, bg: termbox.ColorBlack }
}