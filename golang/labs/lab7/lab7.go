package lab7

import "fmt"

type Product interface {
	GetName() string
	GetPrice() float64
	SetPrice(price float64)
	ApplyDiscount(discount float64)
}

func GetTotalPrice(products []Product) float64 {
	var TotalPrice float64 = 0
	for _, product := range products {
		TotalPrice += product.GetPrice()
	}
	return TotalPrice
}
func Lab7() {
    Clothes := &Clothes{Name: "Шорты", Price: 1000, Size: "S", Color: "White"}
    Vegetables := &Vegetables{Name: "Перец", Weight: 100, Price: 100}
    Book := &Book{Name: "Алиса в стране чудес", Format: "бумажный", Price: 450}
    products := []Product{Clothes, Vegetables, Book}
    Book.SetFormat("электронный")
	Clothes.ApplyDiscount(100)
	Vegetables.ApplyDiscount(50)
	Book.ApplyDiscount(30)
    totalPrice := GetTotalPrice(products)
    fmt.Println("Новая цена", GetTotalPrice(products))
}

