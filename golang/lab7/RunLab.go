package lab7

import "fmt"

func CompleteLab7() {
	product1 := &Clothes{Name: "Худи", Price: 1200, Size: "XL", Color: "Black"}
	product1.make_sale(15)
	product2 := &Food{Name: "Ананас", Weight: 1000, Price: 200}
	product2.make_sale(25)
	product3 := &Vehicle{Brand: "Mitsubishi", Model: "ASX", Color: "Orange", Price: 1400000}
	product3.make_sale(5)

	products := []Product{product1, product2, product3}

	for _, product := range products {
		fmt.Println(product.get_productInfo())
	}

	totalPrice := get_priceAllProducts(products)
	fmt.Printf("Общая стоимость: %.2f\n", totalPrice)
}
