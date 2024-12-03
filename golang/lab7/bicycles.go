package lab7

type Bicycles struct {
	price float32
	color string
	material string
}

func (x Bicycles) GetPrice() float32 {
	return x.price
}

func (c *Bicycles) Discount(x float32) {
	(*c).price = (c.price / 100) * (100 - x)
}

func (c *Bicycles) ChangePrice(x float32) {
	(*c).price = x
}

func (c *Bicycles) ChangeCharacteristic(newcolor string, newmaterial string) {
	c.color = newcolor
	c.material = newmaterial
}