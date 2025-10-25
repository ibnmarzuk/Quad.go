package piscine

func QuadE(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	for i := 0; i < y; i++ {
		if i == 0 {
			drawLine('A', 'B', 'C', x)
		} else if i == y-1 {
			drawLine('C', 'B', 'A', x)
		} else {
			drawLine('B', ' ', 'B', x)
		}
	}
}

