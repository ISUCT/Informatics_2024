package laba6

import "fmt"

type Car struct {
	Name  string
	Speed float64
	Color string
}

func (c *Car) setName(newName string) {
	c.Name = newName
}

func (c Car) getName() string {
	return c.Name
}

func (c *Car) setSpeed(newSpeed float64) {
	c.Speed = newSpeed
}

func (c Car) getSpeed() float64 {
	return c.Speed
}

func (c *Car) setColor(newColor string) {
	c.Color = newColor
}

func (c Car) getColor() string {
	return c.Color
}

func CompleteLaba6() {
	car := Car{Name: "LADA", Speed: 180, Color: "Black"}

	fmt.Println("Марка авто : ", car.getName())
	car.setName("Lada")
	fmt.Println("Марка авто : ", car.getName())

	fmt.Println("Скорость авто :", car.getSpeed())
	car.setSpeed(96)
	fmt.Printf("Скорость авто: %f\n", car.getSpeed())

	fmt.Println("Цвет авто :", car.getColor())
	car.setColor("White")
	fmt.Println("Цвет авто :", car.getColor())
}
