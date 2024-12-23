package Lab_7

type Orange struct {
	price    float32
	color    string
	material string
}

func (o *Orange) GetPrice() float32 {
	return o.price
}

func (o *Orange) Discount(percent float32) {
	o.price = o.price * (1 - percent/100)
}

func (o *Orange) ChangePrice(newPrice float32) {
	o.price = newPrice
}

func (o *Orange) ChangeCharacteristic(newColor, newMaterial string) {
	o.color = newColor
	o.material = newMaterial
}
