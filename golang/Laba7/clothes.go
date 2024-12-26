package lab7

type Clothes struct {
	name     string
	price    float64
	material string
}

func (c *Clothes) GetName() string {
	return c.name
}

func (c *Clothes) SetName(NewName string) {
	c.name = NewName
}

func (c *Clothes) GetPrice() float64 {
	return c.price
}

func (c *Clothes) SetPrice(NewPrice float64) {
	c.price = NewPrice
}

func (c *Clothes) ApplyDiscount(discount float64) {
	c.price = c.price * (100 - discount) / 100
}

func (c *Clothes) SetMaterial(NewMaterial string) {
	c.material = NewMaterial
}
