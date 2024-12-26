package lab7

import "fmt"

type Headphones struct {
	Brand     string
	Colour    string
	MaxVolume int
	Price     float64
}

func (h Headphones) GetPrice() float64 {
	return h.Price
}

func (h *Headphones) EditPrice(NewPrice float64) {
	(*h).Price = NewPrice
}

func (h *Headphones) ApplyDiscount(Percent float64) {
	(*h).Price = (h.Price / 100) * (100 - Percent)
}

func (h *Headphones) EditDescription(NewDescription string) {
	(*h).Colour = NewDescription
}

func (h Headphones) GetInfo() {
	fmt.Println("В продаже есть наушники бренда", h.Brand, "цвета", h.Colour, ", которые имеют громкость до", h.MaxVolume, "дб и стоят:", h.Price)
}
