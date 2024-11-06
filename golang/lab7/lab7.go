package lab7

import "fmt"

func CompleteLab7() {
	var err error

	bentli := NewCar(0, "бентли", 100000, "красный", 200)
	sofa := NewFurniture(1, "диван", 1000, "кожа")
	chear := NewFurniture(2, "стул", 10, "дерево")
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
			panic(err)
		}
	}
	fmt.Printf("общяя стоимость товара после скидок в 10 процентов: %v \n", GetTotalPrice(ProductList))

	err = bentli.SetColor("белый")
	if err != nil {
		panic(err)
	}
	err = sofa.SetMaterial("Фанера")
	if err != nil {
		panic(err)
	}

	Write(ProductList)
}
