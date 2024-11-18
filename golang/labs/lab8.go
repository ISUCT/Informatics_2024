package labs

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func RunLab8() {
	var name string
	var age int
	var zodiacsign string
	fmt.Print("Введите имя: ")
	fmt.Fscan(os.Stdin, &name)
	fmt.Print("Введите возраст: ")
	fmt.Fscan(os.Stdin, &age)
	fmt.Print("Введите знак зодиака: ")
	fmt.Fscan(os.Stdin, &zodiacsign)

	file, err := os.Create("NewFile.txt")
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		return
	}
	defer file.Close()

	_, err = fmt.Fprintf(file, "%s %d %s\n", name, age, zodiacsign)
	if err != nil {
		fmt.Println("Ошибка при записи в файл:", err)
	} else {
		fmt.Println("Данные успешно записаны в файл NewFile.txt")
	}
	fmt.Println(name, age, zodiacsign)

	var searchString string
	fmt.Print("Введите слово для поиска: ")
	fmt.Scanln(&searchString)

	file, err = os.Open("NewFile.txt")
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	found := false
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(strings.ToLower(line), strings.ToLower(searchString)) {
			fmt.Println("Слово" + " " + searchString + "найдено!")
			found = true
			break
		}
	}
	if !found {
		fmt.Println("Слово" + " " + searchString + " " + "не найдено.")
	}
}
