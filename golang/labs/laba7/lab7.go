package laba7

import (
	"fmt"
)

func GetFinalPrice(products []Product) float64 {
	var totalCost float64
	for _, product := range products {
		totalCost += product.getPrice()
	}
	return totalCost
}

func RunLab7() {
	product1 := &Electronics{Name: "Телефон", Price: 20, Brand: "Iphone", Model: "5s"}
	if err := product1.applyDiscount(25); err != nil {
		fmt.Println(err)
		return
	}
	product2 := &Clothing{Name: "Рубашка", Price: 14.32, Size: "L", Color: "white"}
	if err := product2.applyDiscount(10); err != nil {
		fmt.Println(err)
		return
	}
	product3 := &Food{Name: "Яблоко", Price: 2.35, Weight: 1}
	if err := product3.applyDiscount(15); err != nil {
		fmt.Println(err)
		return
	}

	products := []Product{product1, product2, product3}

	for _, product := range products {
		fmt.Println(product.getProductInfo())
	}

	total := GetFinalPrice(products)
	fmt.Printf("Общая стоимость: %.2f\n", total)
}
