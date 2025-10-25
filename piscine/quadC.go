package piscine

func QuadC(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	for i := 0; i < y; i++ {
		if i == 0 {
			drawLine('A', 'B', 'A', x)
		} else if i == y-1 {
			drawLine('C', 'B', 'C', x)
		} else {
			drawLine('B', ' ', 'B', x)
		}
	}
}

