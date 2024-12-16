package lab4

import (
    "math"
    "fmt"
)

func logBase5(x float64) float64 {
    return math.Log(x) / math.Log(5)
}

func CalculateExpression(a float64, b float64, x float64) float64 {
    return (a*math.Cbrt(x) - b*logBase5(x)) / math.Pow(math.Log10(x-1), 3)
}

func TaskA(a, b, xi, xk, deltax float64) []float64 {
    var values []float64
    for x := xi; x <= xk; x += deltax {
        values = append(values, CalculateExpression(a, b, x))
    }
    return values
}

func TaskB(a, b float64, x [5]float64) []float64 {
    var values []float64
    for _, value := range x {
        values = append(values, CalculateExpression(a, b, value))
    }
    return values
}

func RunLab4Tasks() {
    a := 2.0
    b := 1.5
    fmt.Println(TaskA(a, b, 1.08, 1.88, 0.16))
    var s = [5]float64{1.16, 1.35, 1.48, 1.52, 1.96}
    fmt.Println(TaskB(a, b, s))
}
