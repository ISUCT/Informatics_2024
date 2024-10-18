package Lab4

import (
	"math"
)

func СalculateY(x float64) float64 {
	y := math.Pow((math.Pow(math.Asin(x), 2) + math.Pow(math.Acos(x), 4)), 3)
	return y
}

func TaskA(xn, xk, xdel float64) []float64 {
	var yValues []float64
	for x := xn; x <= xk; x += xdel {
		yValues = append(yValues, СalculateY(x))
	}
	return yValues
}

func TaskB(xv []float64) []float64 {
	var yValues []float64
	for _, x := range xv {
		yValues = append(yValues, СalculateY(x))
	}
	return yValues

}
