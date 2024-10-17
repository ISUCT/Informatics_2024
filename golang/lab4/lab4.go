package lab4

import (
	"fmt"
	"math"
)

func RunLab4() {
	xn := 1.42
	xk := 3.62
	delx := 0.44
	x1 := 1.6
	x2 := 1.81
	x3 := 2.24
	x4 := 2.65
	x5 := 3.38
	//Задание А
	fmt.Println("знач у для хn-хk с шагом xk:")
	for x := xn; x <= xk; x += delx {
		fmt.Println(y(x))
	}
	//Задание В
	fmt.Println("знач у для х1-х5:", y(x1), ",", y(x2), ",", y(x3), ",", y(x4), ",", y(x5))
}

func y(x float64) float64 {
	a := 0.8
	b := 0.4
	return (math.Pow(x-a, 1.0/3.0) + math.Pow(x+b, 1.0/5.0)) / (math.Pow(x, 1.0/7.0) - math.Pow(x*x-(a+b)*(a+b), 1.0/9.0))
}
