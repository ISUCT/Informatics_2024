package lab4

import (
	"math"
)

func Function(a, b, x float64) float64 {
	y := (math.Pow(math.Sin(a+b*x), 3.5)) / (1 + math.Cos(math.Log10(a+b*x)))

	return y
}

func TaskA(a, b, xMin, xMax, xDelta float64) []float64 {
	var data []float64
	for i := xMin; i < xMax; i += xDelta {
		data = append(data, Function(a, b, i))
	}
	return data
}

func TaskB(a, b float64, xs []float64) []float64 {
	var data []float64
	for _, i := range xs {
		data = append(data, Function(a, b, i))
	}
	return data
}
