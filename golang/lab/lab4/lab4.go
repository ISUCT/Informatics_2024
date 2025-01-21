package lab4_test

import (
	"fmt"
	"math"
)

func CalculateY(a, b, x float64) float64 {
	return math.Pow((a*x+b), 1.0/3.0) / math.Pow(math.Log10(x), 2)
}

func TaskA(a, b, xStart, xEnd, step float64) []float64 {
	var result []float64
	for x := xStart; x <= xEnd; x += step {
		result = append(result, CalculateY(a, b, x))
	}
	return result
}

func TaskB(a, b float64, xValues []float64) []float64 {
	var result []float64
	for _, x := range xValues {
		result = append(result, CalculateY(a, b, x))
	}
	return result
}

func labs() {
	var a, b, xStart, xEnd, step float64

	a = 0.1
	b = 0.5
	xStart = 0.33
	xEnd = 1.23
	step = 0.18
	var xValues []float64 = []float64{0.5, 0.36, 0.40, 0.62, 0.78}
	var resultA []float64 = TaskA(a, b, xStart, xEnd, step)
	var resultB []float64 = TaskB(a, b, xValues)

	fmt.Println("Результаты задачи A:")
	for i, y := range resultA {
		x := xStart + step*float64(i)
		fmt.Printf("x = %.2f, y = %.4f\n", x, y)
	}

	fmt.Println("\nРезультаты задачи B:")
	for i, y := range resultB {
		fmt.Printf("x = %.2f, y = %.4f\n", xValues[i], y)
	}
}

func lab4() {
	labs()
}
