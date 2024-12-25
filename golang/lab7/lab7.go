package lab7

import "fmt"

func Sum_Price(products []Product) float64 {
	sum := 0.0
	for _, product := range products {
		sum += product.Get_price()
	}
	return sum
}

const sale float64 = 0.12

func Show_lab7() {
	product1 := &Food{price: 10.2, name: "Бородинский", category: "Хлеб", brand: "Спас"}
	product2 := &Furniture{price: 300.5, name: "Диван хай-тек", brand: "Good_furniture", color: "red"}
	product3 := &Tools{price: 39.5, name: "отвёртка", brand: "leomax", category: "инструменты для дома"}
	products := []Product{product1, product2, product3}

	fmt.Println("Общая стоимость продуктов до применения скидки: ", Sum_Price(products))

	product1.Sale(sale)
	product2.Sale(sale)
	product3.Sale(sale)

	fmt.Println("Общая стоимость продуктов после применения скидки: ", Sum_Price(products))
	fmt.Println("-----------------------------------------------------------------------")
	fmt.Println("Характеристики продуктов до применения Change_specifications: ")

	list1 := products[0].Get_info()
	list2 := products[1].Get_info()
	list3 := products[2].Get_info()

	fmt.Println("product1: ", list1[0], "|", list1[1], "|", list1[2])
	fmt.Println("product2: ", list2[0], "|", list2[1], "|", list2[2])
	fmt.Println("product3: ", list3[0], "|", list3[1], "|", list3[2])
	fmt.Println("-----------------------------------------------------------------------")
	products[0].Change_specifications("Кекс", "Новогодний", "Русский дар")
	products[1].Change_specifications("Кресло - качалка", "New age", "black")
	products[2].Change_specifications("Молоток", "Kulon", "инструменты для дачи")

	fmt.Println("Характеристики продуктов после применения Change_specifications: ")

	list1 = products[0].Get_info()
	list2 = products[1].Get_info()
	list3 = products[2].Get_info()

	fmt.Println("product1: ", list1[0], "|", list1[1], "|", list1[2])
	fmt.Println("product2: ", list2[0], "|", list2[1], "|", list2[2])
	fmt.Println("product3: ", list3[0], "|", list3[1], "|", list3[2])
}