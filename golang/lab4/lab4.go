package lab4

import (
	"fmt"
	"math"
)

func Lab424() {
	a := 4.1
	b := 2.7

	xStart := 1.5
	xEnd := 3.5
	xDelta := 0.4

	xValues := []float64{1.9, 2.34, 2.34, 2.74, 3.16}

	for x := xStart; x <= xEnd; x += xDelta {
		y := (math.Pow(a, 3)*math.Sqrt(x) - b*math.Log(x)) / (math.Pow(math.Log10(x-1), 3))
		fmt.Printf("При x = %.2f, y = %.2f\n", x, y)
	}

	for _, x := range xValues {
		y := (math.Pow(a, 3)*math.Sqrt(x) - b*math.Log(x)) / (math.Pow(math.Log10(x-1), 3))
		fmt.Printf("При x = %.2f, y = %.2f\n", x, y)
	}
}
