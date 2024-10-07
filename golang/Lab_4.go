package main

import (
	"fmt"
	"math"
)

func Calculator(a, b, x float64) float64 {
	return (1 + math.Pow(math.Log10(x/a), 2)) / (b - math.Pow(2.7, x/a))
}

func TaskA(a, b, Xn, Xk, delX float64) {
	for x := Xn; x <= Xk; x += delX {
		fmt.Printf("x: %.2f, Result: %.4f\n", x, Calculator(a, b, x))
	}
}

func TaskB(a, b float64, arr []float64) {
	for _, value := range arr {
		fmt.Printf("x: %.2f, Result: %.4f\n", value, Calculator(a, b, value))
	}
}

func main() {
	TaskA(2.0, 5.0, 1.0, 10.0, 1.0)

	values := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	TaskB(2.0, 5.0, values)
}
