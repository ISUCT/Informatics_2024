package lab4

import (
	"math"
	"fmt"
)

func Calculator(a, b, x float64) float64{
	y := math.Cbrt(a*x + b)/(math.Log(x)*math.Log(x))
	return y
}

func TaskA(a, b, Xn, Xk, delX float64) {
	for x := Xn; x <= Xk; x += delX {
		fmt.Println(Calculator(a, b, x))
	}
}

func TaskB(a, b float64, x[5]float64) {
	for _, value := range x {
		fmt.Println(Calculator(a, b, value))

	
	}
}

