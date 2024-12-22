package lab4

import "math"

const a = 0.4
const b = 0.8

func calculateY(a, b, x float64) float64 {
	var y float64 = (math.Pow(a, x) - math.Pow(b, x)) / (math.Log10(a/b) * math.Cbrt(a*b))
	return y
}

func Task_A(xn, xk, delx float64) []float64 {
	var yValues []float64
	for x := xn; x <= xk; x += delx {
		yValues = append(yValues, calculateY(a, b, x))
	}
	return yValues
}

func Task_B(xv []float64) []float64 {
	var yValues []float64
	for _, x := range xv {
		yValues = append(yValues, calculateY(a, b, x))
	}
	return yValues
}
