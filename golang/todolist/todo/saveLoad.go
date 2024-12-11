package todo

import (
	"encoding/json"
	"fmt"
	"os"
)

const link = "todolist/save.json"

func (T *Todo) Save() error {
	data, err := json.Marshal(T.List)
	if err != nil {
		return fmt.Errorf("сериализация: %w", err)
	}

	file, err := os.Create(link)
	if err != nil {
		return fmt.Errorf("создание файла сохронения %s: %w", link, err)
	}
	defer file.Close()
	file.Write(data)
	return nil
}

func (T *Todo) Load() error {
	data, err := os.ReadFile(link)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return fmt.Errorf("чтение файла %s: %w", link, err)
	}

	err = json.Unmarshal(data, &T.List)
	if err != nil {
		return fmt.Errorf("десериализация: %w", err)
	}
	return nil
}
