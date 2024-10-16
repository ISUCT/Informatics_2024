package labs
//9 вариант
import (
	"fmt"
	"math"
)

func calculateY(x float64) float64 {
	var a float64 = 1.1
	var b float64 = 0.09
	numerator := math.Log10(x*x - 1)
	denominator := math.Log(a*x*x-b) / math.Log(5)
	return numerator / denominator
}
func Laba4A() []float64 {
	fmt.Print("Задача A\n")
	var x_a float64 = 1.2
	results := []float64{}
	for i := 1.2; i < 2.2; i += 0.2 {
		y := calculateY(x_a)
		results = append(results, y)
		x_a += 0.2
	}
	return results
}
func Laba4B() (float64) {
	fmt.Print("Задача B\n")
	xn := [5]float64{1.21, 1.76, 2.53, 3.48, 4.52}
	var x_b float64
	for _, x_b := range xn {
		fmt.Print(calculateY(x_b), "\n")
	}
	return x_b
}