package labs

import (
	"fmt"
	"math"
)

func Calculate_y(x float64, a float64, b float64) float64 {
	y := math.Acos(math.Pow(x, 2)-math.Pow(b, 2)) / math.Asin(math.Pow(x, 2)-math.Pow(a, 2))
	return y
}


func Task_a(xn, xk, deltax float64) [][]float64 {
	var taskA [][]float64
	for x := xn; x <= xk; x += deltax {
		y := Calculate_y(x, 0.05, 0.06)
		taskA = append(taskA, []float64{x, y})
	}
	return taskA
}


func Task_b(values []float64) [][]float64 {
	var taskB [][]float64
	for _, x := range values {
		y := Calculate_y(x, 0.05, 0.06)
		taskB = append(taskB, []float64{x, y})
	}
	return taskB
}

func RunLab4() {
	resultsA := Task_a(0.2, 0.95, 0.15)
	for _, result := range resultsA {
		fmt.Printf("x: %.2f, y: %.2f\n", result[0], result[1])
	}

	arr := []float64{0.15, 0.26, 0.37, 0.48, 0.56}
	resultsB := Task_b(arr)
	for _, result := range resultsB {
		fmt.Printf("x: %.2f, y: %.2f\n", result[0], result[1])
	}
}