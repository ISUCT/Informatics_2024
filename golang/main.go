package main

import (
 "fmt"
 "isuct.ru/informatics2022/lab4"
)

func main() {
 var result float64
 var a float64 = 7.2
 var b float64 = 4.2
 var x [5]float64 = [5]float64{2.4, 2.8, 3.9, 4.7, 3.16}

 fmt.Println("\nзадача А:")
 for i := 1.81; i < 5.31; i += 0.7 {
  result = lab4.Colcul(a, b, i)
  fmt.Println("при x=", i, "y=", result)
 }

 fmt.Println("\nзадача B:")
 for _, i := range x {
  result = lab4.Colcul(a, b, i)
  fmt.Println("при x=", i, "y=", result)
 }
}