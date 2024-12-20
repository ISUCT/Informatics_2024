package lab7

type Book struct {
	name      string
	price     float64
	varieties string
}

func (b *Book) GetName() string {
	return b.name
}

func (b *Book) SetName(NewName string) {
	b.name = NewName
}

func (b *Book) GetPrice() float64 {
	return b.price
}

func (b *Book) SetPrice(NewPrice float64) {
	b.price = NewPrice
}

func (b *Book) ApplyDiscount(discount float64) {
	b.price = b.price * (100 - discount) / 100
}

func (b *Book) SetVarieties(NewVarieties string) {
	b.varieties = NewVarieties
}
