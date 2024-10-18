package lab4

import (
	"math"
)

func Colcul(a, b, x float64) float64 {
	y := (math.Log10(math.Pow(x, 2) - 1)) / (((math.Log(a * math.Pow(x, 2))) / (math.Log(5))) - b)
	return y
}
