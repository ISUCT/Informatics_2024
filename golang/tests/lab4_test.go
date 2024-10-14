package tests

import (
	"fmt"
	"math"
	"testing"
)

func y(x float64, a float64, b float64) float64 {
	return math.Pow(a+b*x, 2.5)/1 + math.Log10(a+b*x)
}

func TestY(t *testing.T) {
	a := 2.01
	b := 2.02
	xvalues := []float64{1.52, 5.46, 2.33}
	for _, x := range xvalues {
		result := y(x, a, b)
		if result < 0 {
			fmt.Println("Ожидается положительное значение")
		} else {
			fmt.Println("ок")
		}
	}
}
