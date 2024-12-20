package lab7

type Book struct {
	Name        string
	Price       float64
	OldPrice    float64
	Description string
	Author      string
}

func (b *Book) GetName() string {
	return b.Name
}

func (b *Book) GetPrice() float64 {
	return b.Price
}

func (b *Book) GetDescription() string {
	return b.Description
}

func (b *Book) ApplyDiscount(discount float64) {
	b.OldPrice = b.Price
	b.Price = b.Price - (b.Price * discount)
}

func (b *Book) SetPrice(newPrice float64) {
	b.OldPrice = b.Price
	b.Price = newPrice
}

func (b *Book) SetDescription(newDescription string) {
	b.Description = newDescription
}

func (b *Book) GetOldPrice() float64 {
	return b.OldPrice
}
