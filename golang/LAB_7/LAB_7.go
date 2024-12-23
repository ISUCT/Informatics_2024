package LAB_7

import "fmt"

type Product interface {
	GetName() string
	GetPrice() float64
	SetPrice(price float64)
	ApplyDiscount(discount float64)
}

func Calculate(products []Product) float64 {
	total := 0.0
	for _, product := range products {
		total += product.GetPrice()
	}
	return total
}

func Lab7() {
	iceCream := &IceCream{
		Name:   "Ванильное",
		Brand:  "Веселая коровка",
		Price:  50.00,
		Flavor: "Ваниль",
	}
	sofa := &Sofa{
		Name:     "Классический диван",
		Price:    20000.00,
		Brand:    "IKEA",
		Color:    "Коричневый",
		Material: "Кожа",
	}

	products := []Product{iceCream, sofa}
	fmt.Println("Общая стоимость:", Calculate(products))
	iceCream.ApplyDiscount(10)
	sofa.ApplyDiscount(15)
	fmt.Println("Общая стоимость товаров после применения скидок:", Calculate(products))

	fmt.Println("Материал дивана:", sofa.Material)
	fmt.Println("Цвет дивана:", sofa.Color)
	fmt.Println("Вкус мороженого:", iceCream.Flavor)
	sofa.ChangeColor("Черный")
	sofa.ChangeMaterial("Замша")
	iceCream.ChangeFlavor("Шоколад")
	fmt.Println("Цвет дивана изменен на:", sofa.Color)
	fmt.Println("Материал дивана изменен на:", sofa.Material)
	fmt.Println("Вкус мороженого изменен на:", iceCream.Flavor)
}
