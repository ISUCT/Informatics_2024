package main

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

func main() {
	a := 2.0
	b := 3.0

	xStart := 1.0
	xEnd := 5.0
	step := 1.0
	resultsA := TaskA(a, b, xStart, xEnd, step)
	fmt.Println("Результаты TaskA:")
	for i, y := range resultsA {
		fmt.Printf("x = %.2f, y = %.6f\n", xStart+float64(i)*step, y)
	}

	fmt.Println()

	xValues := []float64{1.5, 2.5, 3.5, 4.5}
	resultsB := TaskB(a, b, xValues)
	fmt.Println("Результаты TaskB:")
	for i, y := range resultsB {
		fmt.Printf("x = %.2f, y = %.6f\n", xValues[i], y)
	}
	fmt.Println()
	x := 3.0
	resultY := CalculateY(a, b, x)
	fmt.Printf("Результат CalculateY: x= %.2f, y = %.6f\n", x, resultY)
}
