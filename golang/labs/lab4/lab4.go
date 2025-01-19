package function

import (
	"fmt"
	"math"
)

const epsilon = 1e-9

func CalculateY(x, a, b float64) float64 {
	arccosArg := x*x - b*b
	arcsinArg := x*x - a*a

	fmt.Printf("x: %f, a: %f, b: %f\n", x, a, b)
	fmt.Printf("arccosArg: %f, arcsinArg: %f\n", arccosArg, arcsinArg)

	if arccosArg < -1 || arccosArg > 1 || arcsinArg < -1 || arcsinArg > 1 {
		fmt.Println("Out of range")
		return math.NaN()
	}

	asinValue := math.Asin(arcsinArg)
	fmt.Printf("asinValue: %f\n", asinValue)

	if math.Abs(asinValue) < epsilon {
		fmt.Println("asinValue too small")
		return math.NaN()
	}

	result := math.Acos(arccosArg) / asinValue
	fmt.Printf("Result: %f\n", result)

	return result
}

func Task_A(begin_x, end_x, delta_x, a, b float64) []float64 {
	var answer_arr []float64
	for x := begin_x; x < end_x; x += delta_x {
		answer_arr = append(answer_arr, CalculateY(x, a, b))
	}
	return answer_arr
}

func Task_B(arguments []float64, a, b float64) []float64 {
	var answer_arr []float64
	for _, x := range arguments {
		answer_arr = append(answer_arr, CalculateY(x, a, b))
	}
	return answer_arr
}

func RunLab4Task() {
	fmt.Println("------------------------------------------")

	resultA := Task_A(0.05, 0.95, 0.15, 0.06, 0.05)
	fmt.Println(resultA)
	fmt.Println("------------------------------------------")

	arr := []float64{0.15, 0.26, 0.37, 0.48, 0.56}
	resultB := Task_B(arr, 0.05, 0.07)
	fmt.Println(resultB)
	fmt.Println("------------------------------------------")
}
