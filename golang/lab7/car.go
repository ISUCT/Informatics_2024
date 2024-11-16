package lab7

type Car struct {
	price    float64
	colour   string
	brandcar string
}

func (c *Car) getPrice() float64 {
	return c.price
}

func (c *Car) setPrice(newPrice float64) {
	c.price = newPrice
}

func (c *Car) discount(discount float64) {
	c.price = c.price * (100 - discount) / 100
}

func (c *Car) setCharacteristics(newColour string, newBrandcar string) {
	c.colour = newColour
	c.brandcar = newBrandcar
}
