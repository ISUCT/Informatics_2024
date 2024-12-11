package todo

import (
	"fmt"
	"time"
)

type Task struct {
	Name       string    `json:"Name"`
	DateCreate time.Time `json:"DateStart"`
	Deadline   time.Time `json:"Deadline"`
	Teg        string    `json:"Teg"`
}

type Todo struct {
	List []Task `json:"List"`
}

func (T *Todo) AddTask(newName string, newTime string, newTeg string) error {
	newDeadline, err := time.Parse("2006-01-02", newTime)
	if err != nil {
		return fmt.Errorf("установка даты дедлайна при создании Task: %w", err)
	}
	if len(T.List) == 0 {
		T.List = append(T.List,
			Task{
				Name:       newName,
				DateCreate: time.Now().UTC(),
				Deadline:   newDeadline,
				Teg:        newTeg,
			},
		)
	} else {
		T.List = append(
			[]Task{
				{
					Name:       newName,
					DateCreate: time.Now().UTC(),
					Deadline:   newDeadline,
					Teg:        newTeg,
				},
			},
			T.List...,
		)
	}
	return nil
}
