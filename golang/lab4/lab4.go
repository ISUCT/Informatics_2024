package lab4

import (
	"fmt"
	"math"
)

func Calculate(a, b, x float64) float64 {
	y := math.Cbrt(a*x+b) / (math.Log(x) * math.Log(x))

	return y
}

func Enter(a, b float64, arr []float64) {
	for index, value := range arr {
		y := Calculate(a, b, value)
		fmt.Printf("X%d = %g\n", index, y)
	}
}

func TaskA(x_f, x_end, step float64) []float64 {
	list := []float64{}

	for i := x_f; i < x_end+step; i += step {
		list = append(list, i)
	}
	return list
}
