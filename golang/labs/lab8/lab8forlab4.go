package lab8

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"isuct.ru/informatics2022/labs/lab4"
)

func ReadDataFromFileForLab4(filename string) ([]float64, error) {
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

func RunLab8forLab4() {
	values, err := ReadDataFromFileForLab4("input.txt")
	if err != nil {
		panic("ошибка при чтении данных из файла")
	}
	var a float64 = values[0]
	var b float64 = values[1]
	var taskBslice = values[5:]
	fmt.Println("Задача А")
	var xn float64 = values[2]
	var xk float64 = values[3]
	var deltax float64 = values[4]
	for _, pair := range lab4.TaskA(a, b, xn, xk, deltax) {
		x := pair[0]
		y := pair[1]
		fmt.Printf("x = %.2f\ty = %f\n", x, y)
	}
	fmt.Println()
	fmt.Println("Задание B")
	for _, pair := range lab4.TaskB(a, b, taskBslice) {
		x := pair[0]
		y := pair[1]
		fmt.Printf("x = %.2f\ty = %f\n", x, y)
	}
}
