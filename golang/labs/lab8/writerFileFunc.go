package lab8

import (
	"bufio"
	"errors"
	"os"
)

func WriteDataToFile(filename string) error {
	var firstLine bool = true
	f, err := os.OpenFile(filename, os.O_WRONLY, 0600)
	if err != nil {
		return errors.New("ошибка открытия файла")
	}
	defer f.Close()
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		if firstLine == bool(true) {
			firstLine = false
			continue
		}
		txt := sc.Text()
		if txt == "" {
			return nil
		}
		_, err := f.WriteString(txt)
		f.WriteString("\n")
		if err != nil {
			return errors.New("ошибка записи в файл")
		}
	}
	return nil
}
