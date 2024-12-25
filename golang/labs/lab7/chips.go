package lab7

import "fmt"

type Chips struct {
	name   string
	taste  string
	weight float64
	price  float64
}

func NewChips(name string, taste string, weight float64, price float64) *Chips {
	chips := &Chips{name: name, taste: taste, weight: weight, price: price}
	return chips
}

func (c *Chips) GetName() string {
	return c.name
}

func (c *Chips) SetPrice(newPrice float64) {
	c.price = newPrice
}

func (c *Chips) GetPrice() float64 {
	return c.price
}

func (c *Chips) GetWeight() float64 {
	return c.weight
}

func (c *Chips) GetTaste() string {
	return c.taste
}

func (c *Chips) GetInfo() {
	fmt.Printf("name: %s\ntaste: %s\nweight: %.2f\nprice: %.2f\n", c.name, c.taste, c.weight, c.price)
}

func (c *Chips) ApplyDiscount(perDiscount float64) {
	c.price = c.price * (1 - (perDiscount / 100))
}
