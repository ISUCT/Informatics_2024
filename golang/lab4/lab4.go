package lab4

import (
	"math"
)

func Calculator(a, b, x float64) float64 {
	y := math.Cbrt(a*x+b) / (math.Log(x) * math.Log(x))

	return y
}
