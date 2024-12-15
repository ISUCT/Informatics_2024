package lab4

import (
	"fmt"
	"math"
)

func RunLab4() {
	a, b := 2.5, 4.6
	xn, xk, deltaX := 1.1, 3.6, 0.5
	values := []float64{1.2, 1.28, 1.36, 1.46, 2.35}

	fmt.Println("Task A:", SolveTaskA(xn, xk, deltaX, a, b))
	fmt.Println("Task B:", SolveTaskB(values, a, b))
}

func SolveTaskA(xn, xk, step, a, b float64) []float64 {
	return CalculateRange(xn, xk, step, a, b)
}

func SolveTaskB(values []float64, a, b float64) []float64 {
	return CalculateList(values, a, b)
}

func CalculateRange(start, end, step, a, b float64) []float64 {
	var results []float64
	for x := start; x <= end; x += step {
		results = append(results, calculateY(x, a, b))
	}
	return results
}

func CalculateList(values []float64, a, b float64) []float64 {
	var results []float64
	for _, x := range values {
		results = append(results, calculateY(x, a, b))
	}
	return results
}

func calculateY(x, a, b float64) float64 {
	return math.Pow(a+b*x, 2.5) / (1 + math.Log10(a+b*x))
}
