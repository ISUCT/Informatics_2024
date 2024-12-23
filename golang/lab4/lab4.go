package lab4

import (
	"fmt"
	"math"
)

func RunLab4() {
	a := 0.8
	b := 0.4
	fmt.Println(TaskA(1.42, 3.62, 0.44, a, b))
	arr := []float64{1.6, 1.81, 2.24, 2.65, 3.38}
	fmt.Println(TaskB(arr, a, b))
}

func TaskA(xn, xk, deltax, a, b float64) []float64 {
	var yValues []float64
	for x := xn; x <= xk; x += deltax {
		yValues = append(yValues, Calculate(x, a, b))
	}
	return yValues
}

func TaskB(values []float64, a, b float64) []float64 {
	var yValues []float64
	for _, x := range values {
		yValues = append(yValues, Calculate(x, a, b))
	}
	return yValues
}

func Calculate(x float64, a float64, b float64) float64 {
	y := (math.Pow(x-a, 1.0/3.0) + math.Pow(x+b, 1.0/5.0)) / (math.Pow(x, 1.0/7.0) - math.Pow(x*x-(a+b)*(a+b), 1.0/9.0))
	return y
}