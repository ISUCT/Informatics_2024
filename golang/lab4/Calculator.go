package lab4

import (
	"fmt"
	"math"
)

func Calculate(a, b, x float64) float64 {
	return (1 + math.Pow(math.Log10(x/a), 2)) / (b - math.Pow(2.7, x/a))
}

func TaskA(a, b, Xn, Xk, delX float64) []float64 {
	var y []float64
	for x := Xn; x <= Xk; x += delX {
		y = append(y, Calculate(a, b, x))
	}
	return y
}

func TaskB(a, b float64, x [5]float64) []float64 {
	var y []float64
	for _, value := range x {
		y = append(y, Calculate(a, b, value))
	}
	return y
}

func RunLab4() {
	a := 2.0
	b := 0.95
	fmt.Println(TaskA(a, b, 1.25, 2.75, 0.3))
	var p = [5]float64{2.2, 3.78, 4.51, 6.58, 1.2}
	fmt.Println(TaskB(a, b, p))
}
