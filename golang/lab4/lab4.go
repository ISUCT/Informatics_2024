package lab4

import (
	"fmt"
	"math"
)

func Calculate(a, b, x float64) float64 {
	tngBx := math.Tan(b * x)
	cotgAx := 1 / math.Tan(a*x)
	y := a + (tngBx*tngBx)/b + (cotgAx * cotgAx)
	return y
}

func T_B(a float64, b float64, x [5]float64) []float64 {
	var y []float64
	for _, value := range x {
		y = append(y, Calculate(a, b, value))
	}
	return y
}

func T_A(a, b, Xn, Xk, delX float64) []float64 {
	var y []float64
	for x := Xn; x <= Xk; x += delX {
		y = append(y, Calculate(a, b, x))
	}
	return y
}

func RunLab4() {
	var a float64 = 0.1
	var b float64 = 0.5
	fmt.Println(T_A(a, b, 0.15, 1.37, 0.25))
	var values = [5]float64{0.2, 0.3, 0.44, 0.6, 0.56}
	fmt.Println(T_B(a, b, values))
}
