package lab7

type mascara struct {
	price float32
	color string
	material string
}

func (x mascara) GetPrice() float32 {
	return x.price
}

func (c *mascara]) Discount(x float32) {
 (*c).price = (c.price / 100) * (100 - x)
}

func (c *mascara]) ChangePrice(x float32) {
 (*c).price = x
}

func (c *mascara) ChangeCharacteristic(newcolor string,newmaterial string ) {
	c.color = newcolor
	c.material = newmaterial
}
