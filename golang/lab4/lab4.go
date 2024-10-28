package laba4

import (
	"fmt"
	"math"
)

func CalculateFunction(x float64) float64 {
	y := (math.Pow(math.Sqrt(x-2.5), 3) + math.Pow(math.Log10(x), 2)) / (1 + math.Cos(math.Log10(x)))
	return y
}

func CompleteTaskA(xMin, xMax, xDelta float64) []float64 {
	var result []float64
	for i := xMin; i < xMax; i += xDelta {
		result = append(result, CalculateFunction(i))
	}
	return result
}

func CompleteTaskB(x []float64) []float64 { // Удалили лишний аргумент float64
	var result []float64
	for _, i := range x {
		result = append(result, CalculateFunction(i))
	}
	return result
}

func CompleteLaba4() {
	var xMin float64 = 1.25
	var xMax float64 = 3.25
	var xDelta float64 = 0.4
	var x []float64 = []float64{1.84, 2.71, 3.81, 4.56, 5.62}

	var resultA []float64 = CompleteTaskA(xMin, xMax, xDelta)
	fmt.Println(resultA)
	var resultB []float64 = CompleteTaskB(x) // Удалили лишний аргумент x
	fmt.Println(resultB)
}
