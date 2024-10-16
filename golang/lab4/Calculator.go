package lab4

import (
	"math"
	"fmt"
)

func Calc(a, b, x float64) float64 {
	var y float64 = math.Asin(math.Pow(x,a)) + math.Acos(math.Pow(x,b))
		return y
}

func TaskA(b, Xn,  Xk, delX float64) {
	for x := Xn; x <= Xk; x += delX {
		fmt.Println(Calc(b, x))
	}
}

func TaskB(b float64, x [5]float64) {
	for _, value := range x {
		fmt.Println(Calc(b, value))
	}
}