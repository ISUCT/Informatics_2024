package lab7

import (
	"fmt"
)

func CalculationSumProduct(listProducts []Product) float64 {
	var sum float64 = 0
	for _, product := range listProducts {
		sum += product.GetPrice()
	}
	return sum
}

func Lab7Run() {
	lipstick := &Cosmetic{name: "помада", price: 400, brand: "influence"}
	yacht := &Ship{name: "яхта", material: "алюминий", price: 9000000}
	doll := &Toy{name: "барби", price: 1500, material: "пластик"}
	
	doll.SetPrice(3000)
	lipstick.SetBrand("NYX")
	yacht.SetMaterial("сталь")

	listProducts := []Product{lipstick, yacht, doll}
	
	fmt.Printf("Сумма товаров, без учета скидки, равна: %v рублей \n", CalculationSumProduct(listProducts))

	discounts := []float64{10, 35, 10}
	for i, product := range listProducts {
		product.ApplyDiscount(discounts[i])
	}
	fmt.Printf("Сумма товаров, без учета скидки, равна: %v рублей \n", CalculationSumProduct(listProducts))
}
