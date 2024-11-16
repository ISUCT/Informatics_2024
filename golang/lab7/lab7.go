package lab7

import (
	"fmt"
	"log"
)

func CompleteLab7() {
	var err error

	bentli := NewCar(0, "бентли", 100000, Red, 200)
	sofa := NewFurniture(1, "диван", 1000, Skin)
	chear := NewFurniture(2, "стул", 10, Wood)
	cucumber := Newfood(3, "огурец", 5)

	var ProductList []Product = []Product{
		bentli,
		sofa,
		chear,
		cucumber,
	}

	fmt.Printf("общяя стоимость товара до скидок: %v \n", GetTotalPrice(ProductList))
	for _, product := range ProductList {
		err = product.SetDiscount(10)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Printf("общяя стоимость товара после скидок в 10 процентов: %v \n", GetTotalPrice(ProductList))

	err = bentli.SetColor(White)
	if err != nil {
		log.Fatal(err)
	}
	err = sofa.SetMaterial(Plywood)
	if err != nil {
		log.Fatal(err)
	}

	displayListProducts(ProductList)
}
