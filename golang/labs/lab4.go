package lab4

import (
  "fmt"
  "math"
)

// Функция для вычисления
func f(x float64, a float64, b float64) float64 {
  return math.Acos(math.Pow(x, 2)-math.Pow(b, 2)) / math.Asin(math.Pow(x, 2)-math.Pow(a, 2))
}

// Функция для выполнения задания
func RunLab4() {
  // Задание А
  fmt.Println("Задание А")
  fmt.Println("x\ty")
  a := 0.05
  b := 0.06
  xn := 0.2
  xk := 0.95
  deltaX := 0.15
  for x := xn; x <= xk; x += deltaX {
    fmt.Printf("%.2f\t%.2f\n", x, f(x, a, b))
  }

  // Задание B
  fmt.Println("Задание B")
  fmt.Println("x\ty")
  iks := []float64{0.15, 0.26, 0.37, 0.48, 0.56}
  for _, x := range iks {
    fmt.Printf("%.2f\t%.2f\n", x, f(x, a, b))
  }
}
