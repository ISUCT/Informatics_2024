package labs

import (
	"fmt"
	"math"
)

// Лабораторная 4, задание по функции 6
func calculate_y(x float64) float64 {
	if math.Abs(x) >= 1 {
		return math.Pow(1.2, x) - math.Pow(x, 1.2)
	} else {
		return math.Acos(x)
	}
}

// Задача под А
func Task_A(begin_x, end_x, delta_x float64) []float64 {
	answer_arr := make([]float64, 0)

	for i := begin_x; i < end_x; i += delta_x {
		x := i
		answer_arr = append(answer_arr, calculate_y(x))
	}

	return answer_arr
}

// Задача под B
func Task_B(arguments []float64) []float64 {
	answer_arr := make([]float64, 0)

	for _, x := range arguments {
		answer_arr = append(answer_arr, calculate_y(x))
	}

	return answer_arr
}

func RunLab4Task() {
	fmt.Println("------------------------------------------")
	fmt.Println(Task_A(0.2, 2.2, 0.4))
	fmt.Println("------------------------------------------")
	arr := [5]float64{0.1, 0.9, 1.2, 1.5, 2.3}
	fmt.Println(Task_B(arr[:]))
	fmt.Println("------------------------------------------")
}
