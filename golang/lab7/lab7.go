package lab7

import "fmt"

type Product interface {
	GetName() string
	GetPrice() float64
	SetPrice(price float64)
	ApplyDiscount(discount float64)
}

type Fruit struct {
	name   string
	price  float64
	weight float64
}

func (f *Fruit) GetName() string {
	return f.name
}

func (f *Fruit) GetPrice() float64 {
	return f.price
}

func (f *Fruit) SetPrice(price float64) {
	f.price = price
}

func (f *Fruit) ApplyDiscount(discount float64) {
	f.price -= f.price * discount / 100
}

type Electronic struct {
	name  string
	price float64
	brand string
}

func (e *Electronic) GetName() string {
	return e.name
}

func (e *Electronic) GetPrice() float64 {
	return e.price
}

func (e *Electronic) SetPrice(price float64) {
	e.price = price
}

func (e *Electronic) ApplyDiscount(discount float64) {
	e.price -= e.price * discount / 100
}

type Chancellery struct {
	name     string
	price    float64
	color    string
	quantity int
}

func (c *Chancellery) GetName() string {
	return c.name
}

func (c *Chancellery) GetPrice() float64 {
	return c.price
}

func (c *Chancellery) SetPrice(price float64) {
	c.price = price
}

func (c *Chancellery) ApplyDiscount(discount float64) {
	c.price -= c.price * discount / 100
}

func CalculateTotalPrice(products []Product) float64 {
	var total float64
	for _, product := range products {
		total += product.GetPrice()
	}
	return total
}

func Runlab7() {
	fruit := &Fruit{name: "Banana", price: 130.00, weight: 0.8}
	electronic := &Electronic{name: "Hairdryer", brand: "Dyson", price: 50000.00}
	chancellery := &Chancellery{name: "Pen", price: 60.00, color: "Blue", quantity: 3}
	products := []Product{fruit, electronic, chancellery}
	fmt.Println("Общая цена товаров:", CalculateTotalPrice(products))
	fruit.ApplyDiscount(10)
	electronic.ApplyDiscount(20)
	chancellery.ApplyDiscount(40)
	fmt.Println("Цена товаров с учетом скидок:", CalculateTotalPrice(products))
}
