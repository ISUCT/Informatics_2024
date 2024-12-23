package Lab_7

type Roller struct {
	price    float32
	color    string
	material string
}

func (r *Roller) GetPrice() float32 {
	return r.price
}

func (r *Roller) Discount(percent float32) {
	r.price = r.price * (1 - percent/100)
}

func (r *Roller) ChangePrice(newPrice float32) {
	r.price = newPrice
}

func (r *Roller) ChangeCharacteristic(newColor, newMaterial string) {
	r.color = newColor
	r.material = newMaterial
}
