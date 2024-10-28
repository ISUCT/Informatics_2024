package lab4

import (
	"fmt"
	"math"
)

func Calculate(a, b, x float64) float64 {
	return (1 + math.Pow(math.Log10(x/a), 2)) / (b - math.Pow(2.7, x/a))
}

func F_A(a, b, Xn, Xk, dX float64) {
	for x := Xn; x <= Xk; x += dX {
		fmt.Println(Calculate(b, a, x))
	}
}

func F_B(b, a float64, x [5]float64) {
	for _, value := range x {
		fmt.Println(Calculate(b, a, value))
	}
}