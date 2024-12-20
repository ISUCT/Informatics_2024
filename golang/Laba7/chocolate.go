package lab7

type Chocolate struct {
	name  string
	price float64
	brand string
}

func (ch *Chocolate) GetName() string {
	return ch.name
}

func (ch *Chocolate) SetName(NewName string) {
	ch.name = NewName
}

func (ch *Chocolate) GetPrice() float64 {
	return ch.price
}

func (ch *Chocolate) SetPrice(NewPrice float64) {
	ch.price = NewPrice
}

func (ch *Chocolate) ApplyDiscount(discount float64) {
	ch.price = ch.price * (100 - discount) / 100
}

func (ch *Chocolate) SetBrand(NewBrand string) {
	ch.brand = NewBrand
}
