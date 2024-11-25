package calculate

import (
	"math"
)

func CalculateY(a, b, x float64) float64 {
	numerator := math.Pow(a, 1.0/3.0) + math.Pow(math.Tan(b*x), 4.5)
	denominator := math.Pow(b, 1.0/5.0) + 1/math.Pow(math.Tan(a*x), 2.7)
	return numerator / denominator
}
