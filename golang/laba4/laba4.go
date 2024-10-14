package laba4

import (
	"fmt"
	"math"
)

func Colcul(a, b, x float64) float64 {
	y := (math.Pow(math.Sin(a+b*x), 3.5)) / (1 + math.Cos(math.Log10(a+b*x)))
	return y
}

func TaskA(a, b, xMin, xMax, xDelta float64) {
	var result float64
	for i := xMin; i < xMax; i += xDelta {
		result = Colcul(a, b, i)
		fmt.Printf("при x = %.2f, y = %.4f \n", i, result)
	}
}

func TaskB(a, b float64, x []float64) {
	var result float64
	for i := 0; i < len(x); i++ {
		result = Colcul(a, b, x[i])
		fmt.Printf("при x = %.2f, y = %.4f \n", x[i], result)
	}
}
