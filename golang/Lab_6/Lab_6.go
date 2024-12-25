package Lab6

import (
	"fmt"
	"strconv"
)

type Character struct {
	Name  string
	Class string
	Level int
}

func (c *Character) SetName(name string) {
	c.Name = name
}
func (c *Character) SetClass(class string) {
	c.Class = class
}
func (c *Character) SetLevel(level int) {
	if level > 0 {
		c.Level = level
	} else {
		fmt.Println("ошибка: уровень > 0")
	}
}
func (character Character) GetInfo() string {
	return "имя: " + character.Name + "\nкласс: " + character.Class + "\nуровень: " + strconv.Itoa(character.Level)
}
func NewCharacter(name string, class string, level int) Character {
	return Character{
		Name:  name,
		Class: class,
		Level: level,
	}
}
func RunLab6() {
	character := NewCharacter("Рикимару Сафронович", "воин", 10)
	fmt.Println(character.GetInfo())
	character.SetName("Инвокер Павлович")
	character.SetClass("Маг")
	character.SetLevel(15)
	fmt.Println(character.GetInfo())
	character.SetLevel(-5)
}
