package lab8

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ReadDataFromFile(filename string) ([]float64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("ошибка открытия файла: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var floats []float64
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.ParseFloat(line, 64)
		if err != nil {
			return nil, fmt.Errorf("ошибка парсинга float '%s': %w", line, err)
		}
		floats = append(floats, num)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("ошибка чтения файла: %w", err)
	}
	return floats, nil
}

func GetParametrs(filename string) (a, b, xn, xk, deltax float64, otherValues []float64, err error) {
	data, err := ReadDataFromFile(filename)
	if err != nil {
		return 0, 0, 0, 0, 0, nil, fmt.Errorf("ошибка чтения данных из файла: %w", err)
	}
	if len(data) < 6 {
		return 0, 0, 0, 0, 0, nil, fmt.Errorf("нехватает значений! попробуй еще раз")
	}
	a = data[0]
	b = data[1]
	xn = data[2]
	xk = data[3]
	deltax = data[4]
	otherValues = data[5:]
	return
}
