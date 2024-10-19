package labs

import (
	"fmt"
	"math"
)

func logarifm(x, base float64) float64 {
	return math.Log(x) / math.Log(base)
}

func calculateY(a, b, x float64) float64 {
	koren := math.Cbrt(x)
	log5 := logarifm(x, 5)
	logValue := math.Log(x - 1)
	lgkub := math.Pow(logValue, 3)
	return (a*koren - b*log5) / lgkub
}

func taskA(a, b float64) {
	fmt.Println("Задача A:")
	for x := 1.5; x <= 3.5; x += 0.4 {
		y := calculateY(a, b, x)
		fmt.Printf("x = %.2f, y = %.6f\n", x, y)
	}
}

func taskB(a, b float64) {
	fmt.Println("\nЗадача B:")
	xValues := []float64{1.9, 2.15, 2.34, 2.74, 3.16}
	for _, x := range xValues {
		y := calculateY(a, b, x)
		fmt.Printf("x = %.2f, y = %.6f\n", x, y)
	}
}

func runlab4() {
	var a float64 = 4.1
	var b float64 = 2.7

	taskA(a, b)
	taskB(a, b)
}

func main() {
	runlab4()
}
