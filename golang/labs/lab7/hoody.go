package lab7

import "fmt"

type Hoody struct {
	size  string
	price float64
	name  string
	color string
}

func (h *Hoody) getPrice() float64 {
	return h.price
}

func (h *Hoody) setPrice(newPrice float64) {
	h.price = newPrice
}

func (h *Hoody) applyDiscount(discount float64) {
	if discount < 0 || discount > 100 {
		fmt.Println("Скидка должна быть в диапазоне от 0 до 100%")
		return
	}
	h.price = h.price * (1 - discount/100)
}

func (h *Hoody) getName() string {
	return h.name
}

func (h *Hoody) setName(newName string) {
	h.name = newName
}
