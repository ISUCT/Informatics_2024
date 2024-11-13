package lab7

import (
	"fmt"
)

type Products interface {
	getName() string
	setPrice(float64)
	getPrice() float64
	changeData(string, float64, string)
	getData()
	applyDiscount(float64)
}

func calculateDiscount(products []Products) float64 {
	var sum float64 = 0
	for _, product := range products {
		sum += product.getPrice()
	}

	return sum
}

func RunLab7Task() {
	product1 := NewBook("The Go Programming Language", "Alan A. A. Donovan", 45.99, "A book on Go programming.")
	product2 := NewCloth("T-shirt", "Nike", 19.99, "A comfortable cotton T-shirt.")
	product3 := NewFurnitures("Sofa", "Leather", 799.99, "A stylish leather sofa.")

	listOfProduct := []Products{product1, product2, product3}

	fmt.Printf("\nПродукты до изменения\n\n")
	product1.getData()
	product2.getData()
	product1.changeData("Advanced Go Programming", 50.00, "An advanced guide to Go programming.")
	product2.changeData("Premium T-shirt", 25.00, "A high-quality cotton T-shirt.")
	fmt.Printf("\nПродукты после изменения\n\n")
	product1.getData()
	product2.getData()

	fmt.Println("Товар-----------Цена")
	for _, product := range listOfProduct {
		fmt.Printf("%s-------%.2f $\n", product.getName(), product.getPrice())
	}
	fmt.Printf("Цена корзины до скидки: %.2f $\n\n", calculateDiscount(listOfProduct))

	product1.applyDiscount(10)
	product2.applyDiscount(20)
	product3.applyDiscount(15)

	fmt.Println("Товар-----------Цена")
	for _, product := range listOfProduct {
		fmt.Printf("%s-------%.2f $\n", product.getName(), product.getPrice())
	}
	fmt.Printf("Цена корзины после скидки: %.2f $\n", calculateDiscount(listOfProduct))
}
