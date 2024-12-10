package lab8

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func RunLab8() {
	filename, err := CreateFile()
	if err != nil {
		fmt.Printf("Ошибка создания файла: %v\n", err)
		return
	}
	err = WriteToFile(filename)
	if err != nil {
		fmt.Printf("Ошибка записи в файл: %v\n", err)
		return
	}
	SearchWord(filename)
}

func WriteToFile(filename string) error {
	var info string
	fmt.Print("Введите информацию, которую хотите записать в файл: ")
	fmt.Scan(&info)
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("ошибка при открытии файла:%w", err)
	}
	defer file.Close()
	_, err = io.WriteString(file, info)
	fmt.Println(info)
	if err != nil {
		return fmt.Errorf("ошибка при записи в файл: %w", err)
	}
	return nil
}

func CreateFile() (string, error) {
	var filename string
	fmt.Print("Введите название файла: ")
	fmt.Scan(&filename)
	file, err := os.Create(filename)
	if err != nil {
		return "", fmt.Errorf("ошибка при создании файла: %w", err)
	}
	defer file.Close()
	return filename, nil
}

func SearchWord(file string) {
	var searchString string
	fmt.Print("Введите слово для поиска: ")
	fmt.Scanln(&searchString)

	f, err := os.Open(file)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	found := false
	for scanner.Scan() {
		line := scanner.Text()
		if line == searchString {
			found = true
			break
		}
	}
	if found {
		fmt.Println("Слово найдено")
	} else {
		fmt.Println("Слово не найдено")
	}
}

func ReadInput(filename string) (xn, xk, a, b float64, xValues []float64, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, 0, 0, 0, nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	values := make([]float64, 0)
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.ParseFloat(line, 64)
		if err != nil {
			return 0, 0, 0, 0, nil, fmt.Errorf("ошибка преобразования строки в число: %w", err)
		}
		values = append(values, num)
	}
	if len(values) < 10 {
		return 0, 0, 0, 0, nil, fmt.Errorf("в файле должно быть как минимум 10 чисел")
	}
	xn = values[0]
	xk = values[1]
	a = values[2]
	b = values[3]
	xValues = values[4:10]

	return xn, xk, a, b, xValues, nil
}
