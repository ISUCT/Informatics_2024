package main

import (
	"fmt"
	"math"
)

func Calculate(a, b, x float64) float64 {
	numerator := (math.Pow(math.Sqrt(math.Pow(x-a, 2)), 1.0/3.0) + math.Pow(math.Sqrt(math.Abs(x+b)), 1.0/5.0))
	denominator := math.Pow(math.Sqrt(math.Pow(x, 2)-math.Pow(a+b, 2)), 1.0/9.0)
	var y float64 = numerator / denominator
	return y
}

func TaskA(a, b, xn, xk, deltax float64) [][]float64 {
	var taskAslice [][]float64

	for x := xn; x <= xk; x += deltax {
		y := Calculate(a, b, x)

		taskAslice = append(taskAslice, []float64{x, y})
	}

	return taskAslice
}

func TaskB(a, b float64, xValue []float64) [][]float64 {
	var taskBslice [][]float64

	for _, x := range xValue {
		y := Calculate(a, b, x)

		taskBslice = append(taskBslice, []float64{x, y})
	}

	return taskBslice
}

func main() {
	var a float64 = 0.8
	var b float64 = 0.4
	var xValue = []float64{1.88, 2.26, 3.84, 4.55, -6.21}
	fmt.Println("Задача А")
	for _, pair := range TaskA(a, b, 1.23, 7.23, 1.2) {
		fmt.Printf("x = %.2f\ty = %f\n", pair[0], pair[1])
	}
	fmt.Println("Задание B")
	for _, pair := range TaskB(a, b, xValue) {
		fmt.Printf("x = %.2f\ty = %f\n", pair[0], pair[1])
	}
}
