package Lab4

import (
	"fmt"
	"math"
)

func Calculate(x, a float64) float64 {
	y := math.Pow(a, math.Sqrt(x)-1) - math.Log10(math.Sqrt(x)-1) + math.Pow(math.Sqrt(x)-1, 1.0/3.0)
	return y
}

func TaskA(b, Xn, Xk, delX float64) []float64 {
	var y []float64
	for x := Xn; x <= Xk; x += delX {
		y = append(y, Calculate(b, x))
	}
	return y
}

func TaskB(b float64, x [5]float64) []float64 {
	var y []float64
	for _, value := range x {
		y = append(y, Calculate(b, value))
	}
	return y
}

func RunLab4() {
	b := 2.5
	fmt.Println(TaskA(b, 1.2, 3.7, 0.5))
	var s = [5]float64{1.28, 1.36, 2.47, 3.68, 4.56}
	fmt.Println(TaskB(b, s))
}
