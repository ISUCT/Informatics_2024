package labs

import (
	"fmt"
	"math"
)

func RunLab4() {
	fmt.Println(TaskA(1.1, 3.6, 0.5))
	arr := []float64{1.2, 1.28, 1.36, 1.46, 2.35}
	fmt.Println(TaskB(arr))
}

func TaskA(xn, xk, deltax float64) []float64 {
	var yValues []float64
	for x := xn; x <= xk; x += deltax {
		yValues = append(yValues, Calculate_y(x, 2.5, 4.6))
	}
	return yValues
}

func TaskB(values []float64) []float64 {
	var yValues []float64
	for _, x := range values {
		yValues = append(yValues, Calculate_y(x, 2.5, 4.6))
	}
	return yValues
}

func Calculate_y(x float64, a float64, b float64) float64 {
	y := math.Pow(a+b*x, 2.5) / (1 + math.Log10(a+b*x))
	return y
}
