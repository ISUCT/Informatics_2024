package LAB_7

import "fmt"

type Product interface {
	GetName() string
	GetPrice() float64
	SetPrice(price float64)
	Discount(discount float64)
}

type Laptop struct {
	Name  string
	Price float64
	Brand string
}

func (l *Laptop) GetName() string {
	return l.Name
}

func (l *Laptop) GetPrice() float64 {
	return l.Price
}

func (l *Laptop) SetPrice(price float64) {
	l.Price = price
}

func (l *Laptop) Discount(discount float64) {
	l.Price -= l.Price * discount / 100
}

type Car struct {
	Name  string
	Price float64
	Brand string
}

func (c *Car) GetName() string {
	return c.Name
}

func (c *Car) GetPrice() float64 {
	return c.Price
}

func (c *Car) SetPrice(price float64) {
	c.Price = price
}

func (c *Car) Discount(discount float64) {
	c.Price -= c.Price * discount / 100
}

func Calculate(products []Product) float64 {
	total := 0.0
	for _, product := range products {
		total += product.GetPrice()
	}
	return total
}

func Lab7() {
	laptop := &Laptop{"Делл XPS 13", 40000.00, "Dell"}
	car := &Car{"Тойота Королла", 260000.00, "Toyota"}

	products := []Product{laptop, car}
	fmt.Println("Общая стоимость", Calculate(products))

	laptop.Discount(17)
	car.Discount(7)
	fmt.Println("Общая стоимость товаров после применения скидок", Calculate(products))
}
