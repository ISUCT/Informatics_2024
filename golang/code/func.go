package func_

import "math"

func Lab_4 (a, x float64) float64 {
	return math.Pow(math.Log(a+x),2)/math.Pow(a+x, 2)
}