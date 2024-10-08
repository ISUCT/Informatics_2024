package lab4

import (
	"math"
	"fmt"
)

const a float64 = 1.1
const b float64 = 0.09

func Calculation(a,b,x float64) float64 {
	y := ((math.Log10(math.Pow(x,2)-1))/(math.Log(a*math.Pow(x,2)-b)/math.Log(5)))
	return y
}

func TaskA(a, b, x1, x2, dx float64) {
	for x := x1; x<=x2; x += dx {
		fmt.Println(Calculation(a,b,x))
	}
}

func TaskB(a float64, b float64, x[5] float64) {
	for _, value := range x {
		fmt.Println(Calculation(a,b,value))
	}
}