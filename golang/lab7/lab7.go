package lab7

import (
	"fmt"
	"log"
)

func CompleteLab7() {
	var err error

	audi := NewCar(0, "ауди", 80000, Red, 240)
	if audi == nil {
		log.Fatal("Ошибка создания автомобиля")
	}

	desk := NewFurniture(1, "офисный стол", 1200, Wood)
	if desk == nil {
		log.Fatal("Ошибка создания стола")
	}

	chair := NewFurniture(2, "кресло", 300, Skin)
	if chair == nil {
		log.Fatal("Ошибка создания кресла")
	}

	apple := NewFood(3, "огурец", 5)

	var ProductList []Product = []Product{
		audi,
		desk,
		chair,
		apple,
	}

	fmt.Printf("Общая стоимость товара до скидок: %.2f\n", GetTotalPrice(ProductList))
	for _, product := range ProductList {
		if err = product.SetDiscount(10); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Printf("Общая стоимость товара после скидок в 10 процентов: %.2f\n", GetTotalPrice(ProductList))

	if err = audi.SetColor(White); err != nil {
		log.Fatal(err)
	}
	if err = desk.SetMaterial(Plywood); err != nil {
		log.Fatal(err)
	}

	displayListProducts(ProductList)
}
