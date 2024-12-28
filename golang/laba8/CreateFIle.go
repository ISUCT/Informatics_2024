package laba8

import (
	"os"
)

func CreateFile(filename string) (string, error) {
	_, _ = os.Stat(filePath + filename)
	f, _ := os.Create(filePath + filename)

	defer f.Close()
	return filename, nil
}
