package lab4

import (
	"fmt"
	"math"
)

func RunLab4() {
	a := 2.5
	b := 4.6
	fmt.Println(TaskA(1.1, 3.6, 0.5, a, b))
	arr := []float64{1.2, 1.28, 1.36, 1.46, 2.35}
	fmt.Println(TaskB(arr, a, b))
}

func TaskA(xn, xk, deltax, a, b float64) []float64 {
	var yValues []float64
	for x := xn; x <= xk; x += deltax {
		yValues = append(yValues, CalculateY(x, a, b))
	}
	return yValues
}

func TaskB(values []float64, a, b float64) []float64 {
	var yValues []float64
	for _, x := range values {
		yValues = append(yValues, CalculateY(x, a, b))
	}
	return yValues
}

func CalculateY(x float64, a float64, b float64) float64 {
	y := math.Pow(a+b*x, 2.5) / (1 + math.Log10(a+b*x))
	return y
}
