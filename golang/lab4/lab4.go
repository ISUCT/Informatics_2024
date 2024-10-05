package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	const const_a, const_b = 2, 3
	task_A(const_a, const_b)
	fmt.Println()
	task_B(const_a, const_b)
}

func return_value_y(x1 float64, const_a float64, const_b float64) float64 {
	y := ((math.Asin(math.Pow(x1, const_a))) + (math.Acos(math.Pow(x1, const_b))))
	return y
}
func task_A(const_a float64, const_b float64) {
	all_x := []float32{}
	var x_n float32 = 0.11
	var x_k float32 = 0.36
	var x_step float32 = 0.05
	for x := x_n; x <= x_k; x += x_step {
		all_x = append(all_x, x)
	}
	for i, x := range all_x {
		fmt.Println(strconv.Itoa(i+1)+":", return_value_y(float64(x), const_a, const_b))
	}
}
func task_B(const_a float64, const_b float64) {
	five_x := [5]float64{0.08, 0.26, 0.35, 0.41, 0.53}
	for i, x := range five_x {
		fmt.Println(strconv.Itoa(i+1)+":", return_value_y(x, const_a, const_b))
	}
}
