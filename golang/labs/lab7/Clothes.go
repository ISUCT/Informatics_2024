package lab7

type Clothes struct {
	Name  string
	Price float64
	Size  string
	Color string
}

func (h *Clothes) GetName() string {
	return h.Name
}

func (h *Clothes) GetPrice() float64 {
	return h.Price
}

func (h *Clothes) SetPrice(price float64) {
	h.Price = price
}

func (h *Clothes) ApplyDiscount(discount float64) {
	h.Price -= h.Price * discount / 100
}
