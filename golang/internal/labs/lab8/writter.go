package lab8

import (
	"bufio"
	"fmt"
	"os"
)

func WriteToFile(file string) error {
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	fmt.Println("Введите текст (для завершения введите пустую строку): ")
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		line := sc.Text()
		if line == "" {
			break
		}
		_, err := f.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return sc.Err()
}
