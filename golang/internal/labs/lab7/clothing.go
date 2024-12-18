package lab7

type Clothing struct {
	name        string
	price       float64
	description string
	size        string
}

func (c *Clothing) GetPrice() float64 {
	return c.price
}

func (c *Clothing) GetName() string {
	return c.name
}

func (c *Clothing) ApplyDiscount(discount float64) {
	c.price -= c.price * discount
}

func (c *Clothing) SetPrice(price float64) {
	c.price = price
}

func (c *Clothing) SetName(name string) {
	c.name = name
}

func (c *Clothing) GetDescription() string {
	return c.description
}

func (c *Clothing) SetDescription(description string) {
	c.description = description
}
