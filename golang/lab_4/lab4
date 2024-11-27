package lab4

import (
	"fmt"
	"math"
)

func Calculate(x, a, b float64) float64 {
	y := (a + math.Pow((math.Tan(b*x)), 2)) / (b + math.Pow(math.Cos(a*x)/math.Sin(a*x), 2))

	return y
}

func TaskA(a, b, Xn, Xk, delX float64) []float64 {
	var y []float64
	for x := Xn; x <= Xk; x += delX {
		y = append(y, Calculate(a, b, x))
	}
	return y
}

func TaskB(a float64, b float64, x [5]float64) []float64 {
	var y []float64
	for _, value := range x {
		y = append(y, Calculate(a, b, value))
	}
	return y
}

func RunLab4() {
	a := 0.1
	b := 0.5
	fmt.Println(TaskA(a, b, 0.15, 1.37, 0.25))
	var s = [5]float64{0.2, 0.3, 0.44, 0.6, 0.56}
	fmt.Println(TaskB(a, b, s))
}
