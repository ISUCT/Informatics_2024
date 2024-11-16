package lab7

import (
	"fmt"
)

func CalculationSumProduct(listproducts []Product) string {
	var sum float64 = 0
	for _, product := range listproducts {
		sum += product.getPrice()
	}
	s := fmt.Sprintf("%.2f", sum)
	return s
}

func RunLab7() {
	product1 := &Food{price: 149.99, section: "сладости", nameproduct: "печенья"}
	product2 := &Car{price: 2000000, colour: "черный", brandcar: "BMW"}
	product3 := &Furniture{price: 150000, tupefurniture: "диван", material: "шенилл"}

	product1.setPrice(169.99)
	product2.setCharacteristics("крастный", "BMW")

	listproducts := []Product{product1, product2, product3}
	fmt.Printf("Сумма товаров, без учёта скидки, равна: %v рублей \n", CalculationSumProduct(listproducts))
	product1.discount(5)
	product2.discount(20)
	product3.discount(15)
	fmt.Printf("Сумма товаров, с учётом скидки, равна: %v рублей \n", CalculationSumProduct(listproducts))
}
