package lab

import (
	"fmt"
	"math"
)

func Calculate(b, x float64) float64 {
	return (1 + math.Pow(math.Sin(math.Pow(b, 3)+math.Pow(x, 3)), 2)) / (math.Cbrt(math.Pow(b, 3) + math.Pow(x, 3)))
}

func TaskA(b, Xn, Xk, delX float64) []float64 {
	var arr []float64
	for x := Xn; x <= Xk; x += delX {
		arr = append(arr, Calculate(b, x))
	}
	return arr
}

func TaskB(b float64, x []float64) []float64 {
	var arr []float64
	for _, value := range x {
		arr = append(arr, Calculate(b, value))
	}
	return arr
}

func PrintValue (Values []float64) {
	for _, value := range Values {
		fmt.Println(value)
	}
}

func RunLab4Tasks() {
	const b float64 = 2.5
	var slice = []float64{1.1, 2.4, 3.6, 1.7, 3.9}
	ValuesA := TaskA(b, 1.28, 3.28, 0.4)
	ValuesB := TaskB(b, slice)
	PrintValue(ValuesA)
	PrintValue(ValuesB)
}