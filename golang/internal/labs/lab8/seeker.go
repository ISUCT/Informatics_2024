package lab8

import (
	"bufio"
	"os"
	"strings"
)

func SearchInFile(f *os.File, search string) ([]string, error) {
	defer f.Close()

	sc := bufio.NewScanner(f)
	result := []string{}
	for sc.Scan() {
		line := sc.Text()
		if strings.Contains(line, search) {
			result = append(result, line)
		}
	}
	return result, sc.Err()
}
