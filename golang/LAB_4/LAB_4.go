package LAB_4

import (
	"fmt"
	"math"
)

func CalculateY(a, b, x float64) float64 {
	numerator := math.Pow(a, 1.0/3.0) + math.Pow(math.Tan(b*x), 4.5)
	denominator := math.Pow(b, 1.0/5.0) + 1/math.Pow(math.Tan(a*x), 2.7)
	return numerator / denominator
}

func TaskA(a, b, xStart, xEnd, step float64) {
	for x := xStart; x <= xEnd; x += step {
		y := CalculateY(a, b, x)
		fmt.Printf("For x = %.2f, y = %.10f\n", x, y)
	}
}

func TaskB(a, b float64, xValues []float64) {
	for _, x := range xValues {
		y := CalculateY(a, b, x)
		fmt.Printf("For x = %.2f, y = %.10f\n", x, y)
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

func Lab4() {
	var a, b, xStart, xEnd, step float64

	a = 0.1
	b = 0.5
	xStart = 0.33
	xEnd = 1.23
	step = 0.18

	TaskA(a, b, xStart, xEnd, step)

	xValues := []float64{0.5, 0.36, 0.40, 0.62, 0.78}
	TaskB(a, b, xValues)
	var xValues []float64 = []float64{0.5, 0.36, 0.40, 0.62, 0.78}
	var resultA []float64 = TaskA(a, b, xStart, xEnd, step)
	var resultB []float64 = TaskB(a, b, xValues)

	fmt.Println(resultA)
	fmt.Println(resultB)
}
