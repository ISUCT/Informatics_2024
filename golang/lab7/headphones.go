package lab7

type Headphones struct {
	Name  string
	Price float64
	Brand string
	Color string
}

func (h *Headphones) GetName() string {
	return h.Name
}

func (h *Headphones) GetPrice() float64 {
	return h.Price
}

func (h *Headphones) SetPrice(price float64) {
	h.Price = price
}

func (h *Headphones) ApplyDiscount(discount float64) {
	h.Price -= h.Price * discount / 100
}
