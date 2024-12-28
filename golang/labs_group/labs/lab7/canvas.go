package lab7

type Canvas struct {
	name  string
	price float64
	size  string
}

func (c *Canvas) GetName() string {
	return c.name
}

func (c *Canvas) GetPrice() float64 {
	return c.price
}

func (c *Canvas) SetPrice(price float64) {
	c.price = price
}

func (c *Canvas) ApplyDiscount(discount float64) {
	c.price -= c.price * discount / 100
}
