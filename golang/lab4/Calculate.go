package Lab4

import (
	"math"
	"fmt"
)

func Calculate(a,b,x float64) float64 {
	y := ((math.Log10(math.Pow(x,2)-1))/(math.Log(a*math.Pow(x,2)-b)/math.Log(5)))
	return y
}

func TaskA(a, b, x1, x2, dx float64) []float64 {
	var y []float64
	for x := x1; x<=x2; x += dx {
		y = append(y, Calculate(a,b,x))
	}
	return y
}

func TaskB(a float64, b float64, x[5] float64) []float64 {
	var y []float64
	for _, value := range x {
		y = append(y, Calculate(a,b,value))
	}
	return y
}

func RunLab4 () {
	a := 1.1
	b := 0.09
	
	fmt.Println(TaskA(a,b,1.2,2.2,0.2))
	var s = [5]float64{1.21,1.76,2.53,3.48,4.52}
	fmt.Println(TaskB(a, b, s))
}
