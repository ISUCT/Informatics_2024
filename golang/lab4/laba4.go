package laba4
import "math"
func Lab4(a,b,x float64) float64{
	y := (math.Pow(a,x)-math.Pow(b,x)*math.Cbrt(a*b))/log10(a/b)
	return y
}