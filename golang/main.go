package main

import (
    "fmt"
    "isuct.ru/informatics2022/lab4"
)

const a float64 = 2.0

func main() {
    fmt.Println("Рукина Полина Владимировна")

    valuesA := lab4.TaskA(a, 1.08, 1.88, 0.16)
    var slice = [5]float64{1.16, 1.35, 1.48, 1.52, 1.96}
    valuesB := lab4.TaskB(a, slice)

    for _, value := range valuesA {
        fmt.Println(value)
    }
    
    for _, value := range valuesB {
        fmt.Println(value)
    }
}
