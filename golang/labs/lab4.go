package labs

import (
	"fmt"
	"math"
)

func Go_lab4() {
	fmt.Println()
	const const_a = 2
	const const_b = 3
	for _, value := range Task_A(const_a, const_b) {
		fmt.Println(value)
	}
	println("")
	for _, value := range Task_B(const_a, const_b) {
		fmt.Println(value)
	}
}
func Return_value_y(x1 float64, const_a float64, const_b float64) float64 {
	y := ((math.Asin(math.Pow(x1, const_a))) + (math.Acos(math.Pow(x1, const_b))))
	return y
}
func Task_A(const_a float64, const_b float64) []float64 {
	all_x := []float32{}
	var x_n float32 = 0.11
	var x_k float32 = 0.36
	var x_step float32 = 0.05
	for x := x_n; x <= x_k; x += x_step {
		all_x = append(all_x, x)
	}

	y_list_taskA := []float64{}
	for _, x := range all_x {
		y_list_taskA = append(y_list_taskA, Return_value_y(float64(x), const_a, const_b))
	}
	return y_list_taskA
}
func Task_B(const_a float64, const_b float64) []float64 {
	five_x := [5]float64{0.08, 0.26, 0.35, 0.41, 0.53}
	y_list_taskB := []float64{}
	for _, x := range five_x {
		y_list_taskB = append(y_list_taskB, Return_value_y(x, const_a, const_b))
	}
	return y_list_taskB
}
