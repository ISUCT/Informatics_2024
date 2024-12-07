package lab7

import (
	"fmt"
	"log"
)

func CompleteLab7() {
	var err error

	audi := NewCar(0, "ауди", 80000, Red, 240)
	desk := NewFurniture(1, "офисный стол", 1200, Wood)
	chair := NewFurniture(2, "кресло", 300, Skin)
	apple := Newfood(3, "огурец", 5)

	var ProductList []Product = []Product{
		audi,
		desk,
		chair,
		apple,
	}

	fmt.Printf("Общая стоимость товара до скидок: %.2f \n", GetTotalPrice(ProductList))
	for _, product := range ProductList {
		err = product.SetDiscount(10)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Printf("Общая стоимость товара после скидок в 10 процентов: %.2f \n", GetTotalPrice(ProductList))

	err = audi.SetColor(White)
	if err != nil {
		log.Fatal(err)
	}
	err = desk.SetMaterial(Plywood)
	if err != nil {
		log.Fatal(err)
	}

	displayListProducts(ProductList)
}
