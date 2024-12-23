package LAB_4

import (
	"fmt"
	"math"
)

const a = 2.0

func calculateY(x float64) float64 {
	y := math.Tan(math.Pow(math.Log10(a+x), 3)) / math.Pow(math.Pow((a+x), 2), 1/7)
	return y
}

func TaskA(xn, xk, delx float64) []float64 {
	var ValuesY []float64
	for x := xn; x <= xk; x += delx {
		ValuesY = append(ValuesY, calculateY(x))
	}
	return ValuesY
}

func TaskB(xv []float64) []float64 {
	var ValuesY []float64
	for _, x := range xv {
		ValuesY = append(ValuesY, calculateY(x))
	}
	return ValuesY
}

func Lab4() {
	xn, xk, delx := 1.08, 1.88, 0.16
	xv := []float64{1.16, 1.35, 1.48, 1.52, 1.96}
	fmt.Println("Задача А:", TaskA(xn, xk, delx))
	fmt.Println("Задача B:", TaskB(xv))
}
