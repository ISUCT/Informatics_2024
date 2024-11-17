package lab8

import (
	"errors"
	"os"
)

func CreateFile(filename string) (string, error) {
	f, err := os.Create(filename)
	if err != nil {
		return "", errors.New("ошибка создания файла")
	}
	defer f.Close()
	return f.Name(), nil
}
