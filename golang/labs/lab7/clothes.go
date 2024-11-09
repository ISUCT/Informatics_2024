package labs

type Clothes struct {
	price       float64
	clothestype string
	material    string
}

func (c *Clothes) getPrice() float64 {
	return c.price
}

func (c *Clothes) discount(discount float64) {
	c.price -= c.price * discount / 100
}

func (c *Clothes) setPrice(newPrice float64) {
	c.price = newPrice
}

func (c *Clothes) changeCharacteristic(newClothestype string, newMaterial string) {
	c.clothestype = newClothestype
	c.material = newMaterial
}
