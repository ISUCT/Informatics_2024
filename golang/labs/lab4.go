package labs

import (
	"fmt"
	"math"
)

func RunLab4() {
	const const_a = 2
	const const_b = 3
	for _, value := range TaskA(const_a, const_b, 0.11, 0.36, 0.05) {
		fmt.Println(value)
	}
	for _, value := range TaskB(const_a, const_b, []float64{0.08, 0.26, 0.35, 0.41, 0.53}) {
		fmt.Println(value)
	}
}
func ReturnY(x float64, const_a float64, const_b float64) float64 {
	y := ((math.Asin(math.Pow(x, const_a))) + (math.Acos(math.Pow(x, const_b))))
	return y
}
func TaskA(const_a float64, const_b float64, x_n, x_k, x_step float32) []float64 {
	var y_list_taskA []float64
	for x := x_n; x <= x_k; x += x_step {
		y_list_taskA = append(y_list_taskA, ReturnY(float64(x), const_a, const_b))
	}
	return y_list_taskA
}
func TaskB(const_a float64, const_b float64, five_x []float64) []float64 {
	var y_list_taskB []float64
	for _, x := range five_x {
		y_list_taskB = append(y_list_taskB, ReturnY(x, const_a, const_b))
	}
	return y_list_taskB
}
