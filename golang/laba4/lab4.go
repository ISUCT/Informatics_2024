package laba4

import (
	"fmt"
	"math"
)

func value_Y(x float64) float64 {
	y := (math.Pow(2.25, (math.Pow(x, 2)-1)) - math.Log10(math.Pow(x, 2)-1) + math.Cbrt((math.Pow(x, 2) - 1)))
	return y
}

func taskA(xstartA, xendA, stepA float64) []float64 {
	var result []float64

	for x := xstartA; x < xendA; x += stepA {
		result = append(result, Value_Y(x))
	}

	return result
}

func taskB(Bslice []float64) []float64 {
	var result []float64
	for _, x := range Bslice {
		result = append(result, Value_Y(x))
	}
	return result
}

func RunLab4() {
	fmt.Println("Задача A:")
	fmt.Println(TaskA(1.2, 2.7, 0.3))
	fmt.Println()
	fmt.Println("Задача Б:")
	x2 := []float64{1.31, 1.39, 1.44, 1.56, 1.92}
	fmt.Println(TaskB(x2))
}
