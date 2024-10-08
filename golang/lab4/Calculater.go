package lab4 //Грубо говоря ваша созданная папка

import (
	"fmt"
	"math"
)

func Calculator(x float64) float64 {
	//Снизу ваша формула
	y := math.Sqrt(math.Pow(math.Asin(x), 4) + math.Pow(math.Acos(x), 4))
	return y
}

func TaskA(Xn, Xk, delX float64) {
	for x := Xn; x <= Xk; x += delX {
		fmt.Println(Calculator(x))
	}
}

func TaskB(x [5]float64) {
	for _, value := range x {
		fmt.Println(Calculator(value))
	}
}
