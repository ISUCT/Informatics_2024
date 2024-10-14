package function

import (
	"fmt"
	"math"
)

func RunLab4Task(x, a, b float64) float64 {

	arccosArg := x*x - b*b
	arcsinArg := x*x - a*a

	if arccosArg < -1 || arccosArg > 1 || arcsinArg < -1 || arcsinArg > 1 {
		return math.NaN()
	}
	asinValue := math.Asin(arcsinArg)
	if asinValue == 0 {
		return math.NaN()
	}
	return math.Acos(arccosArg) / asinValue
}

func Task_A(begin_x, end_x, delta_x, a, b float64) []float64 {
	var answer_arr []float64
	for x := begin_x; x < end_x; x += delta_x {
		answer_arr = append(answer_arr, RunLab4Task(x, a, b))
	}
	return answer_arr
}

func Task_B(arguments []float64, a, b float64) []float64 {
	var answer_arr []float64
	for _, x := range arguments {
		answer_arr = append(answer_arr, RunLab4Task(x, a, b))
	}
	return answer_arr
}

func RunLab4() {
	fmt.Println("------------------------------------------")

	fmt.Println(Task_A(0.05, 0.95, 0.15, 0.06, 0.05))
	fmt.Println("------------------------------------------")

	arr := []float64{0.15, 0.26, 0.37, 0.48, 0.56}
	fmt.Println(Task_B(arr, 0.05, 0.07))
	fmt.Println("------------------------------------------")
}
