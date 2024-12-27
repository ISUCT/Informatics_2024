package lab4

import (
	"fmt"
	"math"
)

func calculateY(a, b, x float64) float64 {
	var y float64 = (math.Pow(a, x) - math.Pow(b, x)) / (math.Log10(a/b) * math.Cbrt(a*b))
	return y
}

func Task_A(a, b, xn, xk, delx float64) []float64 {
	var yValues []float64
	for x := xn; x <= xk; x += delx {
		yValues = append(yValues, calculateY(a, b, x))
	}
	return yValues
}

func Task_B(a, b float64, xv []float64) []float64 {
	var yValues []float64
	for _, x := range xv {
		yValues = append(yValues, calculateY(a, b, x))
	}
	return yValues
}

func RunLab4() {
	fmt.Println("Task A:\n", Task_A(0.4, 0.8, 3.2, 6.2, 0.6))
	fmt.Println("Task B:\n", Task_B(0.4, 0.8, []float64{4.48, 3.56, 2.78, 5.28, 3.21}))
}
