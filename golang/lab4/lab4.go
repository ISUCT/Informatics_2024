package lab4

import (
	"fmt"
	"math"
)

func lab4(a, x float64) float64 {
	return (math.Pow(a, math.Pow(x, 2)-1) - math.Log10(math.Pow(x, 2)-1) + math.Cbrt(math.Pow(x, 2)-1))
}
func lab4A(imin float64, imax float64, razn float64, a float64) {
	for i := imin; i < imax; i += razn {
		fmt.Println(lab4(a, i))
	}
}
func lab4B(x []float64, a float64) {
	for _, value := range x {
		fmt.Println(lab4(a, value))
	}
}
func Lab4AB() {
	var a float64 = 1.6
	var x = []float64{1.28, 1.36, 2.47, 3.68, 4.56}
	lab4A(1.2, 3.7, 0.5, a)
	lab4B(x, a)
}
