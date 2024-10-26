package lab4

import (
	"fmt"
	"math"
)

func CompleteLab4() {
	resultA := CompleteTaskA(2.5, 4.6, 1.15, 3.05, 0.38)
	fmt.Println(resultA)
	resultB := CompleteTaskB(2.5, 4.6, []float64{1.2, 1.36, 1.57, 1.93, 2.25})
	fmt.Println(resultB)
}

func Calculation(a, b, x float64) float64 {
	y := (math.Pow(math.Sin(a+b*x), 3.5)) / (1 + math.Cos(math.Log10(a+b*x)))

	return y
}

func CompleteTaskA(a, b, xMin, xMax, xDelta float64) []float64 {
	var data []float64
	for x := xMin; x < xMax; x += xDelta {
		data = append(data, Calculation(a, b, x))
	}
	return data
}

func CompleteTaskB(a, b float64, xs []float64) []float64 {
	var data []float64
	for _, x := range xs {
		data = append(data, Calculation(a, b, x))
	}
	return data
}
