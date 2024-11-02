package lab4

import (
	"math"
)

func Calculate(x float64) float64 {
	var n float64 = 1.0 / 7
	return math.Pow(math.Pow(math.Asin(x), 4)+math.Pow(math.Acos(x), 6), n)
}

func Task_A(beginx float64, endx float64, stepx float64) []float64 {
	answers := []float64{}
	for x := beginx; x <= endx; x += stepx {
		result := Calculate(x)
		answers = append(answers, result)
	}
	return answers
}

func Task_B(arr []float64) []float64 {
	answers := []float64{}
	for _, x := range arr {
		answers = append(answers, Calculate(x))
	}
	return answers
}
