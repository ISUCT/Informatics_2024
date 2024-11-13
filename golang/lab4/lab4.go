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
func  CalculateExpression (a float64, x float64) float64 {
    return math.Tan(math.Pow(math.Log10(a+x), 3)) / math.Pow(a+x, 2.0/7.0)
}

func TaskA(a , xi , xk , deltax float64) {
     var values []float64
    for x := xi; x <= xk; x += deltaX {
        values = append(values,  CalculateExpression (a, x))
    }
    return values
}

func TaskB(a float64, x [5]float64) []float64 {
    values := []float64{}
    for _, value := range x {
        values = append(values, Calc(a, value))
    }
    return values
    var values []float64{}
    for _, value := range x {
        values = append(values,  CalculateExpression (a, value))
    }
    return values
}
func RunLab4Tasks() {
    const a float64 = 2.0
    valuesA := TaskA(a, 1.08, 1.88, 0.16)
    var slice = [5]float64{1.16, 1.35, 1.48, 1.52, 1.96}
    valuesB := TaskB(a, slice)

    fmt.Println("Результаты TaskA:")
    for _, value := range valuesA {
        fmt.Println(value)
    }

    fmt.Println("Результаты TaskB:")
    for _, value := range valuesB {
        fmt.Println(value)
    }
}
