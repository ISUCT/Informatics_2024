package Lab8

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func RunLab8() {
	Xn, Xk, deltaX := readInputFile("input.txt")

	resultsA, xValuesA := taskA(Xn, Xk, deltaX)
	printResults("Значения Y для диапазона:", xValuesA, resultsA)

	extraValues := readExtraValues("input.txt")
	resultsB := taskB(extraValues)
	printResults("Значения Y для дополнительных значений:", extraValues, resultsB)

	fileName := "output.txt"
	createFile(fileName)
	writeToFile(fileName, "Пример записи данных в файл")
	readFromFile(fileName)
	searchInFile(fileName)
}

func readInputFile(filename string) (float64, float64, float64) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Ошибка при открытии файла: %v\n", err)
		return 0, 0, 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var values []float64

	for scanner.Scan() {
		value, err := strconv.ParseFloat(scanner.Text(), 64)
		if err == nil {
			values = append(values, value)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Ошибка при чтении файла: %v\n", err)
	}

	if len(values) < 3 {
		fmt.Println("Ошибка: недостаточно данных в файле.")
		return 0, 0, 0
	}

	return values[0], values[1], values[2]
}

func readExtraValues(filename string) []float64 {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Ошибка при открытии файла: %v\n", err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var extraValues []float64
	lineCount := 0

	for scanner.Scan() {
		lineCount++
		if lineCount > 3 {
			value, err := strconv.ParseFloat(scanner.Text(), 64)
			if err == nil {
				extraValues = append(extraValues, value)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Ошибка при чтении файла: %v\n", err)
	}

	return extraValues
}

func createFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Ошибка при создании файла: %v\n", err)
		return
	}
	defer file.Close()
	fmt.Printf("Файл %s успешно создан.\n", filename)
}

func writeToFile(filename string, data string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Ошибка при открытии файла для записи: %v\n", err)
		return
	}
	defer file.Close()

	if _, err := file.WriteString(data + "\n"); err != nil {
		fmt.Printf("Ошибка при записи в файл: %v\n", err)
	}
}

func readFromFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Ошибка при открытии файла: %v\n", err)
		return
	}
	defer file.Close()

	fmt.Println("Содержимое файла:")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Ошибка при чтении файла: %v\n", err)
	}
}

func searchInFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Ошибка при открытии файла: %v\n", err)
		return
	}
	defer file.Close()

	fmt.Println("Введите текст для поиска в файле:")
	var searchText string
	fmt.Scanln(&searchText)

	scanner := bufio.NewScanner(file)
	found := false
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, searchText) {
			fmt.Printf("Найдено: %s\n", line)
			found = true
		}
	}
	if !found {
		fmt.Println("Текст не найден.")
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Ошибка при чтении файла: %v\n", err)
	}
}
