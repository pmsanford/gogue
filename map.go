package main

type Map interface {
	draw(term Terminal)
}

type imap struct {
	arr []rune
	width, height int
}

func (m imap) draw(term Terminal) {
	for i := 0; i < m.height; i++ {
		for j := 0; j < m.width; j++ {
			term.draw_char(j, i, m.arr[(i * m.width) + j])
		}
	}
}

func CreateMap() Map {
	m := imap {make([]rune, 1600), 80, 20}
	for index, _ := range m.arr {
		m.arr[index] = 'h'
	}
	return m
}