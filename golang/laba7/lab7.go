package laba7

import "fmt"

func GetEndPrice(products []Product) float64 {
	var endPrice float64
	for _, product := range products {
		endPrice += product.GetPrice()
	}
	return endPrice
}

func RunLab7() {
	product1 := &Book{Name: "Закат Европы", Author: "Освальд Шпеглер", Genre: "научный", Price: 999}
	product1.ApplyDiscount(5)
	product2 := &Juice{Name: "Любимый", Taste: "Апельсин", Volume: 2, Price: 199}
	product2.ApplyDiscount(10)
	product3 := &Movie{Title: "Гарри Поттер и Узник Азкабана", Director: "Джоан Роулинг", Duration: 120, Price: 299}
	product3.ApplyDiscount(15)

	products := []Product{product1, product2, product3}

	for _, product := range products {
		fmt.Println(product.GetProductInfo())
	}

	totalPrice := GetEndPrice(products)
	fmt.Printf("Общая стоимость: %.2f", totalPrice)
}
