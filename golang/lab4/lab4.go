package lab4

import (
	"fmt"
	"math"
)

func calculate(a, b, x float64) float64 {
	y := (math.Pow(a, 3)*math.Sqrt(x) - b*math.Log(x)) / (math.Pow(math.Log10(x-1), 3))

	return y
}
func completeTaskA(a, b, xMin, xMax, xDelta float64) []float64 {
	var result []float64
	for i := xMin; i < xMax; i += xDelta {
		result = append(result, calculate(a, b, i))
	}
	return result
}
func completeTaskB(a, b float64, xPlenty []float64) []float64 {
	var result []float64
	for _, x := range xPlenty {
		result = append(result, calculate(a, b, x))
	}
	return result
}
func CompleteLab4() {
	var a, b, xMin, xMax, xDelta float64
	var resultA, resultB []float64
	var xPlenty []float64 = []float64{1.9, 2.15, 2.34, 2.74, 3.16}
	a = 4.1
	b = 2.7
	xMin = 1.5
	xMax = 3.5
	xDelta = 0.4
	resultA = completeTaskA(a, b, xMin, xMax, xDelta)
	resultB = completeTaskB(a, b, xPlenty)
	fmt.Println(resultA)
	fmt.Println(resultB)
}
