package lab4

import (
	"fmt"
	"math"
)

func calculateY(x float64) (float64, error) {
	if math.Abs(x) >= 1 {
		return math.Pow(1.2, x) - math.Pow(x, 1.2), nil
	} else if math.Abs(x) < 1 {
		return math.Acos(x), nil
	} else {
		return 0, fmt.Errorf("неверный x")
	}
}

func Lab4() {
	xValues := []float64{0.2, 0.6, 1, 1.4, 1.8, 2.2}

	for _, x := range xValues {
		y, err := calculateY(x)
		if err != nil {
			fmt.Printf("Ошибка при x = %f: %v\n", x, err)
		} else {
			fmt.Printf("При x = %f, y = %f\n", x, y)
		}
	}
}
