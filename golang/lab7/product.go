package lab7

import "fmt"

type Product interface {
	GetName() string
	GetPrice() float64
	SetPrice(price float64)
	ApplyDiscount(discount float64)
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
