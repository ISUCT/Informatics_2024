package Lab_7

type Apple struct {
	price    float32
	color    string
	material string
}

func (a *Apple) GetPrice() float32 {
	return a.price
}

func (a *Apple) Discount(percent float32) {
	a.price = a.price * (1 - percent/100)
}

func (a *Apple) ChangePrice(newPrice float32) {
	a.price = newPrice
}

func (a *Apple) ChangeCharacteristic(newColor, newMaterial string) {
	a.color = newColor
	a.material = newMaterial
}
