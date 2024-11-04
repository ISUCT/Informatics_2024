package LAB_4

import (
	"fmt"
	"math"
)

func CalculateY(a, b, x float64) float64 {
	numerator := math.Pow(a, 1.0/3.0) + math.Pow(math.Tan(b*x), 4.5)
	denominator := math.Pow(b, 1.0/5.0) + 1/math.Pow(math.Tan(a*x), 2.7)
	return numerator / denominator
}

func Lab4() {
	var a, b, xStart, xEnd, step float64

	a = 0.1
	b = 0.5
	xStart = 0.33
	xEnd = 1.23
	step = 0.18

	fmt.Printf("Task A:\n")
	for x := xStart; x <= xEnd; x += step {
		y := CalculateY(a, b, x)
		fmt.Printf("For x = %.2f, y = %.10f\n", x, y)
	}

	fmt.Printf("Task B:\n")
	xValues := []float64{0.5, 0.36, 0.40, 0.62, 0.78}

	for _, x := range xValues {
		y := CalculateY(a, b, x)
		fmt.Printf("For x = %.2f, y = %.10f\n", x, y)
	}
}
