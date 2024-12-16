package laba4

import (
	"fmt"
	"math"
)

func calculate(x float64) float64 {
	return math.Pow(1.2, x) - math.Pow(x, 1.2)
}

func Y_A(begin_x, end_x, delta_x float64) {
	for begin_x < end_x {
		x := begin_x
		answer := calculate(x)
		fmt.Println(answer, x)

		begin_x += delta_x
	}
}

func Y_B(arr []float64) {
	for _, x := range arr {
		answer := calculate(x)
		fmt.Println(answer)
	}
}

func RunLaba4() {
	var begin_x float64 = 0.2
	var end_x float64 = 2.2
	var delta_x float64 = 0.4
	var arr = []float64{0.1, 0.9, 1.2, 1.5, 2.3}

	Y_A(begin_x, end_x, delta_x)
	Y_B(arr)
}
