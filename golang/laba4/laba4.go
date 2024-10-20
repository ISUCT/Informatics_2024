package lab4

import "math"

func Colcul(a, b, x float64) float64 {
	y := math.Pow((a+b*x)/(math.Pow(math.Log10(x), 3)), 1/5)
	return y

}
