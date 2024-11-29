package labs

import (
	"fmt"
	"math"
)

func CalculateFunction(x, a, b float64) float64 {
	y := (math.Log(a) * math.Abs(b*b-x*x)) / (math.Pow((math.Abs(x*x - a*a)), 0.2))
	return y
}

func TaskA(a, b, xMin, xMax, xDelta float64) []float64 {
	var y_a = []float64{}
	for i := xMin; i < xMax; i += xDelta {
		y_a = append(y_a, CalculateFunction(a, b, i))
	}
	return y_a
}

func TaskB(xs []float64, a, b float64) []float64 {
	var y_b = []float64{}
	for _, x_b := range xs {
		y_b = append(y_b, CalculateFunction(x_b, a, b))
	}
	return y_b
}

func Laba4run() {
	var y_a, y_b []float64
	y_a = TaskA(2.0, 1.1, 0.08, 1.08, 0.2)
	y_b = TaskB([]float64{0.1, 0.3, 0.4, 0.45, 0.65}, 2.0, 1.1)
	fmt.Println(y_a)
	fmt.Println(y_b)
}