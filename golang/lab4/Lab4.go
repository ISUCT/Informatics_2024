package lab4

import (
	"fmt"
	"math"
)

func Calculate(x float64) float64 {
	y := math.Sqrt(math.Pow(math.Asin(x), 4) + math.Pow(math.Acos(x), 4))
	return y
}

func TaskA(Xn, Xk, delX float64) []float64 {
	results := []float64{}
	for x := Xn; x <= Xk; x += delX {
		results = append(results, Calculate(x))
	}
	return results
}

func TaskB(x [5]float64) []float64 {
	results := []float64{}
	for _, value := range x {
		results = append(results, Calculate(value))
	}
	return results
}

func RunLab4() {
	fmt.Println(TaskA(0.22, 0.92, 0.14))
	var s = [5]float64{0.1, 0.35, 0.4, 0.55, 0.6}
	fmt.Println(TaskB(s))
}
