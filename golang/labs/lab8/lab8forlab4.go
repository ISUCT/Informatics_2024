package lab8

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ReadDataFromFileForLab4(filename string) ([]float64, error) {
	f, err := os.OpenFile(fmt.Sprintf("D:/Informatics/Informatics_2024/golang/labs/lab8/%s", filename), os.O_RDONLY, 0600)
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
