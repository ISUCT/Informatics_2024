package labs

import (
	"math"
)

// Лабораторная 4, задание по функции 6
func y(x float64) float64 {
	return math.Pow(1.2, x) - math.Pow(x, 1.2)
}

// Задача под А
func Task_A(begin_x, end_x, delta_x float64) []float64 {
	answer_arr := make([]float64, 0)

	for begin_x < end_x {
		x := begin_x
		answer_arr = append(answer_arr, y(x))
		begin_x += delta_x
	}

	return answer_arr
}

// Задача под B
func Task_B(arguments []float64) []float64 {
	answer_arr := make([]float64, 0)

	for _, x := range arguments {
		answer_arr = append(answer_arr, y(x))
	}

	return answer_arr
}
