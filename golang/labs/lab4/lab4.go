package lab4

import (
	"fmt"
	"math"
)

func Lab4A(xn, xk, dx float64) float64 {
	fmt.Println("\nЗадача А")
	for x := xn; x <= xk; x += dx {
		fmt.Printf("y(%f) = %f\n", x, Lab4Y(x))
	}
	return 0
}

func Lab4B(xv []float64) float64 {
	fmt.Println("Задача B")
	for i := 0; i<len(xv); i++ {
		fmt.Printf("y(%f) = %f\n", xv[i], Lab4Y(xv[i]))
	}
	return 0
}

func Lab4Y(x float64) float64 {
	var DegreeAsin, DegreeAcos,DegreeNumerator, DegreeDenumenator float64 = 4, 6, 1, 7
	return (math.Pow(math.Pow(math.Asin(x), DegreeAsin)+math.Pow(math.Acos(x), DegreeAcos), DegreeNumerator/DegreeDenumenator))
}

func Lab4() {
	var xn, xk, dx float64 = 0.22, 0.92, 0.14
	var xv []float64 = []float64{0.1, 0.35, 0.4, 0.55, 0.6}
	Lab4A(xn, xk, dx)
	Lab4B(xv)
}
