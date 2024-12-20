package abc

import "fmt"

type Car struct {
	Brand string
	Model string
	Speed int
}

func NewCar(brand, model string, speed int) Car {
	return Car{
		Brand: brand,
		Model: model,
		Speed: speed,
	}
}

func (c Car) GetSpeed() int {
	return c.Speed
}

func (c *Car) SetSpeed(newSpeed int) {
	c.Speed = newSpeed
}

func (c Car) DisplayInfo() {
	fmt.Printf("Автомобиль: %s %s, Скорость: %d км/ч\n", c.Brand, c.Model, c.Speed)
}

func abc() {
	car := NewCar("Toyota", "Corolla", 100)
	car.DisplayInfo()

	car.SetSpeed(120)
	fmt.Printf("Новая скорость: %d км/ч\n", car.GetSpeed())
}
