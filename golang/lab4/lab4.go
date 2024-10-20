gipackage lab4

import (
	"math"
)

func Calculate(a, b, x float64) float64 {
	var y float64 = (1 + math.Pow(math.Log10(x/a), 2)) / (b - math.Pow(math.Exp(1), (x/a)))
	return y
}

func A(a, b, xb, xe, xd float64) []float64 {
	var y []float64
	for x := xb; x <= xe; x += xd {
		y = append(y, Calculate(a, b, x))
	}
	return y
}

func B(a, b float64, x []float64) []float64 {
	var t []float64
	for i := 0; i < len(x); i++ {
		t = append(t, Calculate(a, b, x[i]))
	}
	return t
}
