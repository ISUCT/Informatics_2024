package lab4

import "math"
import "fmt"

func Calculate(x float64)  float64{
	var y float64 = (math.Pow(1.2, x)) - (math.Pow(x, 1.2))
	
	return y
}
func TaskA(Xn, Xk, delX float64){
	for x := Xn; x <= Xk; x += delX {
		fmt.Println(Calculate(x))}
	}
func TaskB(x [5]float64) {
	for _, value := range x {
		fmt.Println(Calculate(value))
	}
}