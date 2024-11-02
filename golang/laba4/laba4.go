package laba4

import (
	"fmt"
	"math"
)

func Calculate(a, b, x float64) float64 {
	return (math.Sqrt(math.Abs(a-b*x)) / math.Pow(math.Log(x), 3))
}

func TaskA(a, b, xn, xk, deltax float64) [][]float64 {
	var taskAanswer [][]float64
	for x := xn; x < xk; x += deltax {
		taskAanswer = append(taskAanswer, []float64{x, Calculate(a, b, x)})
	}
	return taskAanswer
}

func TaskB(a, b float64, taskBslice []float64) [][]float64 {
	var taskBanswer [][]float64
	for _, x := range taskBslice {
		taskBanswer = append(taskBanswer, []float64{x, Calculate(a, b, x)})
	}
	return taskBanswer
}

func Runlab4() {
	var a float64 = 7.2
	var b float64 = 4.2
	var taskBslice = []float64{2.4, 2.8, 3.9, 4.7, 3.16}
	fmt.Println("Задача А")
	for _, pair := range TaskA(a, b, 1.81, 5.31, 0.7) {
		x := pair[0]
		y := pair[1]
		fmt.Printf("x = %.2f\ty = %f\n", x, y)
	}
	fmt.Println()
	fmt.Println("Задание B")
	for _, pair := range TaskB(a, b, taskBslice) {
		x := pair[0]
		y := pair[1]
		fmt.Printf("x = %.2f\ty = %f\n", x, y)
	}
}
