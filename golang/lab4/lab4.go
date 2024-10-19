package lab4

import (
	"math"
)

func Calculator(a, b, x float64) float64 {
	var y float64 = math.Asin(math.Pow(x, a)) + math.Acos(math.Pow(x, b))
	return y
}

func TaskA(a, b, Xn, Xk, delX float64) []float64 {
	var Calc []float64
	for x := Xn; x <= Xk; x += delX {
		Calc = append(Calc, Calculator(a, b, x))
	}
	return Calc
}

func TaskB(a float64, b float64, x [5]float64) []float64 {
	var Calc []float64
	for _, value := range x {
		Calc = append(Calc, Calculator(a, b, value))
	}
	return Calc
}
