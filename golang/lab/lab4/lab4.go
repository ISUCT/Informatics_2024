package lab4_test

import (
 "math"
)

func CalculateY(a, b, x float64) float64 {
 return math.Pow((a*x+b), 1.0/3.0) / math.Pow(math.Log10(x), 2)
}

func TaskA(a, b, xStart, xEnd, step float64) []float64 {
 var result []float64
 for x := xStart; x <= xEnd; x += step {
  result = append(result, CalculateY(a, b, x))
 }
 return result
}

func TaskB(a, b float64, xValues []float64) []float64 {
 var result []float64
 for _, x := range xValues {
  result = append(result, CalculateY(a, b, x))
 }
 return result
}
