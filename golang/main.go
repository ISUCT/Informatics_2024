package main

import (
"fmt"

"isuct.ru/informatics2022/lab4"
)

  func main() {
  var a float64 = -2.5
  var x []float64 = []float64{2.89, 3.54, 5.21, 6.28, 3.48}
  var xn float64 = 3.5
  var xk float64 = 6.5
  var xdel float64 = 0.6
  var resA []float64 = lab4.TaskA(a, xn, xk, xdel)
  fmt.Println("Задача А", resA)
  var resB []float64 = lab4.TaskB(a, x)
  fmt.Println("Задача В", resB)
  fmt.Println("Гоголев Александр Сергеевич")
}
