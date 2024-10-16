package lab4

import "math"
import "fmt"

func Calculate(x float64)  float64{
	var y1 = return((math.Pow(1.2, x)) - (math.Pow(x, 1.2)))
	var y2 = return(math.Acos(x))
}
func TaskA(Xn, Xk, delX float64){
	for x := Xn; x <= Xk; x += delX {
		fmt.Println(Calculator(x))
	}
}