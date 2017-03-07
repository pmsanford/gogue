package main

import (
	"github.com/nsf/termbox-go"
)

type Drawable interface {
	draw(term Terminal)
}

type Color uint16

const (
	DefaultColor Color = iota
	Black              = 0x1
	Red                = 0x2
	Green              = 0x3
	Yellow             = 0x4
	Blue               = 0x5
	Magenta            = 0x6
	Cyan               = 0x7
	White              = 0x8
)

type Terminal interface {
	init()
	draw_char(x, y int, c rune)
	draw_char_ex(x, y int, c rune, fg, bg Color)
	draw_str(x, y int, s string)
	flush()
	close()
}

type tboxterm struct {
	fg, bg Color
}

func (t tboxterm) draw_char(x, y int, c rune) {
	t.draw_char_ex(x, y, c, t.fg, t.bg)
}

func (t tboxterm) draw_char_ex(x, y int, c rune, fg, bg Color) {
	termbox.SetCell(x, y, c, termbox.Attribute(fg), termbox.Attribute(bg))
}

func (t tboxterm) draw_str(x, y int, s string) {
	for _, c := range s {
		t.draw_char(x, y, c)
		x++
	}
}

func (t tboxterm) init() {
	termbox.Init()
	termbox.SetOutputMode(termbox.Output256)
	termbox.Clear(termbox.ColorWhite, termbox.ColorBlack)
}

func (t tboxterm) flush() {
	termbox.Flush()
}

func (t tboxterm) close() {
	termbox.Close()
}

func GetTerminal() Terminal {
	return tboxterm{fg: White, bg: Black}
}
