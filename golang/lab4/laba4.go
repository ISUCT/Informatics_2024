package laba4

import (
	"fmt"
	"math"
)

func Culculation(a, b, x float64) float64 {
	y := (math.Pow(a, x) - math.Pow(b, x)*math.Cbrt(a*b)) / math.Log10(a/b)
	return y
}
func CompleteA(a, b, xMin, xMax, xDelta float64) []float64 {
	var y []float64
	for i := xMin; i < xMax; i += xDelta {
		y = append(y, Culculation(a, b, i))
	}
	return y
}

func CompleteB(a, b float64, x []float64) []float64 {
	var y []float64
	for _, i := range x {
		y = append(y, Culculation(a, b, i))
	}
	return y
}
func Lab4() {
	var a float64 = 0.4
	var b float64 = 0.8
	var x []float64 = []float64{4.48, 3.56, 2.78, 5.28, 3.21}
	var xMin float64 = 3.2
	var xMax float64 = 6.2
	var xDelta float64 = 0.6

	var resultA []float64 = CompleteA(a, b, xMin, xMax, xDelta)
	fmt.Println(resultA)
	var resultB []float64 = CompleteB(a, b, x)
	fmt.Println(resultB)
}
