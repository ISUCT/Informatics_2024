package lab

import (
	"fmt"
	"math"
	"io"
	"os"
	"strings"
	"strconv"
)

func Calculate(b, x float64) float64 {
	return (1 + math.Pow(math.Sin(math.Pow(b, 3)+math.Pow(x, 3)), 2)) / (math.Cbrt(math.Pow(b, 3) + math.Pow(x, 3)))
}

func NumbersTaskA() []string {
	var str_numbers []string
	file, err := os.Open("lab4A.txt")
	if err != nil{
        fmt.Println("Unable to create file:", err) 
        os.Exit(1)
	}

	data := make([]byte, 64)
	
	for{
        n, err := file.Read(data)
        if err == io.EOF{
            break
		}
		str_numbers = strings.Split(string(data[:n]), ",")
	}
	return str_numbers
}

func TaskA(b, Xn, Xk, delX float64) []float64 {
	var arr []float64
	for x := Xn; x <= Xk; x += delX {
		arr = append(arr, Calculate(b, x))
	}
	return arr
}

func NumbersTaskB() []string {
	var str_numbers []string
	file, err := os.Open("lab4B.txt")
	if err != nil{
        fmt.Println("Unable to create file:", err) 
        os.Exit(1)
	}

	data := make([]byte, 64)
	
	for{
        n, err := file.Read(data)
        if err == io.EOF{
            break
		}
		str_numbers = strings.Split(string(data[:n]), ",")
	}
	return str_numbers
}

func TaskB(b float64, x []float64) []float64 {
	var arr []float64
	for _, value := range x {
		arr = append(arr, Calculate(b, value))
	}
	return arr
}

func PrintValue (Values []float64) {
	for _, value := range Values {
		fmt.Println(value)
	}
}

func RunLab4Tasks() {
	RunLab8Tasks()
	const b float64 = 2.5
	
	var arr []string
	arr = NumbersTaskA()
	Xn, _ := strconv.ParseFloat(arr[0], 64)
	Xk, _ := strconv.ParseFloat(arr[1], 64)
	delX, _ := strconv.ParseFloat(arr[2], 64)
	ValuesA := TaskA(b, Xn, Xk, delX)

	arr = NumbersTaskB()
	var slice []float64
	for _, value := range arr {
		if value != "" {
		number, err := strconv.ParseFloat(value, 64)
		if err != nil {
			panic(err)
		}
		slice = append(slice, number)
		}
	}
	ValuesB := TaskB(b, slice)

	PrintValue(ValuesA)
	PrintValue(ValuesB)
}