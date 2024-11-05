package lab7

import "fmt"

type Product interface {
	Discount(float64)
}

type product struct {
	ID    int
	name  string
	price float64
}

func newProduct(ID int, name string, price float64) *product {
	p := &product{
		ID:    ID,
		name:  name,
		price: price,
	}
	return p
}

type car struct {
	*product
	color string
}

func (c *car) Discount(percent float64) {
	c.price -= (c.price / 100) * percent
}

func NewCar(ID int, name string, price float64, color string) *car {
	c := &car{
		product: newProduct(ID, name, price),
		color:   color,
	}
	return c
}

// ___

type furniture struct {
	*product
	material string
}

func (f *furniture) Discount(percent float64) {
	f.price -= (f.price / 100) * percent
}

func NewFurniture(ID int, name string, price float64, material string) *furniture {
	f := &furniture{
		product:  newProduct(ID, name, price),
		material: material,
	}
	return f
}

func Lab7() {
	bentli := NewCar(0, "бентли", 100000, "красный")
	sofa := NewFurniture(0, "диван", 5000, "дерево")

	var ProductList []Product = []Product{bentli, sofa}

	for _, product := range ProductList {
		product.Discount(10)
	}

	fmt.Println(bentli.price)
	fmt.Println(sofa.price)
}
