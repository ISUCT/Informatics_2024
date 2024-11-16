package labs

import (
	"fmt"
	"os"
)

// останется сделать поиск текста в файле
func RunFileLab() {
	file, err := os.Create("text.txt")
	if err != nil {
		fmt.Println("Ошибка создания файла:", err)
		os.Exit(1)
	}
	defer file.Close()

	var name string
	var age int
	var city string

	fmt.Print("Введите имя: ")
	fmt.Fscan(os.Stdin, &name)
	fmt.Print("Введите возраст: ")
	fmt.Fscan(os.Stdin, &age)
	fmt.Print("Введите город: ")
	fmt.Fscan(os.Stdin, &city)

	_, err = fmt.Fprintf(file, "Имя: %s\nВозраст: %d\nГород: %s\n", name, age, city)
	if err != nil {
		fmt.Println("Ошибка записи в файл:", err)
		os.Exit(1)
	}

	fmt.Println("Данные успешно записаны в файл text.txt")

	filePath := "text.txt"
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		os.Exit(1)
	}
	fmt.Println("Содержимое файла:", string(fileData))
}
