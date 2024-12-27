package lab7

import (
	"fmt"
)

func RunLab7() {
	product1 := &Clothes{2000, "серо-буро-малиновый", "XL"}
	product2 := &Car{200000, "белый", "BMW"}
	product3 := &Electronics{35000, "space gray", "256"}
	products := []Product{product1, product2, product3}
	fmt.Println("Общая стоимость товаров:", CalculatePrice(products), "рублей")
	product1.ApplyingTheDiscount(10)
	product2.ApplyingTheDiscount(20)
	product3.ApplyingTheDiscount(30)
	product3.ChangeCharacteristic("black", "512")
	fmt.Println("Общая стоимость товаров после применения скидок:", CalculatePrice(products), "рублей")
	fmt.Println(product3)
}
