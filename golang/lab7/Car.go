package lab7

type Car struct {
	name  string
	price float64
	model string
	color string
	
}

func (c *Car) GetName() string {
	return c.name
}

func (c *Car) GetPrice() float64 {
	return c.price
}

func (c *Car) SetPrice(price float64) {
	c.price = price
}

func (c *Car) ApplyDiscount(discount float64) {
	c.price -= c.price * discount / 100
}