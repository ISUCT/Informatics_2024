package lab4

import (
	"math"
)

func Mathfunc(a, b, elements float64) float64 {
	y := math.Pow((a+b*elements)/(math.Pow(math.Log10(elements), 3)), 0.2)
	return y
}
