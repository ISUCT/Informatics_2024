package lab7

import "fmt"

type Furnitures struct {
	name        string
	material    string
	price       float64
	description string
}

func NewFurnitures(name string, material string, price float64, descrption string) *Furnitures {
	furnitures := &Furnitures{name: name, material: material, price: price, description: descrption}
	return furnitures
}

func (f *Furnitures) getName() string {
	return f.name
}

func (f *Furnitures) setPrice(newPrice float64) {
	f.price = newPrice
}

func (f *Furnitures) getPrice() float64 {
	return f.price
}

func (f *Furnitures) changeData(name string, price float64, descrption string) {
	f.name = name
	f.price = price
	f.description = descrption
}

func (f *Furnitures) getData() {
	fmt.Printf("name: %s\nmaterial: %s\nprice: %.2f\ndescrption: %s\n\n", f.name, f.material, f.price, f.description)
}

func (f *Furnitures) applyDiscount(perDiscount float64) {
	f.price = f.price * (1 - (perDiscount / 100))
}
