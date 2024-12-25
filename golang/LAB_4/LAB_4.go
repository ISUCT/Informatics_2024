package Lab4

import (
	"fmt"
	"math"
)

func Calculate(a, b, x float64) float64 {
	numerator := (a * math.Sqrt(x)) - (b * (math.Log(x) / math.Log(5)))
	denominator := math.Log10(math.Abs(x - 1))
	return numerator / denominator
}

func TaskA(a, b, xStart, xEnd, xStep float64) []float64 {
	var YValues []float64
	for x := xStart; x <= xEnd; x += xStep {
		YValues = append(YValues, Calculate(a, b, x))
	}
	return YValues
}

func TaskB(a, b float64, xValues []float64) []float64 {
	var YValues []float64
	for _, x := range xValues {
		YValues = append(YValues, Calculate(a, b, x))
	}
	return YValues
}

func RunLab4() {
	a := 4.1
	b := 2.7
	xStart := 1.2
	xEnd := 5.2
	xStep := 0.8
	var xValues []float64 = []float64{1.84, 2.71, 3.81, 4.56, 5.62}
	var resultA []float64 = TaskA(a, b, xStart, xEnd, xStep)
	var resultB []float64 = TaskB(a, b, xValues)

	fmt.Println("задача A", resultA)
	fmt.Println("задача B", resultB)
}
