package lab6

import "fmt"

type Table struct {
	title string
	size  int
	color string
}

func (t *Table) change_size(new_size int) {
	(*t).size = new_size
}

func (t *Table) change_color(new_color string) {
	(*t).color = new_color
}

func (t *Table) new_table(new_title, new_color string, new_size int) {
	(*t).title = new_title
	(*t).size = new_size
	(*t).color = new_color
}

func (t Table) Title() string {
	return t.title
}

func (t Table) Size() int {
	return t.size
}

func (t Table) Color() string {
	return t.color
}

func Show_lab6() {
	My_table := Table{
		title: "Хай - тек",
		size:  5,
		color: "красный",
	}

	fmt.Println("Характеристики стола: название - ", My_table.Title(), "| размер - ", My_table.Size(), "| цвет - ", My_table.Color())

	My_table.change_size(6)
	My_table.change_color("бирюзовый")

	fmt.Println("================================================================================================")
	fmt.Println("Изменены размер и цвет с помощью методов change_size и change_color:")
	fmt.Println("Характеристики стола: название - ", My_table.Title(), "| размер - ", My_table.Size(), "| цвет - ", My_table.Color())

	My_table.new_table("Стол-раскладушка", "Черный", 9)

	fmt.Println("================================================================================================")
	fmt.Println("Полностью изменены характеристики стола с помощью метода new_table:")
	fmt.Println("Характеристики стола: название - ", My_table.Title(), "| размер - ", My_table.Size(), "| цвет - ", My_table.Color())
}