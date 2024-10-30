package lab4

import (
	"fmt"
	"math"
)

func ColculateFunction(a, b, x float64) float64 {
	y := (math.Log10(math.Pow(x, 2) - 1)) / (((math.Log(a * math.Pow(x, 2))) / (math.Log(5))) - b)
	return y
}
func completeTaskA(a, b, xn, xk, xd float64) []float64 {
	var result []float64
	for i := xn; i < xk; i += xd {
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
	var a float64 = 1.1
	var b float64 = 0.09
	var x []float64 = []float64{1.21, 1.76, 2.53, 3.48, 4.52}
	var xn float64 = 1.2
	var xk float64 = 2.2
	var xd float64 = 0.2
	var resultA []float64 = completeTaskA(a, b, xn, xk, xd)
	fmt.Println(resultA)
	var resultB []float64 = completeTaskB(a, b, x)
	fmt.Println(resultB)
}
