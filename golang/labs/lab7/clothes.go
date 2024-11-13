package lab7

import "fmt"

type Clothes struct {
	name        string
	brand       string
	price       float64
	description string
}

func NewCloth(name string, brand string, price float64, descrption string) *Clothes {
	clothes := &Clothes{name: name, brand: brand, price: price, description: descrption}
	return clothes
}

func (c *Clothes) getName() string {
	return c.name
}

func (c *Clothes) setPrice(newPrice float64) {
	c.price = newPrice
}

func (c *Clothes) getPrice() float64 {
	return c.price
}

func (c *Clothes) changeData(name string, price float64, descrption string) {
	c.name = name
	c.price = price
	c.description = descrption
}

func (c *Clothes) getData() {
	fmt.Printf("name: %s\nbrand: %s\nprice: %.2f\ndescrption: %s\n\n", c.name, c.brand, c.price, c.description)
}

func (c *Clothes) applyDiscount(perDiscount float64) {
	c.price = c.price * (1 - (perDiscount / 100))
}
