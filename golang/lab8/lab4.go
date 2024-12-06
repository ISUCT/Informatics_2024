package lab8

import (
	"math"
)

func CalculateFunction(x, b float64) float64 {
	y := (1 + math.Pow(math.Sin(math.Pow(b, 3)+math.Pow(x, 3)), 2)) / (math.Cbrt(math.Pow(b, 3) + math.Pow(x, 3)))
	return y
}

func TaskA(b, xa_beginning, xa_end, xa_delta float64) []float64 {
	var y_a = []float64{}
	for i := xa_beginning; i <= xa_end; i = i + xa_delta {
		y_a = append(y_a, CalculateFunction(i, b))
	}
	return y_a
}

func TaskB(xs []float64, b float64) []float64 {
	var y_b = []float64{}
	for _, x_b := range xs {
		y_b = append(y_b, CalculateFunction(x_b, b))
	}
	return y_b
}
