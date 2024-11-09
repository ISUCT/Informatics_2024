package labs

type Cosmetics struct {
	price         float64
	brand         string
	cosmeticstype string
}

func (c *Cosmetics) getPrice() float64 {
	return c.price
}

func (c *Cosmetics) discount(discount float64) {
	c.price -= c.price * discount / 100
}

func (c *Cosmetics) setPrice(newPrice float64) {
	c.price = newPrice
}

func (c *Cosmetics) changeCharacteristic(newBrand string, newCosmeticstype string) {
	c.brand = newBrand
	c.cosmeticstype = newCosmeticstype
}
