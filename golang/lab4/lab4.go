package lab4

import (
	"fmt"
	"math"
)

func Calculate(a, b, x float64) float64 {
	y := math.Pow(a*x+b, 1.0/3.0) / math.Pow(math.Log(x), 2.0)
	return y
}

func TaskA(a, b, Xn, Xk, delX float64) []float64 {
	var Calc []float64
	for x := Xn; x <= Xk; x += delX {
		Calc = append(Calc, Calculate(a, b, x))
	}
	return Calc
}

func TaskB(a float64, b float64, x [5]float64) []float64 {
	var Calc []float64
	for _, value := range x {
		Calc = append(Calc, Calculate(a, b, value))
	}
	return Calc
}

func Run4() {
	a := 2.0
	b := 3.0

	fmt.Println(TaskA(a, b, 1.14, 4.24, 0.62))
	var s = [5]float64{0.35, 1.28, 3.51, 5.21, 4.16}
	fmt.Println(TaskB(a, b, s))
}
