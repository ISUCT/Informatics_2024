package Lab_7

type Bike struct {
	price    float32
	color    string
	material string
}

func (b *Bike) GetPrice() float32 {
	return b.price
}

func (b *Bike) Discount(percent float32) {
	b.price = b.price * (1 - percent/100)
}

func (b *Bike) ChangePrice(newPrice float32) {
	b.price = newPrice
}

func (b *Bike) ChangeCharacteristic(newColor, newMaterial string) {
	b.color = newColor
	b.material = newMaterial
}
