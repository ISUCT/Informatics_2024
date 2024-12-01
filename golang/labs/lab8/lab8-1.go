package lab8

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Parameters struct {
	A      float64
	B      float64
	Xn     float64
	Xk     float64
	Deltax float64
	Others []float64
}

func ReadDataFromFile(filename string) (Parameters, error) {
	file, err := os.Open(filename)
	if err != nil {
		return Parameters{}, fmt.Errorf("ошибка открытия файла: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var values []float64
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.ParseFloat(line, 64)
		if err != nil {
			return Parameters{}, fmt.Errorf("ошибка парсинга float '%s': %w", line, err)
		}
		values = append(values, num)
	}

	if err := scanner.Err(); err != nil {
		return Parameters{}, fmt.Errorf("ошибка чтения файла: %w", err)
	}
	if len(values) < 5 {
		return Parameters{}, fmt.Errorf("нехватает значений! попробуй еще раз")
	}
	params := Parameters{
		A:      values[0],
		B:      values[1],
		Xn:     values[2],
		Xk:     values[3],
		Deltax: values[4],
		Others: values[5:],
	}
	return params, nil
}

func GetParametrs(filename string) (a, b, xn, xk, deltax float64, otherValues []float64, err error) {
	data, err := ReadDataFromFile(filename)
	if err != nil {
		return 0, 0, 0, 0, 0, nil, fmt.Errorf("ошибка чтения данных из файла: %w", err)
	}

	a = data.A
	b = data.B
	xn = data.Xn
	xk = data.Xk
	deltax = data.Deltax
	otherValues = data.Others

	return
}
