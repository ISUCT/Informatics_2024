package lab7

type car struct {
	name string
	price uint
	discount uint
	catigories string
}

func (c *car) setName(newName string) {
	c.name = newName
}

func (c *car) setPrice(newPrice uint) {
	c.price = newPrice
}

func (c *car) setDiscount(newDiscount uint) {
	c.discount = newDiscount
}

func (c *car) setCatigories(newCatigories string) {
	c.catigories = newCatigories
}

func (c car) getInfo() (string, uint, uint, string) {
	return c.name, c.price, c.discount, c.catigories
}