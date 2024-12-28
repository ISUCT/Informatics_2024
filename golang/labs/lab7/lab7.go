package labs

import "fmt"

func CalculateProductsSum(products []Product) float64 {
	var sum float64 = 0
	for _, product := range products {
		sum += product.GetPrice()
	}
	return sum
}

func RunLab7() {
	carrot := &Food{50.15, "морковь", "магнит", 17}
	mascara := &Cosmetics{"тушь для ресниц", 1700, "loreal paris"}
	shirt := &Clothes{"рубашка", 3000.99, "Levi`s", "синтетика", "весна"}
	products := []Product{carrot, mascara, shirt}
	fmt.Println("товары")
	fmt.Println("общая стоимость:", CalculateProductsSum(products), "рублей")
	for _, product := range products {
		fmt.Println(product.GetInf())
	}

	for _, product := range products {
		product.SetDiscount(20)
	}
	fmt.Println("общая стоимость товаров после применения скидки 20%:", CalculateProductsSum(products), "рублей")

	fmt.Println("информация про товар рубашка")
	fmt.Println("материал:", shirt.GetMat())
	fmt.Println("сезон:", shirt.GetSeason())
	fmt.Println("бренд:", shirt.GetBrand())
}
