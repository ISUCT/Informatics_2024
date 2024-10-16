package fourlaba

import (
	"fmt"
	"math"
)

func CalculateY(a, b, x float64) float64 {
	logTerm := math.Log(b*b-x*x) / math.Log(a)
	denominator := math.Cbrt(math.Abs(x*x - a*a))
	return logTerm / denominator
}

// Функция для задачи A, с циклом от начального до конечного значения x с заданным шагом
func runTaskA(a, b, xStart, xEnd, step float64) {
	for x := xStart; x <= xEnd; x += step {
		y := CalculateY(a, b, x)
		fmt.Printf("x = %.2f, y = %.5f\n", x, y)
	}
}

// Функция для задачи B, с набором заранее заданных значений x
func runTaskB(a, b float64, xValues []float64) {
	for _, x := range xValues {
		y := CalculateY(a, b, x)
		fmt.Printf("x = %.2f, y = %.5f\n", x, y)
	}
}
