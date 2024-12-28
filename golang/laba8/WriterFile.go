package laba8

import (
	"fmt"
	"os"
)

func WriteDataToFile(filename string) {
	f, _ := os.OpenFile(filePath+filename, os.O_WRONLY, 0600)
	defer f.Close()
	var name string
	fmt.Scan(&name)
	var surname string
	fmt.Scan(&surname)

	f.WriteString(name)
	f.WriteString(surname)
}
