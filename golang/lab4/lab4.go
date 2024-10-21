package lab4
import (
	"math"
)
func Calculate(a, b, elements float64) float64 {
	y := math.Pow((a+b*elements)/(math.Pow(math.Log10(elements), 3)), 0.2)
	return y
}

func TaskA(a, b, xn, xk, xdel float64) []float64{
	var res []float64
	for i := xn; i < xk; i += xdel {
		res = append(res, Calculate(a,b,i))
	}
	return res
}

func TaskB(a, b float64, x []float64) []float64 {
	var res []float64
	for _, i := range x {
		res = append(res, Calculate(a, b, i))
	}
	return res
}
