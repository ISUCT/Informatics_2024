package labs

import "fmt"

type Product interface {
	GetName() string
	GetPrice() float64
	SetPrice(price float64)
	Discount(discount float64)
}

type Icecream struct {
	Name  string
	Price float64
	Taste string
}

func (i *Icecream) GetName() string {
	return i.Name
}

func (i *Icecream) GetPrice() float64 {
	return i.Price
}

func (i *Icecream) SetPrice(price float64) {
	i.Price = price
}

func (i *Icecream) Discount(discount float64) {
	i.Price -= i.Price * discount / 100
}

type Milk struct {
	Name    string
	Price   float64
	Percent float64
}

func (m *Milk) GetName() string {
	return m.Name
}

func (m *Milk) GetPrice() float64 {
	return m.Price
}

func (m *Milk) SetPrice(price float64) {
	m.Price = price
}

func (m *Milk) Discount(discount float64) {
	m.Price -= m.Price * discount / 100
}

func Calculate(products []Product) float64 {
	total := 0.0
	for _, product := range products {
		total += product.GetPrice()
	}
	return total
}

func RunLab7() {
	icecream := &Icecream{"Магнат", 119.99, "шоколад"}
	milk := &Milk{"Весёлая коровка", 90.00, 3.2}

	products := []Product{icecream, milk}
	fmt.Println("Общая стоимость товаров:", Calculate(products))

	icecream.Discount(5)
	milk.Discount(15)
	fmt.Println("Общая стоимость товаров после применения скидок:", Calculate(products))
}
