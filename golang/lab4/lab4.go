package lab4

import (
	"fmt"
	"math"
)

func Calculate(x float64, a float64, b float64) float64 {
	chislitel := (math.Pow(math.Sqrt(math.Pow(x-a, 2)), 1.0/3.0) + math.Pow(math.Sqrt(math.Abs(x+b)), 1.0/5.0))
	znamenatel := math.Pow(math.Sqrt(math.Pow(x, 2)-math.Pow(a+b, 2)), 1.0/9.0)

	var y float64 = chislitel / znamenatel
	return y
}

func TaskA(a, b, xi, xk, deltax float64) [][]float64 {
	var yValues [][]float64
	for x := xi; x <= xk; x += deltax {
		yValues = append(yValues, []float64{x, Calculate(x, a, b)})
	}
	return yValues
}

func TaskB(a, b float64, values []float64) [][]float64 {
	var yValues [][]float64
	for _, x := range values {
		yValues = append(yValues, []float64{x, Calculate(x, a, b)})
	}
	return yValues
}

func Runlab4() {
	a := 0.8
	b := 0.4
	fmt.Println("Задача А")
	for _, pair := range TaskA(a, b, 1.23, 7.23, 1.2) {
		fmt.Printf("x = %.2f\ty = %f\n", pair[0], pair[1])
	}
	var s = []float64{1.88, 2.26, 3.84, 4.55, -6.21}
	fmt.Println("Задача B")
	for _, pair := range TaskB(a, b, s) {
		fmt.Printf("x = %.2f\ty = %f\n", pair[0], pair[1])
	}
}
