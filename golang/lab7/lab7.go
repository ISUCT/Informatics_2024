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
	cookies := &Food{price: 149.99, name: "печенье"}
	car := &Car{price: 2000000, name: "Audi", color: "черный"}
	sofa := &Furniture{price: 150000, name: "диван", material: "шенилл"}

	cookies.setPrice(169.99)
	car.setName("BMW")
	car.setColor("белый")
	sofa.setMaterial("кожа")

	listproducts := []Product{cookies, car, sofa}
	fmt.Printf("Сумма товаров, без учёта скидки, равна: %v рублей \n", CalculationSumProduct(listproducts))
	cookies.applyDiscount(5)
	car.applyDiscount(20)
	sofa.applyDiscount(15)
	fmt.Printf("Сумма товаров, с учётом скидки, равна: %v рублей \n", CalculationSumProduct(listproducts))
}
