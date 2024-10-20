package lab4

import (
	"fmt"
	"math"
)

func ColculateFunction(a, b, x float64) float64 {
	y := math.Pow((a+b*x)/(math.Pow(math.Log10(x), 3)), 0.2)
	return y

}

func completeTaskA(a, b, xMin, xMax, xDelta float64) []float64 {
	var result []float64
	for i := xMin; i < xMax; i += xDelta {
		result = append(result, ColculateFunction(a, b, i))

	}
	return result
}

func completeTaskB(a, b float64, x []float64) []float64 {
	var result []float64
	for _, i := range x {
		result = append(result, ColculateFunction(a, b, i))
	}
	return result

}

func CompleteLab4() {
	var resultA, resultB []float64
	resultA = completeTaskA(7.2, 1.3, 1.56, 4.71, 0.63)
	resultB = completeTaskB(7.2, 1.3, []float64{2.4, 2.8, 3.9, 4.7, 3.16})
	fmt.Println(resultA)
	fmt.Println(resultB)
}
