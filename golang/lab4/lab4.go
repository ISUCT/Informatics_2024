package lab4

import (
	 "fmt"
	 "math"
)

func Calc(a float64,x float64) float64 {
	return math.Tan (math.Pow(math.Log10(a+x),3))/math.Pow(a+x,2/7)
}
func TaskA( a float64,xi float64, xk float64,deltaX float64) {
	for x:=xi;x<=xk;x +=deltaX{
		fmt.Println (Calc(a,x))
	}
}
func TaskB(a float64, x [5] float64){
	for _, value := range x{
		fmt.Println (Calc(a,value))
	}
}
