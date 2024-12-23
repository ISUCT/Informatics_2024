package lab7

type Book struct {
	name       string
	price      float64
	numAuthors int
}

func (b Book) GetName() string {
	return b.name
}

func (b Book) GetPrice() float64 {
	return b.price
}

func (b *Book) SetPrice(price float64) {
	b.price = price
}

func (b *Book) ApplyDiscount(discount float64) {
	b.price -= b.price * discount
}
