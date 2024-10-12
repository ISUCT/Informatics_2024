package laba42

import (
	"fmt"
	"math"
)

func CalculateYYY() {
	a := 2.0
	xValues := []float64{1.16, 1.35, 1.48, 1.52, 1.96}
	for _, x := range xValues {
		logValue := math.Log10(a + x)
		y := math.Tan(math.Pow(logValue, 3)) / math.Pow((a+x), 2.0/7.0)
		fmt.Printf("При x = %.2f, y = %.6f\n", x, y)
	}
}
