package lab7

import (
	"fmt"
)

type Product interface {
	getName() string
	setName(string)
	getPrice() float64
	setPrice(float64)
	applyDiscount(float64)
}

func CalculateProductSum(productlist []Product) string {
	var sum float64 = 0
	for _, product := range productlist {
		sum += product.getPrice()
	}
	return fmt.Sprintf("%.2f", sum)
}

func RunLab7() {
	pants := &Clothes{name: "штаны", price: 1499, size: "M"}
	chocolate := &Food{name: "шоколад", price: 99.99, calories: 576}
	phone := &Technic{name: "Iphone 15 Pro Max", price: 170990, weight: 227}

	pants.setSize("L")
	chocolate.setPrice(109.99)
	phone.setName("Iphone 16 Pro Max")
	phone.setWeight(225)

	productlist := []Product{pants, chocolate, phone}
	fmt.Printf("Сумма товаров без скидки: %v рублей.\n", CalculateProductSum(productlist))
	pants.applyDiscount(10)
	chocolate.applyDiscount(30)
	phone.applyDiscount(5)
	fmt.Printf("Сумма товаров со скидкой: %v рублей.\n", CalculateProductSum(productlist))
}
