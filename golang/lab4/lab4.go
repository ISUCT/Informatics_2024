package lab4

import (
	"fmt"
	"math"
)

func Function(x float64) float64 {
	var b float64 = 2.5
	y := (1 + math.Pow(math.Sin(math.Pow(b, 3)+math.Pow(x, 3)), 2)) / (math.Cbrt((math.Pow(b, 3) + math.Pow(x, 3))))
	s := math.Floor(y*100) / 100
	return s
}
func RunLab4() {
	fmt.Print("Задача А\n")
	var x_a float64 = 1.28
	for i := 0; i < 6; i++ {
		fmt.Print(Function(x_a), "\n")
		x_a = x_a + 0.4
	}

	fmt.Print("Задача В\n")
	xs := [5]float64{1.1, 2.4, 3.6, 1.7, 3.9}
	for _, x_b := range xs {
		fmt.Print(Function(x_b), "\n")
	}
}
