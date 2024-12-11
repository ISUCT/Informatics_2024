package laba7

import "fmt"

func GetEndPrice(products []Product) float64 {
	var endPrice float64
	for _, product := range products {
		endPrice += product.getPrice()
	}
	return endPrice
}

func RunLab7() {
	product1 := &Book{name: "Закат Европы", author: "Освальд Шпеглер", genre: "научный", price: 999}
	product1.applyDiscount(5)
	product2 := &Juice{name: "Любимый", taste: "Апельсин", volume: 2, price: 199}
	product2.applyDiscount(10)
	product3 := &Movie{title: "Гарри Поттер и Узник Азкабана", director: "Джоан Роулинг", duration: 120, price: 299}
	product3.applyDiscount(15)

	products := []Product{product1, product2, product3}

	for _, product := range products {
		fmt.Println(product.getProductInfo())
	}

	totalPrice := GetEndPrice(products)
	fmt.Printf("Общая стоимость: %.2f", totalPrice)
}
