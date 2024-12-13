package lab8

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CreateFile(filename string) (string, error) {
	if _, err := os.Stat(filename); err == nil {
		return "", fmt.Errorf("файл %s уже существует", filename)
	} else if !os.IsNotExist(err) {
		return "", fmt.Errorf("ошибка проверки существования файла: %w", err)
	}
	file, err := os.Create(filename)
	if err != nil {
		return "", fmt.Errorf("ошибка создания файла: %w", err)
	}
	defer file.Close()
	return filename, nil
}
<<<<<<< Updated upstream
func WriteToFile(filepath string, data string) error {
=======

func WriteToFile(filepath string, name string, age int, city string, university string, additionalText string) error {
>>>>>>> Stashed changes
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("ошибка открытия файла для записи: %w", err)
	}
	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		return fmt.Errorf("ошибка записи в файл: %w", err)
	}
	return nil
}

func ReadFile(filepath string) (string, error) {
	fileData, err := os.ReadFile(filepath)
	if err != nil {
		return "", fmt.Errorf("ошибка чтения файла: %w", err)
	}
	return string(fileData), nil
}

func SearchInFile(filedata string, searchText string) bool {
	return strings.Contains(filedata, searchText)
}
<<<<<<< Updated upstream
func CollectUserData() string {
	var name string
	var age int
	var city string
	var university string
	var additionalText string

	fmt.Print("Введите имя: ")
	fmt.Fscan(os.Stdin, &name)
	fmt.Print("Введите возраст: ")
	fmt.Fscan(os.Stdin, &age)
	fmt.Print("Введите город: ")
	fmt.Fscan(os.Stdin, &city)
	fmt.Print("Введите институт, в котором вы учитесь: ")
	fmt.Fscan(os.Stdin, &university)
	fmt.Print("Введите любой текст: ")
	fmt.Scanln()
	fmt.Scanln(&additionalText)
	return fmt.Sprintf("Имя: %s\nВозраст: %d\nГород: %s\nУниверситет: %s\nДополнительный текст: %s\n", name, age, city, university, additionalText)
}
func RunFileLab() {
	filePath, err := CreateFile()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	data := CollectUserData()
	err = WriteToFile(filePath, data)
=======

func GetUserInput() (string, int, string, string, string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите имя: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Введите возраст: ")
	var age int
	fmt.Scanf("%d\n", &age)

	fmt.Print("Введите город: ")
	city, _ := reader.ReadString('\n')
	city = strings.TrimSpace(city)

	fmt.Print("Введите институт, в котором вы учитесь: ")
	university, _ := reader.ReadString('\n')
	university = strings.TrimSpace(university)

	fmt.Print("Введите любой текст: ")
	additionalText, _ := reader.ReadString('\n')
	additionalText = strings.TrimSpace(additionalText)

	return name, age, city, university, additionalText
}

func RunFileLab() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите имя файла: ")
	filename, _ := reader.ReadString('\n')
	filename = strings.TrimSpace(filename)

	filePath, err := CreateFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	name, age, city, university, additionalText := GetUserInput()

	err = WriteToFile(filePath, name, age, city, university, additionalText)
>>>>>>> Stashed changes
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Данные успешно записаны в файл", filename)

	fileData, err := ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Содержимое файла:", fileData)

	var searchText string
	fmt.Print("Введите текст для поиска: ")
	searchText, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	searchText = strings.TrimSpace(searchText)

	if SearchInFile(fileData, searchText) {
		fmt.Println("Текст найден в файле")
	} else {
		fmt.Println("Текст не найден в файле")
	}
}
