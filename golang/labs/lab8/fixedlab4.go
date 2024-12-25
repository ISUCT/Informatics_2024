package lab8

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"isuct.ru/informatics2022/labs/lab4"
)

var filePath string = "../golang/labs/lab8/"

func ReadFileFor4(filename string) ([]float64, error) {
	f, err := os.Open(filePath + filename)
	if err != nil {
		return nil, fmt.Errorf("Ошибка открытия файла")
	}
	fileScanner := bufio.NewScanner(f)
	var values []float64
	for fileScanner.Scan() {
		chislo, err := strconv.ParseFloat(fileScanner.Text(), 64)
		if err != nil {
			return nil, fmt.Errorf("Ошибка при переводе числа")
		}
		values = append(values, chislo)
	}
	return values, nil
}

func RunFixedLab4() {
	values, err := ReadFileFor4("input.txt")
	if err != nil {
		panic("Ошибка при чтении данных из файла")
	}
	var x = values[3:]
	var begin_x float64 = values[0]
	var end_x float64 = values[1]
	var delta_x float64 = values[2]
	fmt.Println("Задача A", lab4.Task_A(begin_x, end_x, delta_x))
	fmt.Println()
	fmt.Println("Задание B", lab4.Task_B(x))
}
