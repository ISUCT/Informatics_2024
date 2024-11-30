package labs

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RunLab8() {
	WriteToFile("NewFile.txt")
	SearchWord("Newfile.txt")
}

func WriteToFile(filename string) {
	var name string
	var age int
	var zodiacsign string
	fmt.Print("Введите имя: ")
	fmt.Fscan(os.Stdin, &name)
	fmt.Print("Введите возраст: ")
	fmt.Fscan(os.Stdin, &age)
	fmt.Print("Введите знак зодиака: ")
	fmt.Fscan(os.Stdin, &zodiacsign)

	file, err := os.Create(filename)
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
}

func SearchWord(file string) {
	var searchString string
	fmt.Print("Введите слово для поиска: ")
	fmt.Scanln(&searchString)

	f, err := os.Open(file)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	found := false
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(strings.ToLower(line), strings.ToLower(searchString)) {
			fmt.Println("Слово " + searchString + " найдено!")
			found = true
			break
		}
	}
	if !found {
		fmt.Println("Слово " + searchString + " не найдено.")
	}
}

func ReadInput(filename string) (a, b float64, xValues []float64, err error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return 0, 0, nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.ParseFloat(line, 64)
		if err != nil {
			return 0, 0, nil, fmt.Errorf("ошибка преобразования строки в число: %w", err)
		}
		if i == 0 {
			a = num
		} else if i == 1 {
			b = num
		} else {
			xValues = append(xValues, num)
		}
		i++
	}
	return a, b, xValues, nil
}
