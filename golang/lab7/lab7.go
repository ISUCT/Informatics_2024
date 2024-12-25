package lab7

import (
	"fmt"
	"log"
)

func CompleteLab7() {
	var err error

	priora := NewBibika(0, "анаконда", 100000, Gold, 200)
	sofa := NewMebel(1, "диван", 1000, Skin)
	armchear := NewMebel(2, "кресло", 10, Wood)
	shaurma := NewVkusnyshki(3, "шаурма", 5)

	var ProductList []Product = []Product{
		priora,
		sofa,
		armchear,
		shaurma,
	}

	fmt.Printf("общяя стоимость товара до скидок: %v \n", GetTotalPrice(ProductList))
	for _, product := range ProductList {
		err = product.SetDiscount(10)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Printf("общяя стоимость товара после скидок в 10 процентов: %v \n", GetTotalPrice(ProductList))

	err = priora.SetColor(Black)
	if err != nil {
		log.Fatal(err)
	}
	err = sofa.SetMaterial(Diamonds)
	if err != nil {
		log.Fatal(err)
	}

	displayListProducts(ProductList)
}
