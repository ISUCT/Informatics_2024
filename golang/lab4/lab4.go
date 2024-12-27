package lab4

import (
	"fmt"
	"math"
)

func calculate(a, b, x float64) float64 {
	return (1 + math.Pow(math.Log10(x/a), 2)) / (b - math.Pow(2.7, x/a))
}

func TaskA(a, b, Xn, Xk, dX float64) []float64 {
	var yValues []float64
	for x := Xn; x <= Xk; x += dX {
		yValues = append(yValues, calculate(b, a, x))
	}
	return yValues
}

func TaskB(b, a float64, x [5]float64) []float64 {
	var yValues []float64
	for _, value := range x {
		yValues = append(yValues, calculate(b, a, value))
	}
	return yValues
}

func RunLab4() {
	fmt.Println("Task А:\n", TaskA(0.95, 2.0, 1.25, 2.75, 0.3))
	fmt.Println("Task В:\n", TaskB(2.0, 0.95, [5]float64{2.2, 3.78, 4.51, 6.58, 1.2}))
}