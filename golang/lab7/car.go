package lab7

import "fmt"

type Car struct {
	Name string
	Model string
	Price float64
}

func (c *Car) GetInfo() {
	fmt.Println("Автомобиль:", c.Name, "модель:", c.Model, "стоимость:", c.Price)
}

func (c *Car) GetPrice() float64 {
	return c.Price
}

func (c *Car) MakeDiscount(x float64) {
	if x >= 100 {
		c.Price = 0
		return
	}
	c.Price = c.Price * (1 - x/100)
}

func (c *Car) ChangePrice(x float64) {
	c.Price = x
}

func (c *Car) ChangeDescription(x string) {
	c.Model = x
}