package labs

import "fmt"

type Phones struct {
	brand string
	price float64
	name  string
	color string
}

func (p *Phones) getPrice() float64 {
	return p.price
}

func (p *Phones) setPrice(newPrice float64) {
	p.price = newPrice
}

func (p *Phones) applyDiscount(discount float64) {
	if discount < 0 || discount > 100 {
		fmt.Println("Скидка должна быть в диапазоне от 0 до 100%")
		return
	}
	p.price = p.price * (1 - discount/100)
}

func (p *Phones) getName() string {
	return p.name
}

func (p *Phones) setName(newName string) {
	p.name = newName
}

func (p *Phones) getBrand() string {
	return p.brand
}

func (p *Phones) setBrand(newBrand string) {
	p.brand = newBrand
}

func (p *Phones) getColor() string {
	return p.color
}

func (p *Phones) setColor(newColor string) {
	p.color = newColor
}
