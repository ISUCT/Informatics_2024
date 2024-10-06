package lab4

import (
	"math"
)

func F(x float64) float64 {
	var a float64 = 0.1
	var b float64 = 0.5
	return (math.Cbrt(a) + math.Pow(math.Tan(b*x), 4.5)) / (math.Pow(b, 0.2) + (1 / (math.Pow(math.Tan(a*x), 2.7))))
}
