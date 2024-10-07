package lab4

import (
	"math"
	"fmt"
)

func Calculator(x, a float64) float64 {
	y := math.Pow(a, math.Sqrt(x)-1) - math.Log10(math.Sqrt(x) - 1) + math.Pow(math.Sqrt(x)-1, 1.0/3.0)
	return y
}

func TaskA(b, Xn,  Xk, delX float64) {
	for x := Xn; x <= Xk; x += delX {
		fmt.Println(Calculator(b, x))
	}
}

func TaskB(b float64, x [5]float64) {
	for _, value := range x {
		fmt.Println(Calculator(b, value))
	}
}