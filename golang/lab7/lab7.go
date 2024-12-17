package lab7

import (
	"fmt"
	"log"
)

func Completelab7() {
	var err error

	priora := Newbibika(0, "анаконда", 100000, Gold, 200)
	sofa := Newmebel(1, "диван", 1000, Skin)
	armchear := Newmebel(2, "кресло", 10, Wood)
	shaurma := Newvkusnyshki(3, "шаурма", 5)

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
// Ответ на комментарий по строкам 32-39("Используйте методы из интерфейсного экземпляра, а не структуры"):

// Может что путаю, но я меняю характеристики отдельных структур, их вроде бы нельзя через интерфейс менять.
// Или нужно делать через общии методы в виде "func name(a any) {}" ?
