package lab4

import "math"

func Lab4(a, x float64) float64 {
	y := math.Pow(a, math.Pow(x, 2)-1) - math.Log10(math.Pow(x, 2)-1) + math.Cbrt(math.Pow(x, 2)-1)
	return (y)
}
