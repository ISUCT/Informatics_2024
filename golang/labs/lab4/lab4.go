package lab4

import (
	"fmt"
	"math"
)

func calculateY(b float64, x float64) float64 {
	return (1 + math.Pow(math.Sin(math.Pow(b, 3)+math.Pow(x, 3)), 2)) / math.Cbrt(math.Pow(b, 3)+math.Pow(x, 3))
}

func TaskA(b float64, xStart float64, xEnd float64, deltaX float64) []float64 {
	result := []float64{}
	for x := xStart; x <= xEnd; x += deltaX {
		y := calculateY(b, x)
		result = append(result, y)
	}
	return result
}

func TaskB(b float64, xValues []float64) []float64 {
	result := []float64{}
	for _, x := range xValues {
		y := calculateY(b, x)
		result = append(result, y)
	}
	return result
}

func RunLab4() {
	b := 2.5
	xStart := 1.28
	xEnd := 3.28
	deltaX := 0.4

	fmt.Println("Значение A:")
	resultA := TaskA(b, xStart, xEnd, deltaX)
	for _, y := range resultA {
		fmt.Printf("y: %.4f\n", y)
	}

	xValues := []float64{1.1, 2.4, 3.6, 1.7, 3.9}

	fmt.Println("\nЗначение B:")
	resultB := TaskB(b, xValues)
	for _, y := range resultB {
		fmt.Printf("y: %.4f\n", y)
	}
}
