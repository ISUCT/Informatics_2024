package labs

import "fmt"

type Cat struct {
	Name  string
	Age   float64
	Color string
}

func (c *Cat) setName(newName string) {
	c.Name = newName
}

func (c Cat) getName() string {
	return c.Name
}

func (c *Cat) setAge(newAge float64) {
	c.Age = newAge
}

func (c Cat) getAge() float64 {
	return c.Age
}

func (c *Cat) setColor(newColor string) {
	c.Color = newColor
}

func (c Cat) getColor() string {
	return c.Color
}

func Runlab6() {
	cat := Cat{Name: "Барсик", Age: 1, Color: "Серый"}

	fmt.Println("Имя кота : ", cat.getName())
	cat.setName("Рыжик")
	fmt.Println("Имя кота : ", cat.getName())

	fmt.Println("Возраст кота :", cat.getAge())
	cat.setAge(2.5)
	fmt.Printf("Возраст кота: %f\n", cat.getAge())

	fmt.Println("Цвет кота :", cat.getColor())
	cat.setColor("Рыжий")
	fmt.Println("Цвет кота :", cat.getColor())
}
