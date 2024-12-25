package lab8

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"isuct.ru/informatics2022/labs/lab4"
)

var filePath string = "../golang/labs/lab8/"

func ReadDataFromLab4(filename string) ([]float64, error) {
	f, err := os.Open(filePath + filename)
	fmt.Println(err)
	if err != nil {
		return nil, fmt.Errorf("ошибка открытия файла")
	}
	fileScanner := bufio.NewScanner(f)
	var values []float64
	for fileScanner.Scan() {
		chislo, err := strconv.ParseFloat(fileScanner.Text(), 64)
		if err != nil {
			return nil, fmt.Errorf("ошибка при переводе числа")
		}
		values = append(values, chislo)
	}
	return values, nil
}

func RunLab8WithLab4() {
	values, err := ReadDataFromLab4("input.txt")
	if err != nil {
		fmt.Printf("Ошибка при чтении данных: %v\n", err)
		return
	}

	a := values[0]
	b := values[1]
	xn := values[2]
	xk := values[3]
	dx := values[4]
	arguments := values[5:]

	fmt.Print("A: ")
	xResultsA, yResultsA := lab4.TaskA(a, b, xn, xk, dx)
	fmt.Println("X:", xResultsA)
	fmt.Println("Y:", yResultsA)

	fmt.Print("B: ")
	yResultsB := lab4.TaskB(a, b, arguments)
	fmt.Println("Y:", yResultsB)
}
