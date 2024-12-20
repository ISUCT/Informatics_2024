package lab7

import "fmt"

type Headphones struct {
	Name        string
	Color       string
	Price       float64
	ReleaseDate string
}

func (h *Headphones) ApplyDiscount(discount float64) {
	h.Price = h.Price * (1 - discount/100)
}
func (h *Headphones) GetPrice() float64 {
	return h.Price
}
func (h *Headphones) SetPrice(newPrice float64) {
	h.Price = newPrice
}
func (h *Headphones) ChangeColor(newColor string) {
	h.Color = newColor
}
func (h *Headphones) getInfo() string {
	return fmt.Sprintf("Name: %s, Color: %s, Price: %.2f,ReleaseDate: %s", h.Name, h.Color, h.Price, h.ReleaseDate)
}
