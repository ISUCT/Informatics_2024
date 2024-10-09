package laba4

import "math"

func Colcul(a, b, x float64) float64 {
	y := (math.Log(math.Pow(b, 2)-math.Pow(x, 2)) / math.Log(a)) / (math.Pow(x*x-a*a, 1/3))
	return y
}
