package fileutilis

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func WriteToFile(filename string, info string) error {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return fmt.Errorf("ошибка при открытии файла:%w", err)
	}
	defer file.Close()
	_, err = io.WriteString(file, info)
	if err != nil {
		return fmt.Errorf("ошибка при записи в файл: %w", err)
	}
	return nil
}

func CreateFile(filename string) (string, error) {
	file, err := os.Create(filename)
	if err != nil {
		return "", fmt.Errorf("ошибка при создании файла: %w", err)
	}
	defer file.Close()
	return filename, nil
}

func ReadFromFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("ошибка при открытии файла для чтения: %w", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var fileContent strings.Builder
	for scanner.Scan() {
		fileContent.WriteString(scanner.Text() + "\n")
	}
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("ошибка при чтении файла: %w", err)
	}
	return fileContent.String(), nil
}

func SearchWord(filename string, searchString string) (string, error) {

	f, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("ошибка при открытии файла: %w", err)
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
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("ошибка при чтении файла: %w", err)
	}
	if found {
		return "Слово найдено", nil
	} else {
		return "Слово не найдено", nil
	}
}

func ReadInput(filename string) (a, b float64, xValues []float64, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, 0, nil, fmt.Errorf("ошибка открытия файла: %w", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	values := []float64{}
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.ParseFloat(line, 64)
		if err != nil {
			return 0, 0, nil, fmt.Errorf("ошибка преобразования строки '%s' в число: %w", line, err)
		}
		values = append(values, num)
	}

	if len(values) < 10 {
		return 0, 0, nil, fmt.Errorf("недостаточно чисел")
	}

	a = values[0]
	b = values[1]
	xValues = values[2:]

	return a, b, xValues, nil
}
