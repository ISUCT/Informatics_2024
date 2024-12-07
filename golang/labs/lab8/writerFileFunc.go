package lab8

import (
	"bufio"
	"fmt"
	"os"
)

func WriteDataToFile(filename string) error {
	f, err := os.OpenFile(filePath+filename, os.O_WRONLY, 0600)
	if err != nil {
		return fmt.Errorf("открытия файла: %w", err)
	}
	defer f.Close()
	sc := bufio.NewScanner(os.Stdin)
	var firstLine bool = true
	for sc.Scan() {
		if firstLine == bool(true) { // scan выводит первую строку пустой по умолчанию
			firstLine = false
			continue
		}
		txt := sc.Text()
		if txt == "q" {
			return nil
		}
		_, err := f.WriteString(txt)
		f.WriteString("\n")
		if err != nil {
			return fmt.Errorf("запись в файл: %w", err)
		}
	}
	return nil
}
