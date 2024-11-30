package lab7

import "fmt"

type Furnitures struct {
	name        string
	material    string
	price       float64
	color       string
	description string
}

func NewFurnitures(name string, material string, price float64, descrption string, color string) *Furnitures {
	furnitures := &Furnitures{name: name, material: material, price: price, description: descrption, color: color}
	return furnitures
}

func (f *Furnitures) getName() string {
	return f.name
}

func (f *Furnitures) getColor() string {
	return f.color
}

func (f *Furnitures) setPrice(newPrice float64) {
	f.price = newPrice
}

func (f *Furnitures) getPrice() float64 {
	return f.price
}

func (f *Furnitures) getInfo() {
	fmt.Printf("name: %s\nmaterial: %s\ncolor: %s\nprice: %.2f\ndescrption: %s\n\n", f.name, f.material, f.color, f.price, f.description)
}

func (f *Furnitures) applyDiscount(perDiscount float64) {
	f.price = f.price * (1 - (perDiscount / 100))
}
