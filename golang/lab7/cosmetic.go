package lab7

type Cosmetic struct {
	name  string
	price float64
	brand string
}

func (c *Cosmetic) GetName() string {
	return c.name
}

func (c *Cosmetic) SetName(NewName string) {
	c.name = NewName
}

func (c *Cosmetic) GetPrice() float64 {
	return c.price
}

func (c *Cosmetic) SetPrice(NewPrice float64) {
	c.price = NewPrice
}

func (c *Cosmetic) ApplyDiscount(discount float64) {
	c.price = c.price * (100 - discount) / 100
}

func (c *Cosmetic) SetBrand(NewBrand string) {
	c.brand = NewBrand
}
