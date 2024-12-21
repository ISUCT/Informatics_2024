package Lab7

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
	product1 := &Device{Name: "Ноутбук", Price: 30, Brand: "Honor", Model: "1"}
	product2 := &Cloth{Name: "Свитер", Price: 3.52, Size: "L", Color: "Красный"}
	product3 := &Food{Name: "Арбуз", Price: 1.45, Weight: 1}
	products := []Product{product1, product2, product3}
	for _, product := range products {
		fmt.Println(product.getProductInfo())
	}
	total := GetFinalPrice(products)
	fmt.Printf("Общая стоимость до скидки: %.2f\n", total)

	if err := product1.applyDiscount(25); err != nil {
		fmt.Println(err)
		return
	}
	if err := product2.applyDiscount(10); err != nil {
		fmt.Println(err)
		return
	}
	if err := product3.applyDiscount(15); err != nil {
		fmt.Println(err)
		return
	}
	total = GetFinalPrice(products)
	fmt.Printf("Общая стоимость после скидки: %.2f\n", total)
}
