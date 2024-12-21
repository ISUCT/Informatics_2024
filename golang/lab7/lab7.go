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
	mobile := &Mobile{"Айфон 7 plus", 20000.00, "Iphone", "gold"}
	headphones := &Headphones{"Маршал major 4", 7000.00, "Marshall", "black"}
	sneakers := &Sneakers{"Найк GT Cut Academy", 20000.00, "Nike", "white"}
	products := []Product{mobile, headphones, sneakers}
	fmt.Println("Стоимость без скидок:", GetTotalPrice(products))

	mobile.ApplyDiscount(20)
	headphones.ApplyDiscount(30)
	sneakers.ApplyDiscount(50)
	fmt.Println("Стоимость после учёта скидки:", GetTotalPrice(products))
}
