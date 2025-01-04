package lab7

import (
	"fmt"
	"log"
)

func CompleteLab7() {
	audi := NewCar(0, "ауди", 80000, Red, 240)
	desk := NewFurniture(1, "офисный стол", 1200, Wood)
	chair := NewFurniture(2, "кресло", 300, Skin)
	apple := NewFood(3, "огурец", 5)

	ProductList := []Product{
		audi,
		desk,
		chair,
		apple,
	}

	fmt.Printf("Общая стоимость товара до скидок: %.2f \n", GetTotalPrice(ProductList))

	for _, product := range ProductList {
		if err := product.SetDiscount(10); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Printf("Общая стоимость товара после скидок в 10 процентов: %.2f \n", GetTotalPrice(ProductList))

	if err := audi.SetColor(White); err != nil {
		log.Fatal(err)
	}

	if err := desk.SetMaterial(Plywood); err != nil {
		log.Fatal(err)
	}

	displayListProducts(ProductList)
}
