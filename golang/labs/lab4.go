package labs

import (
	"fmt"
	"math"
)

func Euation(x,a,b float64) float64 {
	return math.Acos(math.Pow(x, 2)-math.Pow(b, 2)) / math.Asin(math.Pow(x, 2)-math.Pow(a, 2))
}


func prinvValues(xn,xk,deltaX,a,b float64){
	fmt.Println("x\ty")
	for x := xn; x <= xk; x += deltaX {
		fmt.Printf("%.2f\t%.2f\n", x, Euation(x, a, b))
	}
}

func printValuesArray(iks []float64, a,b float64){
	fmt.Println("x\ty")
	for _, x := range iks {
		fmt.Printf("%.2f\t%.2f\n", x, Euation(x, a, b))
	}
}
func RunLab4() {
	a := 0.05
	b := 0.06
	prinvValues(0.2, 0.95, 0.15, a, b)

	

	iks := []float64{0.15, 0.26, 0.37, 0.48, 0.56}
	printValuesArray(iks,a,b)
}
