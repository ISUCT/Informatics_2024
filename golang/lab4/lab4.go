package lab4

import (
	"math"
	"fmt"
)

func Calculate (a float64, b float64, x float64) float64 {
	return (math.Sin(math.Pow((a+b*x), 3.5))) / (1 + math.Cos(math.Log10(a+b*x)))
}

func TaskA(a, b, xn, xk, delx float64) []float64 {
	var values []float64
	for x := xn; x <= xk; x += delx {
		values = append(values, Calculate (a, b, x))
	}
	return values
}

func TaskB(a float64, b float64, x []float64) []float64 {
	var values []float64
	for _, value := range x {
		values = append(values, Calculate (a, b, value))
	}
	return values
}

func RunLab4Tasks() {
	a := 2.5
	b := 4.6
	fmt.Println(TaskA(a, b, 1.15, 3.05, 0.38))
		var s = []float64{1.2, 1.36, 1.57, 1.93, 1.25}
		fmt.Println(TaskB(a,b,s))
}

