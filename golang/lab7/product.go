package Lab7

type Item struct {
	Name  string
	Price float64
}

func (i *Item) GetName() string {
	return i.Name
}

func (i *Item) GetPrice() float64 {
	return i.Price
}

func (i *Item) SetPrice(price float64) {
	i.Price = price
}

func (i *Item) ApplyDiscount(discount float64) {
	i.Price -= discount
	if i.Price < 0 {
		i.Price = 0
	}
}
