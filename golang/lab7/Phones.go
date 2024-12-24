package lab7

import "fmt"

type Phones struct {
	Name            string
	Characteristics string
	Price           float64
	ReleaseDate     string
}

func (p *Phones) ApplyDiscount(discount float64) {
	p.Price = p.Price * (1 - discount/100)
}
func (p *Phones) GetPrice() float64 {
	return p.Price
}
func (p *Phones) SetPrice(newPrice float64) {
	p.Price = newPrice
}
func (p *Phones) UpdateCharacteristics(newCharacteristics string) {
	p.Characteristics = newCharacteristics
}
func (p *Phones) getInfo() string {
	return fmt.Sprintf("Name: %s, Characteristics: %s, Price: %.2f,ReleaseDate: %s", p.Name, p.Characteristics, p.Price, p.ReleaseDate)
}
