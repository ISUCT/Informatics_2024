package lab4

import (
	"fmt"
	"math"
)

func Calculate(x float64) float64 {
	if math.Abs(x) >= 1 {
		return (math.Pow(1.2, x)) - (math.Pow(x, 1.2))
	}
	return math.Acos(x)
}

func TaskA(Xmin, Xmax, Xdel float64) []float64 {
	var y []float64
	for x := Xmin; x <= Xmax; x += Xdel {
		y = append(y, Calculate(x))
	}
	return y
}

func TaskB(x [5]float64) []float64 {
	var y []float64
	for _, value := range x {
		y = append(y, Calculate(value))
	}
	return y
}

func Runlab4() {
	fmt.Println(TaskA(0.2, 2.2, 0.4))
	fmt.Println(TaskB([5]float64{0.1, 0.9, 1.2, 1.5, 1.3}))
}
