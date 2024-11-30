package lab7

import "fmt"

type Clothes struct {
	name        string
	brand       string
	price       float64
	materials   string
	description string
}

func NewCloth(name string, brand string, price float64, descrption string, materials string) *Clothes {
	clothes := &Clothes{name: name, brand: brand, price: price, description: descrption, materials: materials}
	return clothes
}

func (c *Clothes) getName() string {
	return c.name
}

func (c *Clothes) getMaterials() string {
	return c.materials
}

func (c *Clothes) setPrice(newPrice float64) {
	c.price = newPrice
}

func (c *Clothes) getPrice() float64 {
	return c.price
}

func (c *Clothes) getInfo() {
	fmt.Printf("name: %s\nbrand: %s\nprice: %.2f\ndescrption: %s\n\n", c.name, c.brand, c.price, c.description)
}

func (c *Clothes) applyDiscount(perDiscount float64) {
	c.price = c.price * (1 - (perDiscount / 100))
}
