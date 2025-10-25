package piscine

func QuadA(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	for i := 0; i < y; i++ {
		if i == 0 {
			drawLine('o', '-', 'o', x)
		} else if i == y-1 {
			drawLine('o', '-', 'o', x)
		} else {
			drawLine('|', ' ', '|', x)
		}
	}
}
