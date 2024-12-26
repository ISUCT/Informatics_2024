package lab9

import (
	"encoding/json"
	"fmt"
	"os"
)

func CreateFile(path string) {
	file, err_create := os.Create(path)

	if err_create != nil {
		fmt.Println(err_create)
	}

	defer file.Close()

}

func WriteFile(path string, list *ToDoList) {
	file, errOpen := os.OpenFile(path, os.O_RDWR, 0666)

	if errOpen != nil {
		fmt.Println(errOpen)
	}

	Data, errData := json.Marshal(&list)

	if errData != nil {
		fmt.Println(errData)
	}

	file.WriteString(string(Data))
	defer file.Close()
}
