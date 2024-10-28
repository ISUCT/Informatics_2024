package lab4

import (
	"math"
)

func CalculateValue(a, b, x float64) float64 {
	if x == 0 {
		return 0
	}
	y := math.Sqrt(math.Abs(a-b*x) / math.Pow(math.Log10(x), 3))
	return y
}