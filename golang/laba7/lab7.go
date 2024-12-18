package laba7

import (
	"fmt"
)

func CalculationSumProduct(listproducts []Product) string {
	var sum float64 = 0
	for _, product := range listproducts {
		sum += product.GetPrice()
	}
	s := fmt.Sprintf("%.2f", sum)
	return s
}

func RunLab7() {
	var bubbleTea Product = &Food{name: "Баблти", price: 149.99, weight: 10}
	var audi Product = &Car{price: 150000, name: "Ауди", color: "Абсолютный черный", model: "A4", brand: "Audi"}
	var phone Product = &Technic{price: 200000, name: "Audi", color: "Черный"}

	fmt.Println(bubbleTea.ProductInfo())
	fmt.Println(audi.ProductInfo())
	fmt.Println(phone.ProductInfo())

	listproducts := []Product{bubbleTea, audi, phone}
	fmt.Printf("Сумма товаров, без учёта скидки, равна: %v рублей \n", CalculationSumProduct(listproducts))
	bubbleTea.ApplyDiscount(9)
	audi.ApplyDiscount(23)
	phone.ApplyDiscount(8)
	fmt.Printf("Сумма товаров, с учётом скидки, равна: %v рублей \n", CalculationSumProduct(listproducts))
}
