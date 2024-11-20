package lab7

type tire struct {
	name string
	price uint
	discount uint
	catigories string
}

func (c *tire) setName(newName string) {
	c.name = newName
}

func (c *tire) setPrice(newPrice uint) {
	c.price = newPrice
}

func (c *tire) setDiscount(newDiscount uint) {
	c.discount = newDiscount
}

func (c *tire) setCatigories(newCatigories string) {
	c.catigories = newCatigories
}

func (c tire) getInfo() (string, uint, uint, string) {
	return c.name, c.price, c.discount, c.catigories
}