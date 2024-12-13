package lab8

import (
	"fmt"
	"os"
)

func WriteFile(fileName string, newLine string) error {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return fmt.Errorf("не удается открыть файл")
	}
	defer file.Close()

	file.WriteString(newLine + "\n")
	return nil
}
