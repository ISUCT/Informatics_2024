package lab7

type Car struct {
	price float64
	name  string
	color string
}

func (c *Car) getPrice() float64 {
	return c.price
}

func (c *Car) setPrice(newPrice float64) {
	c.price = newPrice
}

func (c *Car) applyDiscount(discount float64) {
	c.price = c.price * (100 - discount) / 100
}

func (c *Car) getName() string {
	return c.name
}

func (c *Car) setName(newName string) {
	c.name = newName
}

func (c *Car) setColor(newColor string) {
	c.color = newColor
}
