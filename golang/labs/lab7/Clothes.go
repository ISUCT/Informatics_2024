package lab7

type Clothes struct {
	Name  string
	Price float64
	Size  string
	Color string
}

func (c *Clothes) GetName() string {
	return h.Name
}

func (c *Clothes) GetPrice() float64 {
	return c.Price
}

func (c *Clothes) SetPrice(price float64) {
	c.Price = price
}

func (c *Clothes) ApplyDiscount(discount float64) {
	c.Price -= c.Price * discount / 100
}
