package lab7

type Shirt struct {
	price float64
	name  string
	size  string
}

func (c *Shirt) GetName() string {
	return c.name
}

func (c *Shirt) SetName(newName string) {
	c.name = newName
}

func (c *Shirt) GetPrice() float64 {
	return c.price
}

func (c *Shirt) SetPrice(newPrice float64) {
	c.price = newPrice
}

func (c *Shirt) SetSize(newSize string) {
	c.size = newSize
}

func (c *Shirt) ApplyDiscount(discount float64) {
	c.price = c.price * (100 - discount) / 100
}
