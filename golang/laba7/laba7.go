package laba7

import "fmt"

func RunLab7Task() {
	lamp := Electronics{
		name:  "лампа",
		price: 100,
		model: "диод",
	}
	jacket := Clothes{
		name:     "куртка",
		price:    1000,
		material: "дюспо",
	}
	HarryPotter := Books{
		name:   "Гарри Поттер",
		price:  1000,
		author: "Джоан Роулинг",
	}

	ProductsList := []Product{&lamp, &jacket, &HarryPotter}
	commonCost := getTotalPrice(ProductsList)
	fmt.Println("сумма цен товаров до скидки:", commonCost)

	for _, product := range ProductsList {
		product.SetDiscount(10)
	}

	commonCost = getTotalPrice(ProductsList)
	fmt.Println("сумма цен товаров после скидки:", commonCost)

	lamp.SetModel("лампа накаливания")
	jacket.SetMaterial("таслан")
	HarryPotter.SetAuthor("Джек Торн")
}

func getTotalPrice(ProductsList []Product) float64 {
	var commonCost float64

	for _, product := range ProductsList {
		commonCost += product.GetPrice()
	}
	return commonCost
}
