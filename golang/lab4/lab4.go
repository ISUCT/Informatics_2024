package lab4

import (
	"fmt"
	"math"
)

func lab4(a, x float64) float64 {
	return (math.Pow(a, math.Pow(x, 2)-1) - math.Log10(math.Pow(x, 2)-1) + math.Cbrt(math.Pow(x, 2)-1))
}
func lab4A(imin float64, imax float64, razn float64, a float64) []float64 {
	var result []float64
	for i := imin; i < imax; i += razn {
		result = append(result, lab4(a, i))
	}
	return (result)
}
func lab4B(x []float64, a float64) []float64 {
	var result []float64
	for _, value := range x {
		result = append(result, lab4(a, value))
	}
	return (result)
}
func Lab4AB() {
	var a float64 = 1.6
	var x = []float64{1.28, 1.36, 2.47, 3.68, 4.56}
	fmt.Println(lab4A(1.2, 3.7, 0.5, a))
	fmt.Println(lab4B(x, a))
}
