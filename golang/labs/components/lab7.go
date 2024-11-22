package components

import (
	"fmt"
)

func GeneralPrice(products []Product) float64 {
	var totalCost float64
	for _, product := range products {
		totalCost += product.price()
	}
	return totalCost
}

func RunLab7() {
	product1 := &Electronics{Name: "Телефон", Price: 20, Brand: "Iphone", Model: "5s"}
	if err := product1.sale(25); err != nil {
		fmt.Println(err)
		return
	}
	product2 := &Clothing{Name: "Рубашка", Price: 14.32, Size: "L", Color: "white"}
	if err := product2.sale(10); err != nil {
		fmt.Println(err)
		return
	}
	product3 := &Food{Name: "Яблоко", Price: 2.35, Weight: 1}
	if err := product3.sale(15); err != nil {
		fmt.Println(err)
		return
	}

	products := []Product{product1, product2, product3}

	for _, product := range products {
		fmt.Println(product.productFeature())
	}

	total := GeneralPrice(products)
	fmt.Printf("Общая стоимость: %.2f\n", total)
}
