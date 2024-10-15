package laba4

import (
	"math"
)

func Function(a, b, x float64) float64 {
	y := (math.Pow(math.Sin(a+b*x), 3.5)) / (1 + math.Cos(math.Log10(a+b*x)))
	return y
}

func TaskA(a, b, xMin, xMax, xDelta float64) []float64 {
	var result []float64
	for i := xMin; i < xMax; i += xDelta {
		result = append(result, Function(a, b, i))
	}
	return result
}

func TaskB(a, b float64, x []float64) []float64 {
	var result []float64
	for _, i := range x {
		result = append(result, Function(a, b, i))
	}
	return result
}
