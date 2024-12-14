package laba7

import "fmt"

type Gaming struct {
	Name            string
	Characteristics string
	Price           float64
	SailStart       float64
}

func (g *Gaming) ApplyDiscount(discount float64) {
	g.Price = g.Price * (1 - discount/100)
}

func (g *Gaming) GetPrice() float64 {
	return g.Price
}
func (g *Gaming) SetPrice(newPrice float64) {
	g.Price = newPrice
}
func (g *Gaming) UpdateCharacteristics(newCharacteristics string) {
	g.Characteristics = newCharacteristics
}

func (g *Gaming) getInfo() string {
	return fmt.Sprintf("Name: %s, Characteristics: %s, Price: %.2f,SailStart: %2.f", g.Name, g.Characteristics, g.Price, g.SailStart)
}
