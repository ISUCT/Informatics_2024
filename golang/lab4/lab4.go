package lab4

import (
	"fmt"
	"math"
)

func CalculateFunction(x, b float64) float64 {
	y := (1 + math.Pow(math.Sin(math.Pow(b, 3)+math.Pow(x, 3)), 2)) / (math.Cbrt((math.Pow(b, 3) + math.Pow(x, 3))))
	return y
}
func TaskA(xa_beginning, xa_end, xa_delta float64) []float64 {
	y_a := []float64{}
	for i := xa_beginning; i <= xa_end; i = i + xa_delta {
		y_a = append(y_a, CalculateFunction(i, 2.5))
	}
	return y_a
}

func TaskB(x1, x2, x3, x4, x5 float64) []float64 {
	y_b := []float64{}
	xs := [5]float64{x1, x2, x3, x4, x5}
	for _, x_b := range xs {
		y_b = append(y_b, CalculateFunction(x_b, 2.5))
	}
	return y_b
}

func RunLab4() {
	fmt.Print("Задача А\n")
	fmt.Print(TaskA(1.28, 3.28, 0.4), "\n")
	fmt.Print("Задача В\n")
	fmt.Print(TaskB(1.1, 2.4, 3.6, 1.7, 3.9))
}
