package lab4

import (
	"fmt"
	"math"
)

func Calculate(a, b, x float64) float64 {
	return (math.Cbrt(a) + math.Pow(math.Tan(b*x), 4.5)) / (math.Pow(b, 0.2) + (1 / (math.Pow(math.Tan(a*x), 2.7))))
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

func RunLab4() {
	var a float64 = 0.1
	var b float64 = 0.5
	var taskBslice = []float64{1.24, 1.38, 2.38, 3.21, 0.68}
	fmt.Println("Задача А")
	for _, pair := range TaskA(a, b, 0.33, 1.23, 0.18) {
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
