package lab4

import (
	"fmt"
	"math"
)

func Calculate(a, b, elements float64) float64 {
	y := math.Pow((a+b*elements)/(math.Pow(math.Log10(elements), 3)), 0.2)
	return y
}

func TaskA(a, b, xn, xk, xdel float64) []float64 {
	var res []float64
	for i := xn; i < xk; i += xdel {
		res = append(res, Calculate(a, b, i))
	}
	return res
}

func TaskB(a, b float64, x []float64) []float64 {
	var res []float64
	for _, i := range x {
		res = append(res, Calculate(a, b, i))
	}
	return res
}

func Lab4() {
	var a float64 = 7.2
	var b float64 = 1.3
	var x []float64 = []float64{2.4, 2.8, 3.9, 4.7, 3.16}
	var xn float64 = 1.56
	var xk float64 = 4.71
	var xdel float64 = 0.63

	var resA []float64 = TaskA(a, b, xn, xk, xdel)
	fmt.Println("Задача А", resA)
	var resB []float64 = TaskB(a, b, x)
	fmt.Println("Задача В", resB)
}
