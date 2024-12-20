package lab7

type Clothes struct {
	price float64
	name  string
	size  string
}

func (c *Clothes) getName() string {
	return c.name
}

func (c *Clothes) setName(newName string) {
	c.name = newName
}

func (c *Clothes) getPrice() float64 {
	return c.price
}

func (c *Clothes) setPrice(newPrice float64) {
	c.price = newPrice
}

func (c *Clothes) getSize() string {
	return c.size
}

func (c *Clothes) setSize(newSize string) {
	c.size = newSize
}

func (c *Clothes) applyDiscount(discount float64) {
	c.price = c.price * (100 - discount) / 100
}
