package lab7

import "fmt"

type Car struct {
	Name  string
	Model string
	Price float64
}

func (c *Car) GetInfo() {
	fmt.Println("Автомобиль:", c.Name, "модель:", c.Model, "стоимость:", c.Price)
}

func (c *Car) GetPrice() float64 {
	return c.Price
}

func (c *Car) MakeDiscount(discountPercentage float64) {
	if discountPercentage >= 100 {
		c.Price = 0
		return
	}
	c.Price = c.Price * (1 - discountPercentage/100)
}

func (c *Car) ChangePrice(newPrice float64) {
	c.Price = newPrice
}

func (c *Car) ChangeDescription(newModel string) {
	c.Model = newModel
}
