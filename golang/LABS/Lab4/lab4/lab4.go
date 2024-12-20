package lab4

import (
	"fmt"
	"math"
)

func Calculate(x float64) float64 {
	return math.Pow(math.Abs(x*x-2.5), 0.25) + math.Pow(math.Log10(x*x), 0.33333333)
}

func Task_A(begin_x, end_x, delta_x float64) []float64 {
	var result []float64

	for x := begin_x; x < end_x; x += delta_x {
		result = append(result, Calculate(x))
	}
	return result
}

func Task_B(arguments []float64) []float64 {
	var result []float64

	for _, x := range arguments {
		result = append(result, Calculate(x))
	}
	return result
}

func RunLab4() {
	x := []float64{1.84, 2.71, 3.81, 4.56, 5.62}
	fmt.Println("Задача A", Task_A(1.25, 3.25, 0.4))
	fmt.Println("Задача B", Task_B(x))
}
