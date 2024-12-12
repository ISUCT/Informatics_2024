package sevenlab

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

func Start7lab() {
	product1 := &Electronics{Name: "Пейджер", Price: 30, Brand: "Mototola", Model: "Wirelesslinkx"}
	if err := product1.applyDiscount(25); err != nil {
		fmt.Println(err)
		return
	}
	product2 := &Clothing{Name: "Футболка", Price: 12.65, Size: "L", Color: "Black"}
	if err := product2.applyDiscount(10); err != nil {
		fmt.Println(err)
		return
	}
	product3 := &Food{Name: "Ролтон", Price: 3.23, Weight: 1}
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
