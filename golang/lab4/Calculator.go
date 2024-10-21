package lab4

import (
	"fmt"
	"math"
)

func Calculator(a, b, x float64) float64 {
	return (1 + math.Pow(math.Log10(x/a), 2)) / (b - math.Pow(2.7, x/a))
}

func TaskA(a, b, Xn, Xk, delX float64) {
	for x := Xn; x <= Xk; x += delX {
		fmt.Println(Calculator(b, a, x))
	}
}

func TaskB(b, a float64, x [5]float64) {
	for _, value := range x {
		fmt.Println(Calculator(b, a, value))
	}
}
