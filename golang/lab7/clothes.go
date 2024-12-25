package lab7

import "fmt"

type Clothes struct {
	name string
	brand string
	price float64
	description string
}

func NewClothes(name string, brand string, price float64, description string) *Clothes {
	clothes := &Clothes{name: name, brand: brand, price: price, description: description}
	return clothes
}

func (c *Clothes) GetName() string {
	return c.name
}

func (c *Clothes) SetPrice(newPrice float64) {
	c.price = newPrice
}

func (c *Clothes) GetPrice() float64 {
	return c.price
}

func (c *Clothes) ChangeData(name string, price float64, description string) {
	c.name = name
	c.price = price
	c.description = description
}

func (c *Clothes) GetData() {
	fmt.Printf("name: %s\nbrand: %s\nprice: %.2f\ndescription: %s\n\n", c.name, c.brand, c.price, c.description)
}

func (c *Clothes) ApplyDiscount(perDiscount float64) {
	c.price = c.price * (1 - (perDiscount / 100))
}