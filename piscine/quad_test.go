package piscine

import "fmt"

func drawLine(left, middle, right byte, width int) {
	if width <= 0 {
		return
	}
	if width == 1 {
		fmt.Printf("%c\n", left)
		return
	}
	fmt.Printf("%c", left)
	for i := 0; i < width-2; i++ {
		fmt.Printf("%c", middle)
	}
	fmt.Printf("%c\n", right)
}
