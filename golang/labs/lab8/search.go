package lab8

import (
	"os"
	"io"
	"bufio"
	"fmt"
	"errors"
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

func searchingTextIntoFile () {
	var strFile string
	var strSearch string


	fmt.Print("Введите название файла, в котором будет поиск: ")
	fmt.Scanln(&strFile)
	file, err := os.OpenFile(strFile, os.O_RDONLY, 0444)
		if err != nil {
			if errors.Is(err, fs.ErrNotExist) {
				fmt.Print("Файла не существует, выбирите другой файл")
			} else{
				fmt.Print("Не удалось открыть файл, ошибка:", err)
				fmt.Print("\n")
			}
		}
	defer file.Close()

	fmt.Print("Введите искомый текст: ")
	fmt.Scanln(&strSearch)

	search(file, strSearch)
}