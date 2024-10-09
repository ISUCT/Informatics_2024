package lab4

import (
	"math"
	"fmt"
)

func Calculator(a, x float64) float64{
	y := math.Tan(math.Pow(math.Log10(a+x),3)/math.Pow(a+x,2/7))
	return y
}

func TaskA(a, Xn, Xk, delX float64) {
	for x := Xn; x <= Xk; x += delX {
		fmt.Println(Calculator(a, x))
	}
}

func TaskB(a float64, x[5]float64) {
	for _, value := range x {
		fmt.Println(Calculator(a, value))

	
	}
}

