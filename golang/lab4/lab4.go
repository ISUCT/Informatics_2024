package Lab4

import "math"

func Сalcul(x float64) float64 {
	y := math.Pow((math.Pow(math.Asin(x), 2) + math.Pow(math.Acos(x), 4)), 3)
	return y
}
