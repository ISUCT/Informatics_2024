package lab4

import (
    "math"
)

func Calc(a float64, x float64) float64 {
    return math.Tan(math.Pow(math.Log10(a+x), 3)) / math.Pow(a+x, 2.0/7.0)
}

func TaskA(a float64, xi float64, xk float64, deltaX float64) []float64 {
    values := []float64{}
    for x := xi; x <= xk; x += deltaX {
        values = append(values, Calc(a, x))
    }
    return values
}

func TaskB(a float64, x [5]float64) []float64 {
    values := []float64{}
    for _, value := range x {
        values = append(values, Calc(a, value))
    }
    return values
}
