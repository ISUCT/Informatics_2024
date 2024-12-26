package lab4

import (
	"fmt"
	"math"
)

const a float64 = 1.35
const b float64 = 0.98

func Calculate(a, b, x float64) float64 {
	y := math.Cbrt(a*x+b) / (math.Log(x) * math.Log(x))

	return y
}

func Enter(a, b float64, arr []float64) {
	for index, value := range arr {
		y := Calculate(a, b, value)
		fmt.Printf("X%d = %g\n", index, y)
	}
}

func TaskA(x_f, x_end, step float64) []float64 {
	list := []float64{}

	for i := x_f; i < x_end+step; i += step {
		list = append(list, i)
	}
	return list
}

func Show_lab4() {
	A := TaskA(1.14, 4.24, 0.62)
	var B []float64 = []float64{0.35, 1.28, 3.51, 5.21, 4.16}

	fmt.Println("Задача A:")
	Enter(a, b, A)

	fmt.Println("Задача B:")
	Enter(a, b, B)
}
