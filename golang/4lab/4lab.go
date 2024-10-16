package fourlaba

import (
	"math"
)

func CalculateY(a, b, x float64) float64 {
	numeratorPart := b*b - x*x

	logTerm := math.Log(numeratorPart) / math.Log(a)

	denominator := math.Cbrt(math.Abs(x*x - a*a))

	y := logTerm / denominator

	return y
}
