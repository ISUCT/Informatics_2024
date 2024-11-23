package lab8

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func RunLab8() {
	RunLab8forLab4()
	var filename, strForFind string
	fmt.Println("Введите имя файла, в который вы хотите записать данные:")
	fmt.Scan(&filename)
	name, err := CreateFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Введите данные в файл (для завершения ввода введите пустую строку):")
	err = WriteDataToFile(name)
	if err != nil {
		fmt.Println(err)
	}
	slice, err := ReadDataFromFile(name)
	if err != nil {
		fmt.Println(err)
		fmt.Println("поиск данных в этом файле невозможен")
	} else {
		fmt.Printf("Чтение файла %s завершено\n", filename)
		for _, str := range slice {
			fmt.Println(str)
		}
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Введите строку, которую хотите найти в файле: ")
		strForFind, _ = reader.ReadString('\n')
		strForFind = strings.TrimRight(strForFind, "\r\n")
		_, n, err := IsStrInFile(name, strForFind)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Строка '%s' найдена на %d строке", strForFind, n+1)
		}
	}
}
