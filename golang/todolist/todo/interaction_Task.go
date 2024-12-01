package todo

import (
	"errors"
	"fmt"
	"time"
)

var errNonExistentElement error = errors.New("errNonExistentElement")

func (T *Todo) RemoveTask(num int) error {
	if num >= len(T.List) {
		return fmt.Errorf("удаление задачи: %w", errNonExistentElement)
	}
	T.List = append(T.List[:num], T.List[(num+1):]...)
	return nil
}

func (T *Todo) MoveTask(from int, to int) error {
	if from >= len(T.List) || to >= len(T.List) {
		return fmt.Errorf("перемещение задачи: %w", errNonExistentElement)
	}
	list := make([]Task, from)
	copy(list, T.List[:from])

	list = append(list, T.List[(from+1):to]...)
	list = append(list, T.List[from])
	list = append(list, T.List[to:]...)

	T.List = list
	return nil
}

func (T *Todo) EditTask(num int, newName string, newTime string, newTeg string) error {
	newDeadline, err := time.Parse("2006-01-02", newTime)
	if err != nil {
		return fmt.Errorf("установка даты дедлайна при изменении имени: %w", err)
	}

	T.List[num].Name = newName
	T.List[num].Deadline = newDeadline
	T.List[num].Teg = newTeg
	return nil
}

func (T *Todo) EditName(num int, newName string) {
	T.List[num].Name = newName
}

func (T *Todo) EditTime(num int, newTime string) error {
	newDeadline, err := time.Parse("2006-01-02", newTime)
	if err != nil {
		return err
	}

	T.List[num].Deadline = newDeadline
	return nil
}

func (T *Todo) EditTeg(num int, newTeg string) {
	T.List[num].Teg = newTeg
}
