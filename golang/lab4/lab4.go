package lab4

import (
	"fmt"
	"math"
)

func Calculation(x float64) float64 {
	return math.Pow(math.Sqrt(math.Asin(x)*math.Asin(x)*math.Asin(x)*math.Asin(x)+math.Acos(x)*math.Acos(x)*math.Acos(x)*math.Acos(x)), 1.0/7.0)
}

func Laba() {
	xH := 0.22
	xK := 0.92
	x_Delta := 0.14
	resultsA := calculateYValuesForRange(xH, xK, x_Delta)
	header := fmt.Sprintf("Значения y для x от %.2f до %.2f c шагом %.2f", xH, xK, x_Delta)
	printResults(header, xH, xK, resultsA)

	xs := []float64{0.1, 0.35, 0.4, 0.55, 0.6}
	resultsB := calculateYValuesForXList(xs)
	printResults(fmt.Sprintf("Значения y для x от %f до %f", xs[0], xs[len(xs)-1]), xs[0], xs[len(xs)-1], resultsB)
}

func calculateYValuesForRange(xH, xK, xDelta float64) []float64 {
	var results []float64
	for x := xH; x <= xK; x += xDelta {
		results = append(results, Calculation(x))
	}
	return results
}

func calculateYValuesForXList(xs []float64) []float64 {
	var results []float64
	for _, x := range xs {
		results = append(results, Calculation(x))
	}
	return results
}
func printResults(header string, startValue interface{}, endValue interface{}, results []float64) {
	fmt.Println(header, startValue, "до", endValue)
	fmt.Println("Результаты:")
	for _, result := range results {
		fmt.Printf("%v\n", result)
	}
}
