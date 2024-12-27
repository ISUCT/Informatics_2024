package lab4

import (
	"fmt"
	"math"
)

func CalculateY(a, elements float64) float64 {
	numerator := math.Pow(math.Log10(math.Pow(a, 2)+elements), 2)
	denominator := (a + elements) * (a + elements)
	y := numerator / denominator
	return y
}
func TaskA(a, xn, xk, xdel float64) []float64 {
	var res []float64
	for i := xn; i < xk; i += xdel {
		res = append(res, CalculateY(a, i))
	}
	return res
}

func TaskB(a float64, x []float64) []float64 {
	var res []float64
	for _, i := range x {
		res = append(res, CalculateY(a, i))
	}
	return res
}

func RunLab4() {
	var a float64 = -2.5
	var x []float64 = []float64{2.89, 3.54, 5.21, 6.28, 3.48}
	var xn float64 = 3.5
	var xk float64 = 6.5
	var xdel float64 = 0.6
	var resA []float64 = TaskA(a, xn, xk, xdel)
	fmt.Println("Задача А", resA)
	var resB []float64 = TaskB(a, x)
	fmt.Println("Задача В", resB)
}