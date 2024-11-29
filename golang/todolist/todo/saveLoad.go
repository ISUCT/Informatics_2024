package todo

import (
	"encoding/json"
	"os"
)

func (T *Todo) Save() error {
	data, err := json.Marshal(T.List)
	if err != nil {
		return err
	}

	file, err := os.Create("save.json")
	if err != nil {
		return err
	}
	defer file.Close()
	file.Write(data)
	return nil
}

func (T *Todo) Load() error {
	data, err := os.ReadFile("save.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &T.List)
	if err != nil {
		return err
	}
	return nil
}
