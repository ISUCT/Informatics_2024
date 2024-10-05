package function

import (
	"math"
)

func Uravnenie (b, x float64) float64 {
	var b3 float64 = math.Pow(b, 3)
	var x3 float64 = math.Pow(x, 3)

	return (1+math.Pow(math.Sin(b3 + x3), 2))/math.Pow(b3 + x3, 1/3)
}