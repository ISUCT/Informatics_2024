package lab4

import (
	"fmt"
	"math"
)

func Calculator(x, a, b float64) float64 {
	y := math.Pow(a*x+b, 1.0/3.0) / math.Pow(math.Log(x), 2.0)
	return y
}

func TaskA(a, b, Xn, Xk, delX float64) {
	for x := Xn; x <= Xk; x += delX {
		fmt.Println(Calculator(a, b, x))
	}
}

func TaskB(a, b float64, x [5]float64) {
	for _, value := range x {
		fmt.Println(Calculator(a, b, value))
	}
}
