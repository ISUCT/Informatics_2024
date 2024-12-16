package lab8

import (
	"os"
	"io"
	"bufio"
	"fmt"
)

func search(file *os.File, search string) {
	var searchCopy []rune = []rune(search)
	var count int
	var countRune int

	reader := bufio.NewReader(file) 
	
    for { 
        line, err := reader.ReadString('\n') 
		var lineUse []rune = []rune(line)
        if err != nil { 
            if err == io.EOF { 
                break
            } else { 
                fmt.Println(err) 
                return
            } 
        } 

        for _, lineRune := range(lineUse) {
			for _, searchRune := range(searchCopy) {
				if searchRune == lineRune {
					countRune++
						if countRune == len([]rune(search)) {
							countRune = 0
							count++
							searchCopy = []rune(search)
						} else {
							searchCopy = append(searchCopy[:0], searchCopy[1:]...)
							break
						}	
				} else {
					countRune = 0
				}
			}
		}
    }
	fmt.Println("Количество найденных слов: ", count)
}