package lab4

import (
	"math"
	"fmt"
)

func Calculate(a, b, x float64) float64 {
	y := math.Cbrt(a*x+b) / (math.Log(x) * math.Log(x))

	return y
}


func Enter(a, b float64, arr []float64){
	for index, value := range arr {
		// fmt.Printf(index)
		y := Calculate(a, b, value)
		fmt.Printf("X%d = %g\n", index, y)
	}
}