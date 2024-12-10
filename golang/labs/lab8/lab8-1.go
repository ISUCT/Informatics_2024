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
	Slicex []float64
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
		Slicex: values[5:],
	}
	return params, nil
}
