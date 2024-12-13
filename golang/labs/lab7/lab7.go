package lab7

import (
	"fmt"
)

type Products interface {
	getName() string
	setPrice(float64)
	getPrice() float64
	getInfo()
	applyDiscount(float64)
}

func calculateTotalPrice(products []Products) float64 {
	var sum float64 = 0
	for _, product := range products {
		sum += product.getPrice()
	}

	return sum
}

func RunLab7Task() {
	product1 := NewBook("The Go Programming Language", "Alan A. A. Donovan", 45.99, "A book on Go programming.", 259)
	product2 := NewCloth("T-shirt", "Nike", 19.99, "A comfortable cotton T-shirt.", "Leather")
	product3 := NewFurnitures("Sofa", "Leather", 799.99, "A stylish leather sofa.", "white")

	listOfProduct := []Products{product1, product2, product3}

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

	fmt.Println("Информация про книгу")
	fmt.Println("Автор:", product1.getAuthor())
	fmt.Println("Кол-во страниц:", product1.getNumberOfPages())
	fmt.Println("Название:", product1.getName())

	fmt.Println("Информация про одежду")
	fmt.Println("Материал:", product2.getMaterials())
	fmt.Println("Название:", product2.getName())

	fmt.Println("Информация про мебель")
	fmt.Println("Цвет:", product3.getColor())
	fmt.Println("Название:", product3.getName())
}
