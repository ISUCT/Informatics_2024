package Lab4

import "math"

func Calculate(a float64, b float64, x float64) float64 {

	var y float64 = (math.Pow(math.Sqrt(math.Pow(x-a, 2)), 1.0/3.0) + math.Pow(math.Sqrt(math.Abs(x+b)), 1.0/5.0)) / math.Pow(math.Sqrt(math.Pow(x, 2)-math.Pow(a+b, 2)), 1.0/9.0)

	return y
}
