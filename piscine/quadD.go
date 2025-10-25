package piscine

func QuadD(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	for i := 0; i < y; i++ {
		if i == 0 {
			drawLine('A', 'B', 'C', x)
		} else if i == y-1 {
			drawLine('A', 'B', 'C', x)
		} else {
			drawLine('B', ' ', 'B', x)
		}
	}
}

