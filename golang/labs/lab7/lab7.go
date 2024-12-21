package lab7

import "fmt"

func CalculationSumProduct(listproducts []Product) float64 {
	var sum float64 = 0
	for _, product := range listproducts {
		sum += product.getPrice()
	}
	return sum
}

func Laba7run() {
	tv := &Electronics{price: 24999.99, name: "Телевизор", brand: "Samsung", color: "Черный"}
	hoodys := &Hoody{price: 5000, name: "Толстовка Nikifilini", color: "Черный", size: "XXL"}
	trainers := &Shoes{price: 39000, name: "Air Jordan 3 'Black Cat'", brand: "Nike", size: "42"}

	tv.setPrice(29999.99)
	tv.setName("Монитор")
	trainers.setBrand("Puma")

	listproducts := []Product{tv, hoodys, trainers}

	fmt.Printf("Сумма товаров, без учёта скидки, равна: %.2f рублей \n", CalculationSumProduct(listproducts))

	discounts := []float64{5, 25, 15}
	for i, product := range listproducts {
		product.applyDiscount(discounts[i])
	}

	fmt.Printf("Сумма товаров, с учётом скидки, равна: %.2f рублей \n", CalculationSumProduct(listproducts))
}
