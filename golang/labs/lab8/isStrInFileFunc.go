package lab8

import (
	"errors"
	"fmt"
)

func IsStrInFile(filename, valueForSearch string) (bool, int, error) {
	if valueForSearch == "" {
		return false, 0, errors.New("невозможно найти пустое значение в файле")
	}
	values, err := ReadDataFromFile(filename)
	if err != nil {
		return false, 0, fmt.Errorf("ошибка чтения данных из файла %s", filename)
	}
	for i, value := range values {
		if value == valueForSearch { // string.Contains нужен для проверки вхрждения символов в слово, а не для сравнения. Так что другого варианта я не нашел
			return true, i, nil
		}
	}
	return false, 0, nil
}
