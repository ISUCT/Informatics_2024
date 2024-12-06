package lab8

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func SearchInFile(filename, search string) error {
	f, err := os.OpenFile(filename, os.O_RDONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	linenumber := 1
	found := false
	for sc.Scan() {
		line := sc.Text()
		if strings.Contains(line, search) {
			found = true
			fmt.Printf("найдено в строке %d: %s\n", linenumber, line)
		}
		linenumber++
	}

	if !found {
		fmt.Println("ничего не нашлось")
	}
	return sc.Err()
}
