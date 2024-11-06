package labs

import (
	"fmt"
	"math"
)

func CalculateY(x float64, a float64, b float64) float64 {
	y := math.Acos(math.Pow(x, 2)-math.Pow(b, 2)) / math.Asin(math.Pow(x, 2)-math.Pow(a, 2))
	return y
}

func TaskA(a, b, xn, xk, deltax float64) [][]float64 {
	var taskA [][]float64
	for x := xn; x <= xk; x += deltax {
		y := CalculateY(x, a, b)
		taskA = append(taskA, []float64{x, y})
	}
	return taskA
}

func TaskB(a, b float64, values []float64) [][]float64 {
	var taskB [][]float64
	for _, x := range values {
		y := CalculateY(x, a, b)
		taskB = append(taskB, []float64{x, y})
	}
	return taskB
}

func RunLab4() {
	resultsA := TaskA(0.05, 0.06, 0.2, 0.95, 0.15)
	for _, result := range resultsA {
		fmt.Printf("x: %.2f, y: %.2f\n", result[0], result[1])
	}

	arr := []float64{0.15, 0.26, 0.37, 0.48, 0.56}
	resultsB := TaskB(0.2, 0.95, arr)
	for _, result := range resultsB {
		fmt.Printf("x: %.2f, y: %.2f\n", result[0], result[1])
	}
}
