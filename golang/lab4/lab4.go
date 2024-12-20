package lab4

import (
	"fmt"
	"math"
)

func Calculate(a, b, x float64) float64 {
	var y float64 = math.Asin(math.Pow(x, a)) + math.Acos(math.Pow(x, b))
	return y
}

func TaskA(a, b, Xn, Xk, delX float64) []float64 {
	var Calc []float64
	for x := Xn; x <= Xk; x += delX {
		Calc = append(Calc, Calculate(a, b, x))
	}
	return Calc
}

func TaskB(a float64, b float64, x []float64) []float64 {
	var Calc []float64
	for _, value := range x {
		Calc = append(Calc, Calculate(a, b, value))
	}
	return Calc
}

func RunLab4() {
	a := 2.0
	b := 3.0
	fmt.Println(TaskA(a, b, 0.11, 0.36, 0.05))
	var s = []float64{0.08, 0.26, 0.35, 0.41, 0.53}
	fmt.Println(TaskB(a, b, s))
}
