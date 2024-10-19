package lab

import (
	"fmt"
	"math"
)

func Calc(b, x float64) float64 {
	return (1 + math.Pow(math.Sin(math.Pow(b, 3)+math.Pow(x, 3)), 2)) / (math.Cbrt(math.Pow(b, 3) + math.Pow(x, 3)))
}

func TaskA(b, Xn, Xk, delX float64) []float64 {
	var arr []float64
	for x := Xn; x <= Xk; x += delX {
		arr = append(arr, Calc(b, x))
	}
	return arr
}

func TaskB(b float64, x []float64) []float64 {
	var arr []float64
	for _, value := range x {
		arr = append(arr, Calc(b, value))
	}
	return arr
}

func RunLab4Tasks(Values []float64) {
	for _, value := range Values {
		fmt.Println(value)
	}
}
