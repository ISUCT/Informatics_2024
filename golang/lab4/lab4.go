package LAB_4

import (
	"fmt"
	"math"
)

func Calculate(x float64) float64 {
	return math.Pow(math.Abs(x*x-2.5), 0.25) + math.Pow(math.Log10(x*x), 0.33)
}
func TaskA(sx, ex, stepx float64) []float64 {
	var ValueY []float64
	for x := sx; x < ex; x += stepx {
		ValueY = append(ValueY, Calculate(x))
	}
	return ValueY
}
func TaskB(arguments []float64) []float64 {
	var ValueY []float64

	for _, x := range arguments {
		ValueY = append(ValueY, Calculate(x))
	}
	return ValueY
}

func RunLab4() {
	x := []float64{1.84, 2.71, 3.81, 4.56, 5.62}
	fmt.Println("задание a", TaskA(1.25, 3.25, 0.4))
	fmt.Println("задание b", TaskB(x))
}
