package fourlaba

import (
	"fmt"
	"math"
)

func CalculateY(a, b, x float64) float64 {
	numeratorPart := b*b - x*x

	logTerm := math.Log(numeratorPart) / math.Log(a)

	denominator := math.Cbrt(math.Abs(x*x - a*a))

	y := logTerm / denominator

	return y
}

func SolveTaskA(a, b, xStart, xEnd, step float64) []float64 {
	results := []float64{}
	for x := xStart; x <= xEnd; x += step {
		y := CalculateY(a, b, x)
		results = append(results, y)
	}
	return results
}

func SolveTaskB(a, b float64, xValues []float64) []float64 {
	results := []float64{}
	for _, x := range xValues {
		y := CalculateY(a, b, x)
		results = append(results, y)
	}
	return results
}
func StartLab4() {
	a := 2.0
	b := 4.1

	fmt.Println("Задача A:")
	xStartA := 0.77
	xEndA := 1.77
	stepA := 0.2
	resultsA := SolveTaskA(a, b, xStartA, xEndA, stepA)
	fmt.Println(resultsA)

	fmt.Println("Задача B:")
	xValuesB := []float64{1.24, 1.38, 2.38, 3.21, 0.68}
	resultsB := SolveTaskB(a, b, xValuesB)
	fmt.Println(resultsB)
}
