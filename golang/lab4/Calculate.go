package lab4

import (
	"math"
)

func Calculate(a, b, x float64) float64 {
	return (1 + math.Pow(math.Log10(x/a), 2)) / (b - math.Pow(2.7, x/a))
}

func TaskA(a, b, Xn, Xk, dX float64) []float64 {
	var yValues []float64
	for x := Xn; x <= Xk; x += dX {
		yValues = append(yValues, Calculate(b, a, x))
	}
	return yValues
}

func TaskB(b, a float64, x [5]float64) []float64 {
	var yValues []float64
	for _, value := range x {
		yValues = append(yValues, Calculate(b, a, value))
	}
	return yValues
}
