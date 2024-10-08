package Lab4

import (
	"fmt"
	"math"
)

func calculate(x float64) float64 {
	n := 2.0 // Пример значения n
	return math.Pow(math.Pow(math.Asin(x), 4)+math.Pow(math.Acos(x), 6), n)
}

func Task_A() {
	for x := 0.22; x <= 0.92; x += 0.14 {
		result := calculate(x)
		fmt.Printf("результат от %.2f : %.2f\n", x, result)
	}
}

func Task_B() {
	var arr = [5]float64{0.1, 0.35, 0.4, 0.55, 0.6}
	for _, x := range arr {
		answer := calculate(x)
		fmt.Println(answer)
	}
}
