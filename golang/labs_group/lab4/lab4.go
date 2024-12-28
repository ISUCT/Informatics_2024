package lab4

import (
	"fmt"
	"math"
)
func Calculate(a, b, x float64) float64 {
	var y float64 = (1 + math.Pow(math.Log10(x/a), 2)) / (b - math.Pow(math.Exp(1), (x/a)))
	return y
}
func CompleteTaskA(a, b, xbegin, xend, xdelt float64) []float64 {
	var result []float64
	for x := xbegin; x <= xend; x += xdelt {
		result = append(result, Calculate(a, b, x))
	}
	return result
}
func CompleteTaskB(a, b float64, xs []float64) []float64 {
	var result []float64
	for _, x := range xs {
		result = append(result, Calculate(a, b, x))
	}
	return result
}
func Answerlab() {
	var a float64 = 2.0
	var b float64 = 0.95
	var xbegin float64 = 1.25
	var xend float64 = 2.75
	var xdelt float64 = 0.3
	var xs []float64 = []float64{2.2, 3.78, 4.51, 6.58, 1.2}
	var for_a = CompleteTaskA(a, b, xbegin, xend, xdelt)
	var for_b = CompleteTaskB(a, b, xs)
	fmt.Println("Решение задания А:")
	for i, y := range for_a {
		fmt.Printf("y%d = %f\n", i+1, y)
	}
	fmt.Println("Решение задания B:")
	for i, y := range for_b {
		fmt.Printf("y%d = %f\n", i+1, y)
	}
}
