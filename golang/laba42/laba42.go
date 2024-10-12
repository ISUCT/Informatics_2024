package laba42

import (
	"fmt"
	"math"
)

func CalculateB() {
	a := 2.0
	xValues := []float64{1.16, 1.35, 1.48, 1.52, 1.96}
	for _, x := range xValues {
		logValue := math.Log10(a + x)
		y := math.Tan(math.Pow(logValue, 3)) / math.Pow((a+x), 2.0/7.0)
		fmt.Printf("При x = %.2f, y = %.6f\n", x, y)
	}
}
func CalculateA() {
	a := 2.0
	xn := 1.08
	xk := 1.88
	deltaX := 0.16
	for x := xn; x <= xk; x += deltaX {
		logValue := math.Log10(a + x)
		y := math.Tan(math.Pow(logValue, 3)) / math.Pow((a+x), 2.0/7.0)
		fmt.Printf("При x = %.2f, y = %.6f\n", x, y)
	}
}
