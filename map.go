package main

type Map interface {
	draw(term Terminal)
	next()
}

type LevelInfo interface {
	check(loc Location) LocationType
}

type imap struct {
	arr           []rune
	width, height int
}

func (m imap) draw(term Terminal) {
	for i := 0; i < m.height; i++ {
		for j := 0; j < m.width; j++ {
			term.draw_char(j, i, m.arr[(i*m.width)+j])
		}
	}
}

func (m imap) next() {}

func CreateMap() imap {
	m := imap{make([]rune, 1600), 80, 20}
	for index := range m.arr {
		m.arr[index] = '.'
	}
	for i := 0; i < m.width; i++ {
		m.arr[i] = '#'
		m.arr[(m.width*(m.height-1))+i] = '#'
	}
	for i := 0; i < m.height; i++ {
		m.arr[(i * m.width)] = '#'
		m.arr[((i+1)*m.width)-1] = '#'
	}
	return m
}

func (m *imap) check(loc Location) LocationType {
	idx := m.width * loc.y
	idx += loc.x
	if m.arr[idx] == '#' {
		return Blocked
	}
	return Empty
}
