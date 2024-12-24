package lab7

import "fmt"

type Headphones struct {
	Brand      string
	Connection string
	Price      float64
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
	(*h).Connection = NewDescription
}

func (h Headphones) GetInfo() {
	fmt.Println("В продаже есть наушники от производителя", h.Brand, "с", h.Connection, "подключением и стоимостью:", h.Price)
}
