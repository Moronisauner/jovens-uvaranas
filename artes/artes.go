package artes

import (
	"fmt"
)

func DesenhaTriangulo(caracter string, l int) {
	for i := 0; i <= l; i++ {
		for j := 0; j < i; j++ {
			fmt.Print(caracter)
		}
		fmt.Println("")
	}
}

func DesenhaLosango(caracter string, l int) {
	mid := l
	for i := 0; i < l; i++ {
		line := make([]int, l*2)

		line[mid-i] = 1
		line[mid+i] = 1
		if i == l-1 {
			fmt.Print("")
		}
		drawLine(line, caracter)
	}

	for i := l - 2; i >= 0; i-- {
		line := make([]int, l*2)

		line[mid-i] = 1
		line[mid+i] = 1
		if i == l-1 {
			fmt.Print("")
		}
		drawLine(line, caracter)
	}
}

func drawLine(line []int, caracter string) {
	for _, i := range line {
		if i == 0 {

			fmt.Print(" ")
		} else {
			fmt.Print(caracter)
		}
	}
	fmt.Println("")
}
