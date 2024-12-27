package LABA6

import "fmt"

type Fox struct {
	Name  string
	Age   int
	Color string
}

func NewFox(name string, age int, color string) *Fox {
	return &Fox{
		Name:  name,
		Age:   age,
		Color: color,
	}
}

func (f *Fox) GetView() string {
	return "   /\\   /\\\n" +
		"  (  o o  )\n" +
		"   \\  ^  / \n" +
		"    |   |  \n" +
		"   /     \\ \n" +
		"  /       \\ \n" +
		" /  |   |  \\ \n" +
		"(   |   |   )"
}
func (f *Fox) Describe() string {
	return fmt.Sprintf("Имя лисы: %s, возраст: %d года, цвет: %s.", f.Name, f.Age, f.Color)
}

func Lab6() {
	fox := NewFox("Феликс", 3, "рыжий")
	fmt.Println(fox.Describe())
	fmt.Println(fox.GetView())
}
