package piscine

func QuadB(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	for i := 0; i < y; i++ {
		if i == 0 {
			drawLine('/', '*', '\\', x)
		} else if i == y-1 {
			drawLine('\\', '*', '/', x)
		} else {
			drawLine('*', ' ', '*', x)
		}
	}
}

