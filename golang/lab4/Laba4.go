package LABA4

import (
	"fmt"
	"math"
)

func CalculateY(a, b, x float64) float64 {
	numerator := math.Cbrt(a*x + b)
	denominator := math.Pow(math.Log10(x), 2)
	return numerator / denominator
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
	a := 1.35
	b := 0.98

	fmt.Println("Задача A")
	xStartA := 1.14
	xEndA := 4.24
	dxA := 0.62
	resultA := TaskA(a, b, xStartA, xEndA, dxA)

	fmt.Println("\nЗадача B")
	xValuesB := []float64{0.35, 1.28, 3.51, 5.21, 4.16}
	resultB := TaskB(a, b, xValuesB)

	fmt.Println(resultA)
	fmt.Println(resultB)
}
