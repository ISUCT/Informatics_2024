package labs

import (
	"fmt"
	"math"
)

// Лабораторная 4, задание по функции 6
func y(x float64) float64 {
	return math.Pow(1.2, x) - math.Pow(x, 1.2)
}

// Задача под А
func Task_A() {
	var begin_x float64 = 0.2
	var end_x float64 = 2.2
	var delta_x float64 = 0.4

	for begin_x < end_x {
		x := begin_x
		answer := y(x)
		fmt.Println(answer, x)

		begin_x += delta_x
	}
}

// Задача под B
func Task_B() {
	var arr = [5]float64{0.1, 0.9, 1.2, 1.5, 2.3}
	for _, x := range arr {
		answer := y(x)
		fmt.Println(answer)
	}
}
