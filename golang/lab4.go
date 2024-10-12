package main

import (
"fmt"
"math"
)

func calculateY(a, x float64) float64 {
// Вычисляем выражения
xSquaredMinus1 := x*x - 1
part1 := math.Pow(a, xSquaredMinus1) // a^(x^2 - 1)
part2 := math.Log(xSquaredMinus1) // ln(x^2 - 1)
part3 := math.Cbrt(xSquaredMinus1) // корень кубический (x^2 - 1)

// Итоговое значение y
y := part1 - part2 + part3
return y
}

func main() {
a := 1.6
x1 := 1.28

y := calculateY(a, x1)

fmt.Printf("Значение y для x1=%.2f: %.4f\n", x1, y)
}