package lab7

import "fmt"

func RunLab7() {
	product1 := &Clothes{Name: "Футболка", Price: 1000, Size: "L", Color: "Black"}
	product1.sale(15)
	product2 := &Food{Name: "Хотдог", Weight: 100, Price: 100}
	product2.sale(25)
	product3 := &Vehicle{Brand: "BMW", Model: "M5", Color: "Black", MphMax: 250, Price: 4500000}
	product3.sale(5)

	products := []Product{product1, product2, product3}

	for _, product := range products {
		fmt.Println(product.productInfo())
	}

	totalPrice := priceAllProducts(products)
	fmt.Printf("Общая стоимость: %.2f\n", totalPrice)
}