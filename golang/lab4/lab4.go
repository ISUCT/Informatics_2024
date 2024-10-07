package lab4

import (
	"math"
)

func Function(x float64) float64 {
	var b float64 = 2.5
	y := (1 + math.Pow(math.Sin(math.Pow(b, 3)+math.Pow(x, 3)), 2)) / (math.Cbrt((math.Pow(b, 3) + math.Pow(x, 3))))
	return y
}
