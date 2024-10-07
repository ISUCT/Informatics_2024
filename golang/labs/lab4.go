package labs

import (
	"fmt"
	"math"
)

func RunLab4() {
	xn := 1.1
	xk := 3.6
	deltax := 0.5
	x1 := 1.2
	x2 := 1.28
	x3 := 1.36
	x4 := 1.46
	x5 := 2.35
	//Задание А
	fmt.Println("значения функции у для хn-хk с шагом дельта х:")
	for x := xn; x <= xk; x += deltax {
		fmt.Println(y(x))
	}
	//Задание В
	fmt.Println("значения функции у для х1-х5:", y(x1), ",", y(x2), ",", y(x3), ",", y(x4), ",", y(x5))
}

func y(x float64) float64 {
	a := 2.5
	b := 4.6
	return math.Pow(a+b*x, 2.5)/1 + math.Log10(a+b*x)
}
