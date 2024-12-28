package laba8

import (
	"bufio"
	"os"
)

func ReadDataFromFile(filename string) []string {
	var textSlice []string
	f, _ := os.OpenFile(filePath+filename, os.O_RDONLY, 0600)
	defer f.Close()
	fileScanner := bufio.NewScanner(f)
	for fileScanner.Scan() {
		textSlice = append(textSlice, fileScanner.Text())
	}

	return textSlice
}
