package laba4

import (
	"fmt"
	"math"
)

func СalculateFunction(a, b, x float64) float64 {
	y := (math.Pow(math.Sin(a+b*x), 3.5)) / (1 + math.Cos(math.Log10(a+b*x)))
	return y
}

func completeTaskA(a, b, xMin, xMax, xDelta float64) []float64 {
	var result []float64
	for i := xMin; i < xMax; i += xDelta {
		result = append(result, СalculateFunction(a, b, i))
	}
	return result
}

func completeTaskB(a, b float64, x []float64) []float64 {
	var result []float64
	for _, i := range x {
		result = append(result, СalculateFunction(a, b, i))
	}
	return result
}

func СompleteLaba4() {
	var a float64 = 2.5
	var b float64 = 4.6
	var xMin float64 = 1.15
	var xMax float64 = 3.05
	var xDelta float64 = 0.38
	var x []float64 = []float64{1.2, 1.36, 1.57, 1.93, 2.25}

	var resultA []float64 = completeTaskA(a, b, xMin, xMax, xDelta)
	fmt.Println(resultA)
	var resultB []float64 = completeTaskB(a, b, x)
	fmt.Println(resultB)
}
