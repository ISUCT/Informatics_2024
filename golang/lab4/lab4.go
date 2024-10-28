package lab4

import "math"

func Calculate(a, b, x float64) float64 {
	var y float64 = (1 + math.Pow(math.Log10(x/a), 2)) / (b - math.Pow(math.Exp(1), (x/a)))
	return y
}

func CompleteTaskA(a, b, xbegin, xend, xdelt float64) []float64 {
	var result []float64
	for  x := xbegin; x <= xend; x += xdelt {
		result = append(result, Calculate(a, b, x))
	}
	return result
}

func CompleteTaskB(a, b float64, x []float64) []float64 {
	var result []float64
	for i := 0; i < len(x); i++ {
		result = append(result, Calculate(a, b, x[i]))
	}
	return result
}
