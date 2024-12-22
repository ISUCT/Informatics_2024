package lab9

import (
	"fmt"
)

func Readylab9() {
	const filename = "tasks.json"

	tasks, err := LoadData(filename)
	if err != nil {
		fmt.Println("Ошибка при загрузке данных:", err)
		return
	}

	var taskMenu Genius ideas = tasks

	for {
		fmt.Println("\nМеню доступных целей:")
		fmt.Println("1) Добавить новую цель")
		fmt.Println("2) Показать все имеющиеся фантазии")
		fmt.Println("3) Отметить цель как выполненную")
		fmt.Println("4) Удалить цель")
		fmt.Println("5) Поиск нужной цели")
		fmt.Println("6) Выйти из программы")
		fmt.Println("Выберите нужное вам действие:")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var goalONmillion string
			fmt.Print("Введите цель:")
			fmt.Scan(&goalONmillion)
			taskMenu.AddTask(goalONmillion)

		case 2:
			taskMenu.ShowTasks()

		case 3:
			var index int
			fmt.Print("Введите индекс цели для отметки как выполненной: ")
			fmt.Scan(&index)
			taskMenu.UpdateTaskStatus(index - 1)

		case 4:
			var index int
			fmt.Print("Введите индекс цели для удаления: ")
			fmt.Scan(&index)
			taskMenu.DeleteTask(index - 1)

		case 5:
			var word string
			fmt.Print("Введите нужное слово для поиска: ")
			fmt.Scan(&word)
			taskMenu.SearchTask(word)

		case 6:
			err := SaveData(filename, tasks)
			if err != nil {
				fmt.Println("Всё пропало!!! Не удалось сохранить данные, ОШИБКА: КОД 6", err)
 } else {
				fmt.Println("Ура!!! Данные успешно сохранены.")
			}
			return

		default:
			fmt.Println("Выбор неверный, пожалуйста, подумай ещё.")
		}
	}
}
